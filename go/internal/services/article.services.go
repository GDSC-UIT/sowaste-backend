package services

import (
	"context"
	"net/http"

	"github.com/GDSC-UIT/sowaste-backend/go/internal/model"
	"github.com/GDSC-UIT/sowaste-backend/go/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleServices struct {
	Db *mongo.Client
}

func GetArticleCollection(as *ArticleServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.ArticleCollection, as.Db)
}

func (as *ArticleServices) GetArticles(c *gin.Context) {
	ctx := c.Request.Context()
	var articles []model.Article
	cursor, err := GetArticleCollection(as).Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &articles); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all articles"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(articles, responseMessage))
}

func (as *ArticleServices) GetAnArticle(c *gin.Context) {
	ctx := c.Request.Context()
	//** Get param of the request uri **//
	param := c.Param("id")
	//** Convert id to mongodb object id **//
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var article model.Article

	err = GetArticleCollection(as).FindOne(ctx, filter).Decode(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get an article"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(article, responseMessage))
}

func (as *ArticleServices) CreateArticle(c *gin.Context) {
	ctx := c.Request.Context()
	var article model.Article

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	article.ID = primitive.NewObjectID()
	article.CreatedAt = utils.GetCurrentTime()

	result, err := GetArticleCollection(as).InsertOne(ctx, article)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create an article"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "article": article}, responseMessage))
}

func (as *ArticleServices) UpdateArticle(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var article model.Article

	err = GetArticleCollection(as).FindOne(ctx, filter).Decode(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&article); err != nil {
		return
	}

	update := bson.M{"$set": article}

	result, err := GetArticleCollection(as).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update an article"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "article": article}, responseMessage))
}

func (as *ArticleServices) DeleteArticle(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetArticleCollection(as).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete an article"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}
