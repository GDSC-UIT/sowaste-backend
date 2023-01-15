package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type DictionaryHandlers struct {
	Handler services.DictionaryServices
}

func (dh *DictionaryHandlers) GetDictionaries(c *gin.Context) {

	//* called from dictionary.services.go **//
	/*
	* @params c; type *gin.Context;
	* @return => void
	 */
	dh.Handler.GetDictionaries(c)
}

func (dh *DictionaryHandlers) GetADictionary(c *gin.Context) {
	dh.Handler.GetADictionary(c)
}

func (dh *DictionaryHandlers) CreateADictionary(c *gin.Context) {
	dh.Handler.CreateADictionary(c)
}

func (dh *DictionaryHandlers) UpdateADictionary(c *gin.Context) {
	dh.Handler.UpdateADictionary(c)
}

func (dh *DictionaryHandlers) DeleteADictionary(c *gin.Context) {
	dh.Handler.DeleteADictionary(c)
}
