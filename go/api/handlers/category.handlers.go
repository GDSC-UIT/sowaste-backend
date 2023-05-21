package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type CategoryHandlers struct {
	Handler services.CategoryServices
}

func (ch *CategoryHandlers) GetCategories(c *gin.Context) {
	ch.Handler.GetCategories(c)
}

func (ch *CategoryHandlers) GetAnCategory(c *gin.Context) {
	ch.Handler.GetAnCategory(c)
}

func (ch *CategoryHandlers) CreateAnCategory(c *gin.Context) {
	ch.Handler.CreateCategory(c)
}

func (ch *CategoryHandlers) UpdateAnCategory(c *gin.Context) {
	ch.Handler.UpdateCategory(c)
}

func (ch *CategoryHandlers) DeleteAnCategory(c *gin.Context) {
	ch.Handler.DeleteCategory(c)
}
