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

type DIYServices struct {
	Db *mongo.Client
}

func GetDIYCollection(as *DIYServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.DIYCollection, as.Db)
}

func (as *DIYServices) GetDIYs(c *gin.Context) {
	ctx := c.Request.Context()
	var diys []model.DIY
	var projectDIYs = bson.M{
		"$project": bson.M{
			"description": 0,
		},
	}
	cursor, err := GetDIYCollection(as).Aggregate(context.TODO(), []bson.M{
		projectDIYs,
	})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &diys); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all diys"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(diys, responseMessage))
}

func (as *DIYServices) GetAnDIY(c *gin.Context) {
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

	var diy model.DIY

	err = GetDIYCollection(as).FindOne(ctx, filter).Decode(&diy)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get an diy"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(diy, responseMessage))
}

func (as *DIYServices) CreateDIY(c *gin.Context) {
	ctx := c.Request.Context()
	var diy model.DIY

	if err := c.ShouldBindJSON(&diy); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	diy.ID = primitive.NewObjectID()
	diy.CreatedAt = utils.GetCurrentTime()

	result, err := GetDIYCollection(as).InsertOne(ctx, diy)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create a diy"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "diy": diy}, responseMessage))
}

func (as *DIYServices) UpdateDIY(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var diy model.DIY

	err = GetDIYCollection(as).FindOne(ctx, filter).Decode(&diy)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&diy); err != nil {
		return
	}

	update := bson.M{"$set": diy}

	result, err := GetDIYCollection(as).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update a diy"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "diy": diy}, responseMessage))
}

func (as *DIYServices) DeleteDIY(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetDIYCollection(as).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a diy"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}
