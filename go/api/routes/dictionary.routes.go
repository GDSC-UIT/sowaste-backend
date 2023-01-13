package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func DictonaryRoutes(group *gin.RouterGroup, db *mongo.Client) {

	handler := handlers.DictionaryHandlers{
		Handler: services.DictionaryServices{
			Db: db,
		},
	}
	dictionaries := group.Group("/dictionaries")
	{
		dictionaries.GET("", handler.GetDictionaries)
		dictionaries.GET("/:id", handler.GetADictionary)
		dictionaries.POST("", handler.CreateADictionary)
		dictionaries.PUT("/:id", handler.UpdateADictionary)
		dictionaries.DELETE("/:id", handler.DeleteADictionary)
	}
}
