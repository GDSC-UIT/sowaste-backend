package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func QuestionRoutes(group *gin.RouterGroup, db *mongo.Client) {

	handler := handlers.QuestionHandlers{
		Handler: services.QuestionServices{
			Db: db,
		},
	}
	dictionaries := group.Group("/questions")
	{
		dictionaries.GET("", handler.GetQuestions)
		dictionaries.POST("", handler.CreateAQuestion)
		dictionaries.PUT("/:id", handler.UpdateAQuestion)
		dictionaries.DELETE("/:id", handler.DeleteAQuestion)
	}
}
