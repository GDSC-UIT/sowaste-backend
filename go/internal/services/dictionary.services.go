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

type DictionaryServices struct {
	Db *mongo.Client
}

const dictionaryDbCollectionName = "dictionaries"

func getDictionaryCollection(ds *DictionaryServices) *mongo.Collection {
	return utils.GetDatabaseCollection(dictionaryDbCollectionName, ds.Db)
}

func (ds *DictionaryServices) GetDictionaries(c *gin.Context) {
	ctx := c.Request.Context()
	var dictionaries []model.Dictionary

	//** utils.GetDatabaseCollection **//
	//* called from collection.utils.go **//
	/*
	* @params dictionaryDbCollectionName; type string;
	* @params ds.Db; type *mongo.Client
	* @return => *mongo.Client
	 */
	cursor, err := getDictionaryCollection(ds).Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &dictionaries); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	responseMessage := "Successfully get all dictionaries"
	c.JSON(http.StatusOK, utils.SuccessfulResponse(dictionaries, responseMessage))
}

func (ds *DictionaryServices) GetADictionary(c *gin.Context) {
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

	var dictionary model.Dictionary
	err = getDictionaryCollection(ds).FindOne(ctx, filter).Decode(&dictionary)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	responseMessage := "Successfully get the dictionary with id " + param
	c.JSON(http.StatusOK, utils.SuccessfulResponse(dictionary, responseMessage))
}

func (ds *DictionaryServices) CreateADictionary(c *gin.Context) {
	ctx := c.Request.Context()

	var dictionary model.Dictionary
	if err := c.ShouldBindJSON(&dictionary); err != nil {
		return
	}
	_, err := getDictionaryCollection(ds).InsertOne(ctx, dictionary)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully created a dictionary"
	c.JSON(http.StatusCreated, utils.SuccessfulResponse(dictionary, responseMessage))
}

func (ds *DictionaryServices) UpdateADictionary(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var dictionary model.Dictionary
	if err := c.ShouldBindJSON(&dictionary); err != nil {
		return
	}

	update := bson.M{"$set": dictionary}

	_, err = getDictionaryCollection(ds).UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully updated the dictionary with id " + param
	c.JSON(http.StatusOK, utils.SuccessfulResponse(dictionary, responseMessage))
}

func (ds *DictionaryServices) DeleteADictionary(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	_, err = getDictionaryCollection(ds).DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully deleted the dictionary with id " + param

	c.JSON(http.StatusOK, utils.SuccessfulResponse(nil, responseMessage))
}
