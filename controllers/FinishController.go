package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func FinishController(c *gin.Context) {
	os.Remove("temp/file.json")

	c.JSON(http.StatusCreated, gin.H{
		"message": "File has removed!",
	})
}
