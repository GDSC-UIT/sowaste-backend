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

type BadgeServices struct {
	Db *mongo.Client
}

func GetBadgeCollection(bs *BadgeServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.BadgeCollection, bs.Db)
}

func (bs *BadgeServices) GetBadges(c *gin.Context) {
	ctx := c.Request.Context()
	var badges []model.Badge
	cursor, err := GetBadgeCollection(bs).Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &badges); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all badges"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(badges, responseMessage))
}

func (bs *BadgeServices) GetABadge(c *gin.Context) {
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

	var badge model.Badge

	err = GetBadgeCollection(bs).FindOne(ctx, filter).Decode(&badge)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get an badge"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(badge, responseMessage))
}

func (bs *BadgeServices) CreateBadge(c *gin.Context) {
	ctx := c.Request.Context()
	var badge model.Badge

	if err := c.ShouldBindJSON(&badge); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	badge.ID = primitive.NewObjectID()

	result, err := GetBadgeCollection(bs).InsertOne(ctx, badge)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create a badge"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "badge": badge}, responseMessage))
}

func (bs *BadgeServices) UpdateBadge(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var badge model.Badge

	err = GetBadgeCollection(bs).FindOne(ctx, filter).Decode(&badge)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&badge); err != nil {
		return
	}

	update := bson.M{"$set": badge}

	result, err := GetBadgeCollection(bs).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update a badge"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "badge": badge}, responseMessage))
}

func (bs *BadgeServices) DeleteBadge(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetBadgeCollection(bs).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a badge"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}
