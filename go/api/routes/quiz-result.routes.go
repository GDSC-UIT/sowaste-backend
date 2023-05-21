package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func QuizResultRoutes(group *gin.RouterGroup, db *mongo.Client) {
	handler := handlers.QuizResultHandlers{
		Handler: services.QuizResultServices{
			Db: db,
		},
	}

	quizResults := group.Group("/quiz-result")
	{
		quizResults.GET("", handler.GetQuizResults)
		quizResults.GET("/:id", handler.GetAQuizResult)
		quizResults.GET("/user/:user_id", handler.GetQuizResultsByUserId)
		quizResults.POST("", handler.CreateAQuizResult)
		quizResults.PUT("/:id", handler.UpdateAQuizResult)
		quizResults.DELETE("/:id", handler.DeleteAQuizResult)
	}
}
