package repository

import (
	agregator "github.com/seedovan19/Agregator"
	"gorm.io/gorm"
)

type WarehousePostgres struct {
	db *gorm.DB
}

func NewWarehousePostgres(db *gorm.DB) *WarehousePostgres {
	return &WarehousePostgres{db: db}
}

func (r *WarehousePostgres) Add_warehouse_record(warehouse_data agregator.Warehouse) (int, error) {
	result := r.db.Create(&warehouse_data)
	return int(warehouse_data.ID), result.Error
}

func (r *WarehousePostgres) Show_warehouse_records() ([]agregator.Warehouse, error) {
	var posts = []agregator.Warehouse{}
	var warehouse agregator.Warehouse

	res, err := r.db.Raw("SELECT name, square, adress, shelf_storage_cost, floor_storage_cost, description FROM warehouses").Rows()

	for res.Next() {
		res.Scan(&warehouse.Name, &warehouse.Square, &warehouse.Adress, &warehouse.Shelf_storage_cost, &warehouse.Floor_storage_cost, &warehouse.Description)
		posts = append(posts, warehouse)
	}
	return posts, err
}
