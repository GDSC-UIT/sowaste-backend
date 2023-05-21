package services

import (
	"net/http"

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
	//** Get param of the request uri **//
	param := c.Param("id")
	//** Convert id to mongodb object id **//
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

	responseMessage := "Successfully get a user"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(user, responseMessage))
}

func (as *UserServices) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user.ID = primitive.NewObjectID()

	result, err := GetUserCollection(as).InsertOne(ctx, user)
	if err != nil {
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