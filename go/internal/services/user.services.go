package services

import (
	"fmt"
	"net/http"

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
		c.JSON(http.StatusBadRequest, err)
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
	fmt.Println(uid)
	c.Set("UUID", uid)
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
