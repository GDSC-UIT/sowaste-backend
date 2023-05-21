package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func DIYRoutes(group *gin.RouterGroup, db *mongo.Client) {
	handler := handlers.DIYHandlers{
		Handler: services.DIYServices{
			Db: db,
		},
	}

	diys := group.Group("/diy")
	{
		diys.GET("", handler.GetDIYs)
		diys.GET("/:id", handler.GetADIY)
		diys.POST("", handler.CreateADIY)
		diys.PUT("/:id", handler.UpdateADIY)
		diys.DELETE("/:id", handler.DeleteADIY)
	}
}
