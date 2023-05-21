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

type BadgeCollectionServices struct {
	Db *mongo.Client
}

func GetBadgeCollectionCollection(bs *BadgeCollectionServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.Badge_CollectionCollection, bs.Db)
}

func (bs *BadgeCollectionServices) GetBadgeCollections(c *gin.Context) {
	ctx := c.Request.Context()
	var badges []model.BadgeCollection
	var projectBadgeCollections = bson.M{}
	var lookupBadge = bson.M{
		"$lookup": bson.M{
			"from":         "badges",
			"localField":   "badge_id",
			"foreignField": "_id",
			"as":           "badges",
		},
	}
	cursor, err := GetBadgeCollectionCollection(bs).Aggregate(context.TODO(), []bson.M{
		lookupBadge,
		projectBadgeCollections,
	})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &badges); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all badge collection"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(badges, responseMessage))
}

func (bs *BadgeCollectionServices) GetABadgeCollection(c *gin.Context) {
	ctx := c.Request.Context()
	//** Get param of the request uri **//
	param := c.Param("id")
	//** Convert id to mongodb object id **//
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var lookupBadge = bson.M{
		"$lookup": bson.M{
			"from":         "badges",
			"localField":   "badge_id",
			"foreignField": "_id",
			"as":           "badges",
		},
	}

	filter := bson.M{"_id": id}

	var match = bson.M{"$match": filter}

	var badge model.BadgeCollection

	curosr, err := GetBadgeCollectionCollection(bs).Aggregate(ctx, []bson.M{
		lookupBadge,
		match,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err = curosr.All(ctx, &badge); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get a badge collection"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(badge, responseMessage))
}

func (bs *BadgeCollectionServices) GetBadgeCollectionByUserId(c *gin.Context) {
	ctx := c.Request.Context()
	//** Get param of the request uri **//
	param := c.Param("id")
	//** Convert id to mongodb object id **//
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"user_id": id}

	var lookupBadge = bson.M{
		"$lookup": bson.M{
			"from":         "badges",
			"localField":   "badge_id",
			"foreignField": "_id",
			"as":           "badges",
		},
	}

	var match = bson.M{"$match": filter}

	var badge model.BadgeCollection

	cursor, err := GetBadgeCollectionCollection(bs).Aggregate(ctx, []bson.M{
		lookupBadge,
		match,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err = cursor.All(ctx, &badge); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get a badge collection"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(badge, responseMessage))
}

func (bs *BadgeCollectionServices) CreateBadgeCollection(c *gin.Context) {
	ctx := c.Request.Context()
	var badge model.BadgeCollection

	if err := c.ShouldBindJSON(&badge); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if badge.UserID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, "User id is required")
		return
	}

	badge.ID = primitive.NewObjectID()

	result, err := GetBadgeCollectionCollection(bs).InsertOne(ctx, badge)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create a badge"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "badge_collection": badge}, responseMessage))
}

func (bs *BadgeCollectionServices) UpdateBadgeCollection(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var badge model.BadgeCollection

	err = GetBadgeCollectionCollection(bs).FindOne(ctx, filter).Decode(&badge)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&badge); err != nil {
		return
	}

	update := bson.M{"$set": badge}

	result, err := GetBadgeCollectionCollection(bs).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update a badge"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "badge_collection": badge}, responseMessage))
}

func (bs *BadgeCollectionServices) DeleteBadgeCollection(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetBadgeCollectionCollection(bs).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a badge"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}
