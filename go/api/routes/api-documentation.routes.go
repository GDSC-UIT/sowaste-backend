package routes

import (
	"github.com/gin-gonic/gin"
)

func APIDocumentationRoutes(root *gin.Engine) {
	root.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"title": "Sowaste API Documentation",
		})
	})
}
