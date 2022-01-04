package main

/*
1. Имплементировать регистрацию и вход
2. Разрешить проблему с хендлерами
	a) Со складами
	б) Со страницами
3. Проверить работу
*/
import (
	"log"
	"os"

	"github.com/joho/godotenv"
	agregator "github.com/seedovan19/Agregator"
	"github.com/seedovan19/Agregator/pkg/handler"
	"github.com/seedovan19/Agregator/pkg/repository"
	"github.com/seedovan19/Agregator/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	// if err := initConfig(); err != nil {
	// 	log.Fatalf("Error initializing configs: %s", err.Error())
	// }

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading enviroment variables: %s", err.Error())
	}

	db, err := repository.Open_connection(repository.Config{
		// Host:     viper.GetString("db.host"),
		// Port:     viper.GetString("db.port"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
		// DBName:   viper.GetString("db.dbname"),
		// SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Failed to initialize database: %s", err.Error())
	}
	// делаем миграцию в базу данных (создаем таблицы, если они еще не были созданы)
	db.AutoMigrate(&agregator.Warehouse{})
	db.AutoMigrate(&agregator.Building{})

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(agregator.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("/../configs")
	viper.SetConfigFile("config")
	return viper.ReadInConfig()
}
