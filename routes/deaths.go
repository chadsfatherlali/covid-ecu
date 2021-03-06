package routes

import (
	"github.com/heroku/go-getting-started/structures"

	"github.com/gin-gonic/gin"
)

func Deaths(routes *gin.RouterGroup) {
	{
		routes.GET("/deaths", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"data": structures.Deaths(),
			})
		})
	}
}
