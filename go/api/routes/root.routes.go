package routes

import "github.com/gin-gonic/gin"

func RootRoutes(group *gin.RouterGroup) {
	group.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Sowaste API!\nFollow the documentation at: https://documenter.getpostman.com/view/21306535/2s8ZDSdRaf#intro",
		})
	})
}
