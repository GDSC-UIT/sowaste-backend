package handlers

import (
	"github.com/GDSC-UIT/sowaste-backend/go/internal/services"
	"github.com/gin-gonic/gin"
)

type ArticleHandlers struct {
	Handler services.ArticleServices
}

func (ah *ArticleHandlers) GetArticles(c *gin.Context) {
	ah.Handler.GetArticles(c)
}

func (ah *ArticleHandlers) GetAnArticle(c *gin.Context) {
	ah.Handler.GetAnArticle(c)
}

func (ah *ArticleHandlers) CreateAnArticle(c *gin.Context) {
	ah.Handler.CreateArticle(c)
}

func (ah *ArticleHandlers) UpdateAnArticle(c *gin.Context) {
	ah.Handler.UpdateArticle(c)
}

func (ah *ArticleHandlers) DeleteAnArticle(c *gin.Context) {
	ah.Handler.DeleteArticle(c)
}
