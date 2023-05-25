package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type BadgeCollectionHandlers struct {
	Handler services.BadgeCollectionServices
}

func (bch *BadgeCollectionHandlers) GetBadgeCollections(c *gin.Context) {
	bch.Handler.GetBadgeCollections(c)
}

func (bch *BadgeCollectionHandlers) GetABadgeCollection(c *gin.Context) {
	bch.Handler.GetABadgeCollection(c)
}

func (bch *BadgeCollectionHandlers) GetCurrentUserBadgeCollections(c *gin.Context) {
	bch.Handler.GetCurrentUserBadgeCollection(c)
}

func (bch *BadgeCollectionHandlers) GetBadgeCollectionsByUserId(c *gin.Context) {
	bch.Handler.GetBadgeCollectionByUserId(c)
}

func (bch *BadgeCollectionHandlers) CreateABadgeCollection(c *gin.Context) {
	bch.Handler.CreateBadgeCollection(c)
}

func (bch *BadgeCollectionHandlers) UpdateABadgeCollection(c *gin.Context) {
	bch.Handler.UpdateBadgeCollection(c)
}

func (bch *BadgeCollectionHandlers) DeleteABadgeCollection(c *gin.Context) {
	bch.Handler.DeleteBadgeCollection(c)
}
