package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	agregator "github.com/seedovan19/Agregator"
	"github.com/seedovan19/Agregator/pkg/handler"
	"github.com/seedovan19/Agregator/pkg/repository"
	"github.com/seedovan19/Agregator/pkg/service"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Printf("Error loading enviroment variables: %s", err.Error())
	}

	db, err := repository.Open_connection(repository.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DBPORT"),
		Username: os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	})

	if err != nil {
		log.Fatalf("Failed to initialize database: %s", err.Error())
	}
	// делаем миграцию в базу данных (создаем таблицы, если они еще не были созданы)
	db.AutoMigrate(&agregator.Warehouse{})
	db.AutoMigrate(&agregator.Building{})
	db.AutoMigrate(&agregator.User{})

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(agregator.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

// func initConfig() error {
// 	viper.AddConfigPath("/../configs")
// 	viper.SetConfigFile("config")
// 	return viper.ReadInConfig()
// }
