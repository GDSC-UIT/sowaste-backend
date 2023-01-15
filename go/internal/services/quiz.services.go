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

type QuizServices struct {
	Db *mongo.Client
}

func GetQuizCollection(qs *QuizServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.QuizCollection, qs.Db)
}

func (qs *QuizServices) GetQuizzes(c *gin.Context) {
	ctx := c.Request.Context()
	var quizzes []model.Quiz
	var lookup = bson.M{
		"$lookup": bson.M{
			"from":         "dictionaries",
			"localField":   "dictionary_id",
			"foreignField": "_id",
			"as":           "dictionaries",
		},
	}
	var project = bson.M{
		"$project": bson.M{
			"dictionaries": bson.M{
				"quizzes": 0,
				"lessons": 0,
			},
		},
	}
	cursor, err := GetQuizCollection(qs).Aggregate(context.TODO(), []bson.M{
		lookup,
		project,
	})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &quizzes); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all quizzes"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(quizzes, responseMessage))
}

func (qs *QuizServices) GetAQuiz(c *gin.Context) {
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

	var quiz model.Quiz

	err = GetQuizCollection(qs).FindOne(ctx, filter).Decode(&quiz)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get a quiz"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(quiz, responseMessage))
}

func (qs *QuizServices) CreateAQuiz(c *gin.Context) {
	ctx := c.Request.Context()
	var quiz model.Quiz
	if err := c.ShouldBindJSON(&quiz); err != nil {
		return
	}
	if quiz.DictionaryID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, "Dictionary id is required")
		return
	}
	quiz.ID = primitive.NewObjectID()
	quiz.Question = []model.Question{}
	//** Insert a quiz to the database **//
	result, err := GetQuizCollection(qs).InsertOne(ctx, quiz)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	filter := bson.M{"_id": quiz.DictionaryID}

	updateDictionary, err := utils.GetDatabaseCollection(utils.DbCollectionConstant.DictionaryCollection, qs.Db).UpdateOne(ctx, filter, bson.M{"$push": bson.M{"quizzes": quiz.DictionaryID}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updateDictionary.MatchedCount)

	responseMessage := "Successfully create a quiz"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}

func (qs *QuizServices) UpdateAQuiz(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var quiz model.Quiz

	err = GetQuizCollection(qs).FindOne(ctx, filter).Decode(&quiz)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&quiz); err != nil {
		return
	}

	update := bson.M{"$set": quiz}

	result, err := GetQuizCollection(qs).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update a quiz"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}

func (qs *QuizServices) DeleteAQuiz(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetQuizCollection(qs).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a quiz"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))

}
