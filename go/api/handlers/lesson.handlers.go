package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type LessonHandlers struct {
	Handler services.LessonServices
}

func (lh *LessonHandlers) GetLessons(c *gin.Context) {

	//* called from lesson.services.go **//
	/*
	* @params c; type *gin.Context;
	* @return => void
	 */
	lh.Handler.GetLessons(c)
}

func (lh *LessonHandlers) GetALesson(c *gin.Context) {
	lh.Handler.GetALesson(c)
}

func (lh *LessonHandlers) CreateALesson(c *gin.Context) {
	lh.Handler.CreateALesson(c)
}

func (lh *LessonHandlers) UpdateALesson(c *gin.Context) {
	lh.Handler.UpdateALesson(c)
}

func (lh *LessonHandlers) DeleteALesson(c *gin.Context) {
	lh.Handler.DeleteALesson(c)
}
