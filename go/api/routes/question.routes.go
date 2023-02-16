package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func QuizRoutes(group *gin.RouterGroup, db *mongo.Client) {

	handler := handlers.QuizHandlers{
		Handler: services.QuizServices{
			Db: db,
		},
	}
	dictionaries := group.Group("/questions")
	{
		dictionaries.GET("", handler.GetQuizzes)
		dictionaries.GET("/:id", handler.GetAQuiz)
		dictionaries.POST("", handler.CreateAQuiz)
		dictionaries.PUT("/:id", handler.UpdateAQuiz)
		dictionaries.DELETE("/:id", handler.DeleteAQuiz)
	}
}
