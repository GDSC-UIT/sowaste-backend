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

type CategoryServices struct {
	Db *mongo.Client
}

func GetCategoryCollection(cs *CategoryServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.CategoryCollection, cs.Db)
}

func (cs *CategoryServices) GetCategories(c *gin.Context) {
	ctx := c.Request.Context()
	var categories []model.Category

	cursor, err := GetCategoryCollection(cs).Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &categories); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all categories"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(categories, responseMessage))
}

func (cs *CategoryServices) GetAnCategory(c *gin.Context) {
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

	var category model.Category

	err = GetCategoryCollection(cs).FindOne(ctx, filter).Decode(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get an category"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(category, responseMessage))
}

func (cs *CategoryServices) CreateCategory(c *gin.Context) {
	ctx := c.Request.Context()
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if category.DictonaryID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, "Dictionary id is required")
		return
	}

	category.ID = primitive.NewObjectID()

	result, err := GetCategoryCollection(cs).InsertOne(ctx, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create an category"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "category": category}, responseMessage))
}

func (cs *CategoryServices) UpdateCategory(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var category model.Category

	err = GetCategoryCollection(cs).FindOne(ctx, filter).Decode(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		return
	}

	update := bson.M{"$set": category}

	result, err := GetCategoryCollection(cs).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update an category"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "category": category}, responseMessage))
}

func (cs *CategoryServices) DeleteCategory(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetCategoryCollection(cs).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete an category"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}
