package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GDSC-UIT/sowaste-backend/go/internal/model"
	"github.com/GDSC-UIT/sowaste-backend/go/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExchangedServices struct {
	Db *mongo.Client
}

func GetExchangedCollection(es *ExchangedServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.ExchangedCollection, es.Db)
}

func (es *ExchangedServices) GetExchangeds(c *gin.Context) {
	ctx := c.Request.Context()
	var exchanged []model.Exchanged

	var lookupRewards = bson.M{
		"$lookup": bson.M{
			"from":         "rewards",
			"localField":   "reward_id",
			"foreignField": "_id",
			"as":           "rewards",
		},
	}
	var lookupUser = bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "_id",
			"foreignField": "user_id",
			"as":           "user",
		},
	}

	var projectDictionaries = bson.M{
		"$project": bson.M{},
	}

	var match = bson.M{
		"$match": bson.M{
			"dictionary_id": bson.M{
				"$exists": true,
			},
		},
	}

	cursor, err := GetExchangedCollection(es).Aggregate(context.TODO(), []bson.M{
		lookupUser,
		lookupRewards,
		projectDictionaries,
		match,
		// project,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &exchanged); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	responseMessage := "Successfully get all exchangeds"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(exchanged, responseMessage))
}

func (es *ExchangedServices) GetAnExchanged(c *gin.Context) {
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

	var exchanged []model.Exchanged

	//** Get a lesson with dictionary populated **//
	var lookupRewards = bson.M{
		"$lookup": bson.M{
			"from":         "rewards",
			"localField":   "reward_id",
			"foreignField": "_id",
			"as":           "rewards",
		},
	}
	var lookupUser = bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "_id",
			"foreignField": "user_id",
			"as":           "user",
		},
	}
	var project = bson.M{}
	var match = bson.M{
		"$match": filter,
	}
	cursor, err := GetExchangedCollection(es).Aggregate(ctx, []bson.M{
		lookupRewards,
		lookupUser,
		project,
		match,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &exchanged); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get an exchanged"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(exchanged[0], responseMessage))
}

func (es *ExchangedServices) GetExchangedsByUserId(c *gin.Context) {
	ctx := c.Request.Context()
	//** Get param of the request uri **//
	param := c.Param("user_id")
	//** Convert id to mongodb object id **//
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"user_id": id}

	var exchanged []model.Exchanged

	var lookupRewards = bson.M{
		"$lookup": bson.M{
			"from":         "rewards",
			"localField":   "reward_id",
			"foreignField": "_id",
			"as":           "rewards",
		},
	}
	var lookupUser = bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "_id",
			"foreignField": "user_id",
			"as":           "user",
		},
	}
	var project = bson.M{}
	var match = bson.M{
		"$match": filter,
	}
	cursor, err := GetExchangedCollection(es).Aggregate(ctx, []bson.M{
		lookupRewards,
		lookupUser,
		project,
		match,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &exchanged); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all exchanged"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(exchanged, responseMessage))
}

func (es *ExchangedServices) CreateAExchanged(c *gin.Context) {
	ctx := c.Request.Context()
	var exchanged model.Exchanged
	if err := c.ShouldBindJSON(&exchanged); err != nil {
		return
	}
	if exchanged.UserID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, "User id is required")
		return
	}
	if exchanged.RewardID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, "Reward id is required")
		return
	}
	exchanged.ID = primitive.NewObjectID()
	//** Insert a exchanged to the database **//
	result, err := GetExchangedCollection(es).InsertOne(ctx, exchanged)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create an exchanged"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "exchanged": exchanged}, responseMessage))
}

func (es *ExchangedServices) UpdateAExchanged(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var exchanged model.Exchanged

	err = GetExchangedCollection(es).FindOne(ctx, filter).Decode(&exchanged)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&exchanged); err != nil {
		return
	}

	update := bson.M{"$set": exchanged}

	result, err := GetExchangedCollection(es).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update a exchanged"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "exchanged": exchanged}, responseMessage))
}

func (es *ExchangedServices) DeleteAExchanged(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetExchangedCollection(es).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete an exchanged"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))

}
