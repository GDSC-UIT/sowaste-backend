package routes

import (
	"github.com/GDSC-UIT/sowaste-backend/go/api/handlers"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func QRRoutes(group *gin.RouterGroup, db *mongo.Client) {
	handler := handlers.QRHandlers{
		Handler: services.QRServices{
			Db: db,
		},
	}

	qrs := group.Group("/qr")
	{
		qrs.GET("", handler.GetQRs)
		qrs.GET("/:id", handler.GetAQR)
		qrs.POST("", handler.CreateAQR)
		qrs.PUT("/:id", handler.UpdateAQR)
		qrs.DELETE("/:id", handler.DeleteAQR)
		qrs.POST("/scan/:id", handler.ScanQR)
		qrs.GET("/generate", handler.GenerateQRCode)
	}
}
