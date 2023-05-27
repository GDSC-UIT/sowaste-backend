package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type QRHandlers struct {
	Handler services.QRServices
}

func (qrh *QRHandlers) GetQRs(c *gin.Context) {
	qrh.Handler.GetQRs(c)
}

func (qrh *QRHandlers) GetAQR(c *gin.Context) {
	qrh.Handler.GetAQR(c)
}

func (qrh *QRHandlers) CreateAQR(c *gin.Context) {
	qrh.Handler.CreateQR(c)
}

func (qrh *QRHandlers) GenerateQRCode(c *gin.Context) {
	qrh.Handler.GenerateQRCode(c)
}

func (qrh *QRHandlers) UpdateAQR(c *gin.Context) {
	qrh.Handler.UpdateQr(c)
}

func (qrh *QRHandlers) DeleteAQR(c *gin.Context) {
	qrh.Handler.DeleteQR(c)
}

func (qrh *QRHandlers) ScanQR(c *gin.Context) {
	qrh.Handler.ScanQR(c)
}
