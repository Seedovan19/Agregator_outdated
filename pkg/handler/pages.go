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

	header := c.GetHeader(authorizationHeader)
	if header != "" {
		log.Print(header)
		_, err := h.services.Authorization.ParseToken(header)
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		c.HTML(
			http.StatusOK,
			"home_page.gohtml",
			gin.H{
				"posts":  posts,
				"IsAuth": true,
			},
		)
		return
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

func (h *Handler) Auth_form_page(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"auth_form.gohtml",
		nil,
	)
}

func (h *Handler) Sign_up_form_page(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"sign_up_form.gohtml",
		nil,
	)
}
