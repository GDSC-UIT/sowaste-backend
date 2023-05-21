package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func BadgeCollectionRoutes(group *gin.RouterGroup, db *mongo.Client) {
	handler := handlers.BadgeCollectionHandlers{
		Handler: services.BadgeCollectionServices{
			Db: db,
		},
	}

	badgeCollections := group.Group("/badge-collection")
	{
		badgeCollections.GET("", handler.GetBadgeCollections)
		badgeCollections.GET("/:id", handler.GetABadgeCollection)
		badgeCollections.GET("/user/:user_id", handler.GetBadgeCollectionsByUserId)
		badgeCollections.POST("", handler.CreateABadgeCollection)
		badgeCollections.PUT("/:id", handler.UpdateABadgeCollection)
		badgeCollections.DELETE("/:id", handler.DeleteABadgeCollection)
	}
}
