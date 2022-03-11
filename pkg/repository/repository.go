package repository

import (
	agregator "github.com/seedovan19/Agregator"
	"gorm.io/gorm"
)

type Warehouse interface {
	Add_warehouse_record(warehouse_data agregator.Warehouse) (int, error)
	Show_warehouse_records() ([]agregator.Warehouse, error)
}

type Repository struct {
	Warehouse
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Warehouse: NewWarehousePostgres(db),
	}
}
