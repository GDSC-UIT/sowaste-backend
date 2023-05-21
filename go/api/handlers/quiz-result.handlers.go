package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type QuizResultHandlers struct {
	Handler services.QuizResultServices
}

func (qrh *QuizResultHandlers) GetQuizResults(c *gin.Context) {
	qrh.Handler.GetQuizResults(c)
}

func (qrh *QuizResultHandlers) GetAQuizResult(c *gin.Context) {
	qrh.Handler.GetAQuizResult(c)
}

func (qrh *QuizResultHandlers) GetQuizResultsByUserId(c *gin.Context) {
	qrh.Handler.GetQuizResultsByUserId(c)
}

func (qrh *QuizResultHandlers) CreateAQuizResult(c *gin.Context) {
	qrh.Handler.CreateAQuizResult(c)
}

func (qrh *QuizResultHandlers) UpdateAQuizResult(c *gin.Context) {
	qrh.Handler.UpdateAQuizResult(c)
}

func (qrh *QuizResultHandlers) DeleteAQuizResult(c *gin.Context) {
	qrh.Handler.DeleteAQuizResult(c)
}
