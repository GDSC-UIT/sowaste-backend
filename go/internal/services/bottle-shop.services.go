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

type BottleShopServices struct {
	Db *mongo.Client
}

func GetBottleShopCollection(bss *BottleShopServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.BottleShopCollection, bss.Db)
}

func (bss *BottleShopServices) GetBottleShops(c *gin.Context) {
	ctx := c.Request.Context()
	var bottleShops []model.BottleShop
	cursor, err := GetBottleShopCollection(bss).Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &bottleShops); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all quizzes"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bottleShops, responseMessage))
}

func (bss *BottleShopServices) GetABottleShop(c *gin.Context) {
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

	var bottleShop model.BottleShop

	err = GetBottleShopCollection(bss).FindOne(ctx, filter).Decode(&bottleShop)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get a bottleShop"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bottleShop, responseMessage))
}

func (bss *BottleShopServices) CreateABottleShop(c *gin.Context) {
	ctx := c.Request.Context()
	var bottleShop model.BottleShop
	if err := c.ShouldBindJSON(&bottleShop); err != nil {
		return
	}

	bottleShop.ID = primitive.NewObjectID()
	//** Insert a quiz to the database **//
	result, err := GetBottleShopCollection(bss).InsertOne(ctx, bottleShop)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create a bottleShop"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}

func (bss *BottleShopServices) UpdateAbottleShop(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var bottleShop model.BottleShop

	err = GetBottleShopCollection(bss).FindOne(ctx, filter).Decode(&bottleShop)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&bottleShop); err != nil {
		return
	}

	update := bson.M{"$set": bottleShop}

	result, err := GetBottleShopCollection(bss).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update a bottleShop"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}

func (bss *BottleShopServices) DeleteAbottleShop(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetBottleShopCollection(bss).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a bottleShop"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))

}
