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

type DictionaryServices struct {
	Db *mongo.Client
}

func GetDictionaryCollection(ds *DictionaryServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.DictionaryCollection, ds.Db)
}

func (ds *DictionaryServices) GetDictionaries(c *gin.Context) {
	ctx := c.Request.Context()
	var dictionaries []model.Dictionary

	/*
		! Example of the result of the aggregation
		TODO https://vidler.app/blog/data/populate-golang-relationship-field-using-mongodb-aggregate-and-lookup/
	*/
	// var lookupQuizzes = bson.M{"$lookup": bson.M{
	// 	"from":         "questions",     //** collection name **//
	// 	"localField":   "_id",           //** field in the input documents **//
	// 	"foreignField": "dictionary_id", //** field in the documents of the "from" collection **//
	// 	"as":           "questions",     //** output array field **//
	// }}
	// var projectQuizzes = bson.M{"$project": bson.M{
	// 	"questions": bson.M{
	// 		"options": 0,
	// 	},
	// 	"description": 0,
	// }}
	var projectDictionaries = bson.M{"$project": bson.M{
		"description": 0,
	}}

	cursor, err := GetDictionaryCollection(ds).Aggregate(context.TODO(), []bson.M{
		projectDictionaries,
		// lookupQuizzes,
		// projectQuizzes,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &dictionaries); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
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
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	filter := bson.M{"_id": id}

	var dictionary []model.Dictionary

	// get the dictionary with the lessons and quizzes inside it
	var match = bson.M{"$match": filter}
	// var lookupQuizzes = bson.M{"$lookup": bson.M{
	// 	"from":         "questions",     //** collection name **//
	// 	"localField":   "_id",           //** field in the input documents **//
	// 	"foreignField": "dictionary_id", //** field in the documents of the "from" collection **//
	// 	"as":           "questions",     //** output array field **//
	// }}
	// var project = bson.M{"$project": bson.M{
	// 	"questions": bson.M{
	// 		"options": 0,
	// 		"title":   0,
	// 		"point":   0,
	// 	},
	// }}
	cursor, err := GetDictionaryCollection(ds).Aggregate(context.TODO(), []bson.M{
		match,
		// lookupQuizzes,
		// project,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err = cursor.All(ctx, &dictionary); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	responseMessage := "Successfully get the dictionary with id " + param
	c.JSON(http.StatusOK, utils.SuccessfulResponse(dictionary[0], responseMessage))
}

func (ds *DictionaryServices) CreateADictionary(c *gin.Context) {
	ctx := c.Request.Context()

	var dictionary model.Dictionary
	if err := c.ShouldBindJSON(&dictionary); err != nil {
		return
	}
	dictionary.ID = primitive.NewObjectID()
	dictionary.Questions = []model.Question{}

	result, err := GetDictionaryCollection(ds).InsertOne(ctx, dictionary)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully created a dictionary"
	c.JSON(http.StatusCreated, utils.SuccessfulResponse(bson.M{"result": result, "dictionary": dictionary}, responseMessage))
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

	err = GetDictionaryCollection(ds).FindOne(ctx, filter).Decode(&dictionary)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&dictionary); err != nil {
		return
	}

	fmt.Println(dictionary)

	update := bson.M{"$set": dictionary}

	result, err := GetDictionaryCollection(ds).UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully updated the dictionary with id " + param
	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "dictionary": dictionary}, responseMessage))
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
	deleteFilter := bson.M{"dictionary_id": id}

	// delete quizzes
	_, err = utils.GetDatabaseCollection(utils.DbCollectionConstant.QuestionCollection, ds.Db).DeleteMany(ctx, deleteFilter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = utils.GetDatabaseCollection(utils.DbCollectionConstant.OptionCollection, ds.Db).DeleteMany(ctx, deleteFilter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// delete the dictionary
	_, err = GetDictionaryCollection(ds).DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully deleted the dictionary with id " + param

	c.JSON(http.StatusOK, utils.SuccessfulResponse(nil, responseMessage))

}
