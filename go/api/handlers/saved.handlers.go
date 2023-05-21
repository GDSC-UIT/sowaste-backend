package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type SavedHandlers struct {
	Handler services.SavedServices
}

func (sh *SavedHandlers) GetSaveds(c *gin.Context) {
	sh.Handler.GetSaveds(c)
}

func (sh *SavedHandlers) GetASaved(c *gin.Context) {
	sh.Handler.GetASaved(c)
}

func (sh *SavedHandlers) GetSavedsByUserId(c *gin.Context) {
	sh.Handler.GetSavedsByUserId(c)
}

func (sh *SavedHandlers) CreateASaved(c *gin.Context) {
	sh.Handler.CreateASaved(c)
}

func (sh *SavedHandlers) UpdateASaved(c *gin.Context) {
	sh.Handler.UpdateASaved(c)
}

func (sh *SavedHandlers) DeleteASaved(c *gin.Context) {
	sh.Handler.DeleteASaved(c)
}

func (sh *SavedHandlers) DeleteASavedByDictionaryId(c *gin.Context) {
	sh.Handler.DeleteASavedByDictionaryId(c)
}
