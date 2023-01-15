package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type QuizHandlers struct {
	Handler services.QuizServices
}

func (qh *QuizHandlers) GetQuizzes(c *gin.Context) {

	//* called from quiz.services.go **//
	/*
	* @params c; type *gin.Context;
	* @return => void
	 */
	qh.Handler.GetQuizzes(c)
}

func (qh *QuizHandlers) GetAQuiz(c *gin.Context) {
	qh.Handler.GetAQuiz(c)
}

func (qh *QuizHandlers) CreateAQuiz(c *gin.Context) {
	qh.Handler.CreateAQuiz(c)
}

func (qh *QuizHandlers) UpdateAQuiz(c *gin.Context) {
	qh.Handler.UpdateAQuiz(c)
}

func (qh *QuizHandlers) DeleteAQuiz(c *gin.Context) {
	qh.Handler.DeleteAQuiz(c)
}
