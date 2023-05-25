package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoutes(group *gin.RouterGroup, db *mongo.Client) {
	handler := handlers.UserHandlers{
		Handler: services.UserServices{
			Db: db,
		},
	}

	users := group.Group("/user")
	{
		// users.GET("", handler.GetUsers)
		users.GET("", handler.GetAUser)
		users.GET("/:id", handler.GetAUserById)
		users.POST("", handler.CreateAUser)
		users.PUT("/:id", handler.UpdateAUser)
		users.DELETE("/:id", handler.DeleteAUser)
	}
}
