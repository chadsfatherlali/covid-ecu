package main

import (
	"log"
	"os"

	"github.com/heroku/go-getting-started/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	routes.Ocupations(router)
	routes.Deaths(router)

	router.Run(":" + port)
}
