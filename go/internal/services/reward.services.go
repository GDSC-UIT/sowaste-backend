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

type RewardServices struct {
	Db *mongo.Client
}

func GetRewardCollection(as *RewardServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.RewardCollection, as.Db)
}

func (as *RewardServices) GetRewards(c *gin.Context) {
	ctx := c.Request.Context()
	var rewards []model.Reward
	var projectRewards = bson.M{}
	cursor, err := GetRewardCollection(as).Aggregate(context.TODO(), []bson.M{
		projectRewards,
	})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &rewards); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all rewards"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(rewards, responseMessage))
}

func (as *RewardServices) GetAReward(c *gin.Context) {
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

	var reward model.Reward

	err = GetRewardCollection(as).FindOne(ctx, filter).Decode(&reward)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get an reward"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(reward, responseMessage))
}

func (as *RewardServices) GetUserRewards(c *gin.Context) {
	ctx := c.Request.Context()

	user := GetCurrentUser(c)

	uid := user.UID

	filter := bson.M{"uid": uid}

	var rewards []model.Reward

	cursor, err := GetRewardCollection(as).Find(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err = cursor.All(ctx, &rewards); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get user rewards with uid " + uid

	c.JSON(http.StatusOK, utils.SuccessfulResponse(rewards, responseMessage))
}

func (as *RewardServices) CreateReward(c *gin.Context) {
	ctx := c.Request.Context()
	var reward model.Reward

	if err := c.ShouldBindJSON(&reward); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	reward.ID = primitive.NewObjectID()

	result, err := GetRewardCollection(as).InsertOne(ctx, reward)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create a reward"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "reward": reward}, responseMessage))
}

func (as *RewardServices) UpdateReward(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var reward model.Reward

	err = GetRewardCollection(as).FindOne(ctx, filter).Decode(&reward)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&reward); err != nil {
		return
	}

	update := bson.M{"$set": reward}

	result, err := GetRewardCollection(as).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update a reward"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "reward": reward}, responseMessage))
}

func (as *RewardServices) DeleteReward(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetRewardCollection(as).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a reward"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}
