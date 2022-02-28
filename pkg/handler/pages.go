package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Home_page(c *gin.Context) {
	posts, err := h.services.Warehouse.Show_warehouse_records()
	if err != nil {
		log.Print(err.Error())
	}

	c.HTML(
		http.StatusOK,
		"home_page.gohtml",
		gin.H{
			"posts":  posts,
			"IsAuth": false,
		},
	)
}

func (h *Handler) Warehouse_form_page(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"form.gohtml",
		nil,
	)
}
