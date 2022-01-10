package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	agregator "github.com/seedovan19/Agregator"
)

type WarehouseInsert struct {
	Name                 string    `json: "name" binding: "required"`
	Square               int64     `json: "square" binding: "required"`
	Adress               string    `json: "adress" binding: "required"`
	Shelf_storage_cost   int64     `json: "shelf_cost" binding: "required"`
	Floor_storage_cost   int64     `json: "floor_cost" binding: "required"`
	Description          string    `json: "description" binding: "required"`
	Warehouse_class      string    `json: "class" binding: "required"`
	Year_of_construction time.Time `json: "age_of_construction" binding: "required"`
}

func (h *Handler) Add_warehouse(c *gin.Context) {

	// adress := c.PostForm("adress")
	// name := c.PostForm("name")
	// square, _ := strconv.ParseInt(c.PostForm("square")[0:], 10, 64)
	// class := c.PostForm("class")
	// age_of_construction, _ := time.Parse(c.PostForm("age_of_construction"), "2019-07-10")
	// shelf_cost, _ := strconv.ParseInt(c.PostForm("shelf_cost")[0:], 10, 64)
	// floor_cost, _ := strconv.ParseInt(c.PostForm("floor_cost")[0:], 10, 64)
	// description := c.PostForm("description")

	// img, err := imageupload.Process(c.Request, "file")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// thumb.Save("../images/saved_images") // правильно сохраняет фотку??

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
	building := agregator.Building{
		Warehouse_class:      insertingValue.Warehouse_class,
		Year_of_construction: insertingValue.Year_of_construction,
	}

	id, err := h.services.Warehouse.Add_warehouse_record(warehouse, building)

	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Inserted warehouse with id: %d", id)

	c.Redirect(http.StatusMovedPermanently, "/")
}
