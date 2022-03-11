package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	agregator "github.com/seedovan19/Agregator"
)

type WarehouseInsert struct {
	Name               string `json:"name" binding:"required"`
	Square             int64  `json:"square" binding:"required"`
	Adress             string `json:"adress" binding:"required"`
	Shelf_storage_cost int64  `json:"shelf_cost" binding:"required"`
	Floor_storage_cost int64  `json:"floor_cost" binding:"required"`
	Description        string `json:"description" binding:"required"`
	Warehouse_class    string `json:"class" binding:"required"`
	// Year_of_construction time.Time `json:"age_of_construction" binding:"required"` //TODO: ему не нравится с форматом времени что-то
}

func (h *Handler) Add_warehouse(c *gin.Context) {
	var insertingValue WarehouseInsert

	if err := c.BindJSON(&insertingValue); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Print(insertingValue.Floor_storage_cost)
	warehouse := agregator.Warehouse{
		Name:               insertingValue.Name,
		Square:             insertingValue.Square,
		Adress:             insertingValue.Adress,
		Shelf_storage_cost: insertingValue.Shelf_storage_cost,
		Floor_storage_cost: insertingValue.Floor_storage_cost,
		Description:        insertingValue.Description,
	}

	id, err := h.services.Warehouse.Add_warehouse_record(warehouse)

	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Inserted warehouse with id: %d", id)
}

func (h *Handler) getWarehouseRecords(c *gin.Context) {
	posts, err := h.services.Warehouse.Show_warehouse_records()
	if err != nil {
		log.Print(err.Error())
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, posts)
}
