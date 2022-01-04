package repository

import (
	"fmt"
	"log"

	agregator "github.com/seedovan19/Agregator"
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
	dbURI := fmt.Sprintf("host = %s user = %s dbname = %s sslmode = disable password = %s port = %s", cfg.Host, cfg.Username, cfg.DBName, cfg.Password, cfg.Port)

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

func Add_warehouse_record(warehouse_data agregator.Warehouse, building_data agregator.Building) {
	db.Create(&warehouse_data)
	db.Create(&building_data)
}

func Show_warehouse_records() ([]agregator.Warehouse, error) {
	var posts = []agregator.Warehouse{}
	var warehouse agregator.Warehouse

	res, err := db.Raw("SELECT name, square, adress, shelf_storage_cost, floor_storage_cost, image, description FROM warehouses").Rows()

	for res.Next() {
		res.Scan(&warehouse.Name, &warehouse.Square, &warehouse.Adress, &warehouse.Shelf_storage_cost, &warehouse.Floor_storage_cost, &warehouse.Image, &warehouse.Description)
		posts = append(posts, warehouse)
	}
	return posts, err
}
