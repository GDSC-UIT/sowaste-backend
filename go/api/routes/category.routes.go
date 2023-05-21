package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CategoryRoutes(group *gin.RouterGroup, db *mongo.Client) {

	handler := handlers.CategoryHandlers{
		Handler: services.CategoryServices{
			Db: db,
		},
	}
	articles := group.Group("/categories")
	{
		articles.GET("", handler.GetCategories)
		articles.GET("/:id", handler.CreateAnCategory)
		articles.POST("", handler.CreateAnCategory)
		articles.PUT("/:id", handler.UpdateAnCategory)
		articles.DELETE("/:id", handler.DeleteAnCategory)
	}
}
