package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type BottleShopHandlers struct {
	Handler services.BottleShopServices
}

func (bh *BottleShopHandlers) GetBottleShops(c *gin.Context) {

	//* called from bottle-shop.services.go **//
	/*
	* @params c; type *gin.Context;
	* @return => void
	 */
	bh.Handler.GetBottleShops(c)
}

func (bh *BottleShopHandlers) GetABottleShop(c *gin.Context) {
	bh.Handler.GetABottleShop(c)
}

func (bh *BottleShopHandlers) CreateABottleShop(c *gin.Context) {
	bh.Handler.CreateABottleShop(c)
}

func (bh *BottleShopHandlers) UpdateABottleShop(c *gin.Context) {
	bh.Handler.UpdateABottleShop(c)
}

func (bh *BottleShopHandlers) DeleteABottleShop(c *gin.Context) {
	bh.Handler.DeleteABottleShop(c)
}
