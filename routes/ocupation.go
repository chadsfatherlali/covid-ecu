package routes

import (
	"github.com/heroku/go-getting-started/structures"

	"github.com/gin-gonic/gin"
)

func Ocupations(router *gin.Engine) {
	ocupationsRoutes := router.Group("/api/v1/ocupations")

	{
		ocupationsRoutes.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"data": structures.Ocupations(),
			})
		})
	}
}
