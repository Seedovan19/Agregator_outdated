package repository

import (
	agregator "github.com/seedovan19/Agregator"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user agregator.User) (int, error)
	GetUser(email, password string) (agregator.User, error)
}

type Warehouse interface {
	Add_warehouse_record(warehouse_data agregator.Warehouse, building_data agregator.Building) (int, error)
	Show_warehouse_records() ([]agregator.Warehouse, error)
}

type Repository struct {
	Authorization
	Warehouse
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Warehouse:     NewWarehousePostgres(db),
	}
}
