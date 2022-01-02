package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type Warehouse struct {
	gorm.Model

	ID                 int64  `gorm:"primaryKey"`
	Name               string `gorm:"not null"`
	Square             int64  `gorm:"not null"`
	Adress             string `gorm:"not null"`
	Shelf_storage_cost int64  `gorm:"not null"`
	Floor_storage_cost int64  `gorm:"not null"`
	Image              string `gorm:"image"`
	Description        string
	Comment            string
	Building           Building `gorm:"foreignkey: id; references: ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Building struct {
	gorm.Model

	Warehouse_class      string `gorm:"not null"` //выпадайка
	Year_of_construction time.Time
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

func Add_warehouse_record(warehouse_data Warehouse, building_data Building) {
	db.Create(&warehouse_data)
	db.Create(&building_data)
}
