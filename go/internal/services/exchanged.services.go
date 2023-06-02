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
			"as":           "reward",
		},
	}
	var lookupUser = bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "uid",
			"foreignField": "uid",
			"as":           "user",
		},
	}

	cursor, err := GetExchangedCollection(es).Aggregate(context.TODO(), []bson.M{
		lookupUser,
		lookupRewards,
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
			"as":           "reward",
		},
	}
	var match = bson.M{
		"$match": filter,
	}
	cursor, err := GetExchangedCollection(es).Aggregate(ctx, []bson.M{
		lookupRewards,
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

func (es *ExchangedServices) GetCurrentUserExchangeds(c *gin.Context) {
	ctx := c.Request.Context()
	user := GetCurrentUser(c)

	filter := bson.M{"uid": user.UID}

	var exchanged []model.Exchanged

	var lookupRewards = bson.M{
		"$lookup": bson.M{
			"from":         "rewards",
			"localField":   "reward_id",
			"foreignField": "_id",
			"as":           "reward",
		},
	}

	var match = bson.M{
		"$match": filter,
	}

	cursor, err := GetExchangedCollection(es).Aggregate(ctx, []bson.M{
		lookupRewards,
		match,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err = cursor.All(ctx, &exchanged); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all exchanged by current user"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(exchanged, responseMessage))
}

func (es *ExchangedServices) CreateAExchanged(c *gin.Context) {
	ctx := c.Request.Context()
	var exchanged model.Exchanged
	if err := c.ShouldBindJSON(&exchanged); err != nil {
		return
	}
	if exchanged.RewardID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, "Reward id is required")
		return
	}
	exchanged.ID = primitive.NewObjectID()
	user := GetCurrentUser(c)
	exchanged.UserID = user.UID
	//** Insert a exchanged to the database **//
	result, err := GetExchangedCollection(es).InsertOne(ctx, exchanged)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create an exchanged"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "exchanged": exchanged}, responseMessage))
}

func (es *ExchangedServices) RefundAExchanged(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}
	user := GetCurrentUser(c)

	var exchanged model.Exchanged

	err = GetExchangedCollection(es).FindOne(ctx, filter).Decode(&exchanged)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if exchanged.UserID != user.UID {
		c.JSON(http.StatusBadRequest, "You can not refund this exchanged")
		return
	}

	// delete the exchange
	_, err = GetExchangedCollection(es).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var reward model.Reward

	err = utils.GetDatabaseCollection(utils.DbCollectionConstant.RewardCollection, es.Db).FindOne(ctx, bson.M{"_id": exchanged.RewardID}).Decode(&reward)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// refund the point for user

	user.RewardPoint += reward.Point

	_, err = utils.GetDatabaseCollection(utils.DbCollectionConstant.UserCollection, es.Db).UpdateOne(ctx, bson.M{"uid": user.UID}, bson.M{"$set": user})

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// update the reward

	responseMessage := "Successfully refund an exchange"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"user": user}, responseMessage))

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
