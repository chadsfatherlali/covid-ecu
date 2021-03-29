package routes

import (
	"github.com/heroku/go-getting-started/structures"

	"github.com/gin-gonic/gin"
)

func Occupations(routes *gin.RouterGroup) {
	{
		routes.GET("/occupations", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"data": structures.Occupations(),
			})
		})
	}
}
