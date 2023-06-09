package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageController struct{}

func NewPageController() *PageController {
	return &PageController{}
}

func (controller *PageController) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": gin.H{
			"name": "Awesome API",
		},
		"error": nil,
	})
}
