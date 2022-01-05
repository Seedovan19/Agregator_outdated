package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	agregator "github.com/seedovan19/Agregator"
)

func (h *Handler) Add_warehouse(c *gin.Context) {

	adress := c.PostForm("adress")
	name := c.PostForm("name")
	square, _ := strconv.ParseInt(c.PostForm("square")[0:], 10, 64)
	class := c.PostForm("class")
	age_of_construction, _ := time.Parse(c.PostForm("age_of_construction"), "2019-07-10")
	shelf_cost, _ := strconv.ParseInt(c.PostForm("shelf_cost")[0:], 10, 64)
	floor_cost, _ := strconv.ParseInt(c.PostForm("floor_cost")[0:], 10, 64)
	description := c.PostForm("description")

	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		log.Fatalf(err.Error())
	}

	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

	if err != nil {
		log.Fatalf(err.Error())
	}

	thumb.Write(c.Writer)

	if adress == "" || name == "" || class == "" {
		log.Printf("Форма заполнена некорректно")
	} else {
		warehouse := agregator.Warehouse{
			Name:               name,
			Square:             square,
			Adress:             adress,
			Shelf_storage_cost: shelf_cost,
			Floor_storage_cost: floor_cost,
			Description:        description,
		}

		building := agregator.Building{
			Warehouse_class:      class,
			Year_of_construction: age_of_construction,
		}

		h.services.Warehouse.Add_warehouse_record(warehouse, building)

		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
