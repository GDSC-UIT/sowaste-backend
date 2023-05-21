package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SavedRoutes(group *gin.RouterGroup, db *mongo.Client) {
	handler := handlers.SavedHandlers{
		Handler: services.SavedServices{
			Db: db,
		},
	}

	saveds := group.Group("/saved")
	{
		saveds.GET("", handler.GetSaveds)
		saveds.GET("/:id", handler.GetASaved)
		saveds.GET("/user/:user_id", handler.GetSavedsByUserId)
		saveds.POST("", handler.CreateASaved)
		saveds.PUT("/:id", handler.UpdateASaved)
		saveds.DELETE("/:id", handler.DeleteASaved)
		saveds.DELETE("/user/:dictionary_id", handler.DeleteASavedByDictionaryId)
	}
}
