package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type DIYHandlers struct {
	Handler services.DIYServices
}

func (dh *DIYHandlers) GetDIYs(c *gin.Context) {
	dh.Handler.GetDIYs(c)
}

func (dh *DIYHandlers) GetADIY(c *gin.Context) {
	dh.Handler.GetAnDIY(c)
}

func (dh *DIYHandlers) CreateADIY(c *gin.Context) {
	dh.Handler.CreateDIY(c)
}

func (dh *DIYHandlers) UpdateADIY(c *gin.Context) {
	dh.Handler.UpdateDIY(c)
}

func (dh *DIYHandlers) DeleteADIY(c *gin.Context) {
	dh.Handler.DeleteDIY(c)
}
