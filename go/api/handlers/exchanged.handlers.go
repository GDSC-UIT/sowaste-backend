package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type ExchangedHandlers struct {
	Handler services.ExchangedServices
}

func (eh *ExchangedHandlers) GetExchanges(c *gin.Context) {
	eh.Handler.GetExchangeds(c)
}

func (eh *ExchangedHandlers) GetAExchange(c *gin.Context) {
	eh.Handler.GetAnExchanged(c)
}

func (eh *ExchangedHandlers) GetExchangedsByUserId(c *gin.Context) {
	eh.Handler.GetExchangedsByUserId(c)
}

func (eh *ExchangedHandlers) GetCurrentUserExchangeds(c *gin.Context) {
	eh.Handler.GetCurrentUserExchangeds(c)
}

func (eh *ExchangedHandlers) CreateAExchange(c *gin.Context) {
	eh.Handler.CreateAExchanged(c)
}

func (eh *ExchangedHandlers) UpdateAExchange(c *gin.Context) {
	eh.Handler.UpdateAExchanged(c)
}

func (eh *ExchangedHandlers) DeleteAExchange(c *gin.Context) {
	eh.Handler.DeleteAExchanged(c)
}
