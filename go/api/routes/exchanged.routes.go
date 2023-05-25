package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ExchangedRoutes(group *gin.RouterGroup, db *mongo.Client) {
	handler := handlers.ExchangedHandlers{
		Handler: services.ExchangedServices{
			Db: db,
		},
	}

	exchangeds := group.Group("/exchanged")
	{
		exchangeds.GET("", handler.GetExchanges)
		exchangeds.GET("/:id", handler.GetAExchange)
		exchangeds.GET("/user/:user_id", handler.GetExchangedsByUserId)
		exchangeds.GET("/user", handler.GetCurrentUserExchangeds)
		exchangeds.POST("", handler.CreateAExchange)
		exchangeds.PUT("/:id", handler.UpdateAExchange)
		exchangeds.DELETE("/:id", handler.DeleteAExchange)
	}
}
