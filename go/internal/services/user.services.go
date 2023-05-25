package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GDSC-UIT/sowaste-backend/go/internal/database"
	"github.com/GDSC-UIT/sowaste-backend/go/internal/model"
	"github.com/GDSC-UIT/sowaste-backend/go/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServices struct {
	Db *mongo.Client
}

func GetUserCollection(as *UserServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.UserCollection, as.Db)
}

func (as *UserServices) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	claims := c.MustGet("CLAIMS").(utils.UserClaims)

	filter := bson.M{"email": claims.Email}

	var user model.User

	err := GetUserCollection(as).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("User not found"))
		return
	}

	responseMessage := "Successfully get current user"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(user, responseMessage))
}

func (as *UserServices) GetUserById(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var user model.User

	err = GetUserCollection(as).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get user with id " + param

	c.JSON(http.StatusOK, utils.SuccessfulResponse(user, responseMessage))
}

func (as *UserServices) CheckIfUserNotExists(c *gin.Context) bool {
	claims := c.MustGet("CLAIMS").(utils.UserClaims)
	ctx := c.Request.Context()
	filter := bson.M{"email": claims.Email}
	var user model.User
	err := GetUserCollection(as).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return true
	}
	if user.Email == claims.Email {
		return false
	}
	return true
}

func (as *UserServices) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	claims := c.MustGet("CLAIMS").(utils.UserClaims)
	uid := c.Query("uid")
	if uid == "" {
		c.JSON(http.StatusBadRequest, "uid is required")
		return
	}
	if !as.CheckIfUserNotExists(c) {
		c.JSON(http.StatusOK, "User already exists and can log in")
		return
	}
	var user model.User
	user.FullName = claims.Name
	user.Email = claims.Email
	user.DisplayImage = claims.Picture
	user.RewardPoint = 0
	user.ID = primitive.NewObjectID()
	user.UID = uid
	result, err := GetUserCollection(as).InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create a user"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "user": user}, responseMessage))
}

func (as *UserServices) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var user model.User

	err = GetUserCollection(as).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		return
	}

	update := bson.M{"$set": user}

	result, err := GetUserCollection(as).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update an user"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "user": user}, responseMessage))
}

func (as *UserServices) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetUserCollection(as).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a user"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}

func (us *UserServices) UpdateCurrentUserPoint(c *gin.Context) {
	ctx := c.Request.Context()
	user := GetCurrentUser(c)
	point := c.Query("point")
	if point == "" {
		c.JSON(http.StatusBadRequest, "Point is required")
		return
	}
	action := c.Query("action")
	if action == "" {
		c.JSON(http.StatusBadRequest, "Action is required")
		return
	}
	pointInt, err := strconv.Atoi(point)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if action == "decrease" {
		if user.RewardPoint < pointInt {
			c.JSON(http.StatusBadRequest, "Not enough point")
			return
		}
		user.RewardPoint = user.RewardPoint - pointInt
	} else if action == "increase" {
		user.RewardPoint = user.RewardPoint + pointInt
	}

	filter := bson.M{"_id": user.ID}

	update := bson.M{"$set": user}

	result, err := GetUserCollection(us).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully " + action + " point for user " + user.FullName

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "user": user}, responseMessage))
}

func (us *UserServices) CreateCurrentUserExchanges(c *gin.Context) {
	ctx := c.Request.Context()
	user := GetCurrentUser(c)

	rewardId := c.Query("reward_id")

	if rewardId == "" {
		c.JSON(http.StatusBadRequest, "Reward id is required")
		return
	}

	var exchange model.Exchanged

	exchange.ID = primitive.NewObjectID()
	rewardIdObject, err := primitive.ObjectIDFromHex(rewardId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	exchange.RewardID = rewardIdObject

	var reward model.Reward

	filter := bson.M{"_id": rewardIdObject}

	err = utils.GetDatabaseCollection(utils.DbCollectionConstant.RewardCollection, us.Db).FindOne(ctx, filter).Decode(&reward)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if reward.Point > user.RewardPoint {
		c.JSON(http.StatusBadRequest, "Not enough point")
		return
	}

	user.RewardPoint = user.RewardPoint - reward.Point

	filter = bson.M{"_id": user.ID}

	update := bson.M{"$set": user}

	_, err = GetUserCollection(us).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Error while update user point")
		return
	}

	exchange.UserID = user.UID

	result, err := utils.GetDatabaseCollection(utils.DbCollectionConstant.ExchangedCollection, us.Db).InsertOne(ctx, exchange)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create an exchange"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "exchange": exchange, "user": user}, responseMessage))

}

func (us *UserServices) GetCurrentUserBadges(c *gin.Context) {
	ctx := c.Request.Context()
	user := GetCurrentUser(c)

	filter := bson.M{"uid": user.UID}

	var badges []model.BadgeCollection

	cursor, err := utils.GetDatabaseCollection(utils.DbCollectionConstant.Badge_CollectionCollection, us.Db).Find(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err = cursor.All(ctx, &badges); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all badges for user " + user.FullName

	c.JSON(http.StatusOK, utils.SuccessfulResponse(badges, responseMessage))
}

func GetCurrentUser(c *gin.Context) model.User {
	ctx := c.Request.Context()
	claims := c.MustGet("CLAIMS").(utils.UserClaims)

	filter := bson.M{"email": claims.Email}

	var user model.User

	err := utils.GetDatabaseCollection(utils.DbCollectionConstant.UserCollection, database.Client.Source).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return model.User{}
	}

	return user
}
