package controllers

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
)

type Request struct {
	Path  string
	Value string
}

func EditController(c *gin.Context) {
	var Req Request

	c.ShouldBindJSON(&Req)

	data, err := ioutil.ReadFile("temp/file.json")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	scontent := string(data)
	scontent, err = sjson.Set(scontent, Req.Path, Req.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	f, _ := os.Create("temp/file.json")
	f.WriteString(scontent)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Element has changed!",
	})
}
