package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func LessonRoutes(group *gin.RouterGroup, db *mongo.Client) {

	handler := handlers.LessonHandlers{
		Handler: services.LessonServices{
			Db: db,
		},
	}
	dictionaries := group.Group("/lessons")
	{
		dictionaries.GET("", handler.GetLessons)
		dictionaries.GET("/:id", handler.GetALesson)
		dictionaries.POST("", handler.CreateALesson)
		dictionaries.PUT("/:id", handler.UpdateALesson)
		dictionaries.DELETE("/:id", handler.DeleteALesson)
	}
}
