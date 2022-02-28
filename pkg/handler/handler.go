package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/seedovan19/Agregator/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/assets", "./assets")
	router.Static("/images", "./images")
	router.Static("/javascript", "./javascript")
	router.LoadHTMLGlob("./templates/*")

	router.GET("/", h.Home_page)
	router.GET("/form", h.Warehouse_form_page)

	api := router.Group("/api")
	{
		warehouses := api.Group("/warehouses")
		{
			warehouses.POST("/add_warehouse", h.Add_warehouse)
		}
	}
	return router
}
