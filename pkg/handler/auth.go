package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	agregator "github.com/seedovan19/Agregator"
)

func (h *Handler) signUp(c *gin.Context) {
	var input agregator.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

}

func (h *Handler) signIn(c *gin.Context) {}
