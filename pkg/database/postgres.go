package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gorm.io/datatypes"
)

var db *gorm.DB
var err error

type Warehouse struct {
	gorm.Model

	Id                 int    `gorm:"primaryKey;autoIncrement:true"`
	Name               string `gorm:"not null"`
	Square             int    `gorm:"not null"`
	Adress             string `gorm:"not null"`
	Shelf_storage_cost int    `gorm:"not null"`
	Floor_storage_cost int    `gorm:"not null"`
	Description        string
	Comment            string
}

type Building struct {
	WarehouseID          int    `gorm:"not null"`
	Warehouse_class      string `gorm:"not null"` //выпадайка
	Year_of_construction datatypes.Date
	Warehouse            Warehouse `gorm:"ForeignKey:WarehouseID;References: Id"`
}

func Open_connection() {
	// загружаем переменные окружения
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// строка соединения с базой данных
	dbURI := fmt.Sprintf("host = %s user = %s dbname = %s sslmode = disable password = %s port = %s", host, user, dbName, password, dbPort)

	// открываем соединение с базой данных
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}
}

func Migrate() {
	db.AutoMigrate(&Warehouse{})
	db.AutoMigrate(&Building{})
}

func Close() {
	//закрываем соединение с базой данных
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}
