package repository

import agregator "github.com/seedovan19/Agregator"

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
