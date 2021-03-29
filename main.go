package main

import (
	"os"

	"github.com/heroku/go-getting-started/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	covidRoutes := router.Group("/api/v1")

	routes.Occupations(covidRoutes)
	routes.Deaths(covidRoutes)

	router.Run(":" + port)
}
