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

	routes.Ocupations(router)
	routes.Deaths(router)

	router.Run(":" + port)
}
