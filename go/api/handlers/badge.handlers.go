package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type BadgeHandlers struct {
	Handler services.BadgeServices
}

func (bh *BadgeHandlers) GetBadges(c *gin.Context) {
	bh.Handler.GetBadges(c)
}

func (bh *BadgeHandlers) GetABadge(c *gin.Context) {
	bh.Handler.GetABadge(c)
}

func (bh *BadgeHandlers) CreateABadge(c *gin.Context) {
	bh.Handler.CreateBadge(c)
}

func (bh *BadgeHandlers) UpdateABadge(c *gin.Context) {
	bh.Handler.UpdateBadge(c)
}

func (bh *BadgeHandlers) DeleteABadge(c *gin.Context) {
	bh.Handler.DeleteBadge(c)
}
