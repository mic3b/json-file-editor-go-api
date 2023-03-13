package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func FinishController(c *gin.Context) {
	c.File("temp/file.json")
	os.Remove("temp/file.json")
}
