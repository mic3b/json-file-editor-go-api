package main

import (
	"translations/editor/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes:
	translations := router.Group("/translations")
	{
		translations.POST("/split", controllers.SplitController)
		translations.POST("/edit", controllers.EditController)
		translations.POST("/finish", controllers.FinishController)

	}

	// Router Setup
	router.Use(cors.Default())
	router.Run()
}
