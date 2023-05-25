package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RewardRoutes(group *gin.RouterGroup, db *mongo.Client) {
	handler := handlers.RewardHandlers{
		Handler: services.RewardServices{
			Db: db,
		},
	}

	rewards := group.Group("/reward")
	{
		rewards.GET("", handler.GetRewards)
		rewards.GET("/:id", handler.GetAReward)
		rewards.GET("/user", handler.GetUserRewards)
		rewards.POST("", handler.CreateAReward)
		rewards.PUT("/:id", handler.UpdateAReward)
		rewards.DELETE("/:id", handler.DeleteAReward)
	}
}
