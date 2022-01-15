package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

var db *gorm.DB

func Open_connection(cfg Config) (*gorm.DB, error) {
	// строка соединения с базой данных
	dbURI := fmt.Sprintf("host = %s user = %s dbname = %s password = %s port = %s", cfg.Host, cfg.Username, cfg.DBName, cfg.Password, cfg.Port)
	// connection := os.Getenv("DATABASE_URL")

	// открываем соединение с базой данных
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database connection: %s", err.Error())
		return nil, err
	} else {
		fmt.Println("Successfully connected to database")
	}

	return db, nil
}

func Close() {
	//закрываем соединение с базой данных
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}
