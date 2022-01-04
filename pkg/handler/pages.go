package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seedovan19/Agregator/pkg/repository"
)

func (h *Handler) Home_page(c *gin.Context) {
	// tpl, err := template.ParseFiles("../templates/home_page.gohtml")
	// if err != nil {
	// 	log.Print(err.Error())
	// }

	posts, err := repository.Show_warehouse_records()
	if err != nil {
		log.Print(err.Error())
	}

	c.HTML(
		http.StatusOK,
		"home_page.gohtml",
		posts,
	)
	// tpl.Execute(w, posts)
}

func (h *Handler) Warehouse_form_page(c *gin.Context) {
	// tpl, err := template.ParseFiles("../templates/form.gohtml")
	// if err != nil {
	// 	log.Print(err.Error())
	// }
	// tpl.Execute(w, nil)

	c.HTML(
		http.StatusOK,
		"form.gohtml",
		nil,
	)
}
