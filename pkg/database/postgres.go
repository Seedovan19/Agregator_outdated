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

type Person struct {
	ID       int64  `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Surname  string `gorm:"not null"`
	Company  string `gorm:"not null"`
	Phone    string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}

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
	// Phone              string
}

type Add_warehouse struct {
	Warehouse_ID int64
	Person_ID    int64
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

func Show_warehouse_records() ([]Warehouse, error) {
	var posts = []Warehouse{}
	var warehouse Warehouse

	res, err := db.Raw("SELECT name, square, adress, shelf_storage_cost, floor_storage_cost, image, description FROM warehouses").Rows()

	for res.Next() {
		res.Scan(&warehouse.Name, &warehouse.Square, &warehouse.Adress, &warehouse.Shelf_storage_cost, &warehouse.Floor_storage_cost, &warehouse.Image, &warehouse.Description)
		posts = append(posts, warehouse)
	}
	return posts, err
}
