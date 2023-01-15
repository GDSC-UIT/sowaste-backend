package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type QuestionHandlers struct {
	Handler services.QuestionServices
}

func (qh *QuestionHandlers) GetQuestions(c *gin.Context) {

	//* called from question.services.go **//
	/*
	* @params c; type *gin.Context;
	* @return => void
	 */
	qh.Handler.GetQuestions(c)
}

func (qh *QuestionHandlers) GetAQuestion(c *gin.Context) {
	qh.Handler.GetAQuestion(c)
}

func (qh *QuestionHandlers) CreateAQuestion(c *gin.Context) {
	qh.Handler.CreateAQuestion(c)
}

func (qh *QuestionHandlers) UpdateAQuestion(c *gin.Context) {
	qh.Handler.UpdateAQuestion(c)
}

func (qh *QuestionHandlers) DeleteAQuestion(c *gin.Context) {
	qh.Handler.DeleteAQuestion(c)
}
