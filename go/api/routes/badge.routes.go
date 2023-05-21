package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func BadgeRoutes(group *gin.RouterGroup, db *mongo.Client) {
	handler := handlers.BadgeHandlers{
		Handler: services.BadgeServices{
			Db: db,
		},
	}

	badges := group.Group("/badge")
	{
		badges.GET("", handler.GetBadges)
		badges.GET("/:id", handler.GetABadge)
		badges.POST("", handler.CreateABadge)
		badges.PUT("/:id", handler.UpdateABadge)
		badges.DELETE("/:id", handler.DeleteABadge)
	}
}
