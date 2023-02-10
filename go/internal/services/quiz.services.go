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

	var lookupDictionaries = bson.M{
		"$lookup": bson.M{
			"from":         "dictionaries",
			"localField":   "dictionary_id",
			"foreignField": "_id",
			"as":           "dictionaries",
		},
	}
	var lookupQuestions = bson.M{
		"$lookup": bson.M{
			"from":         "questions",
			"localField":   "_id",
			"foreignField": "quiz_id",
			"as":           "questions",
		},
	}

	// var project = bson.M{
	// 	"$project": bson.M{
	// 		"dictionaries": bson.M{
	// 			"quizzes": 0,
	// 			"lessons": 0,
	// 		},
	// 	},
	// }

	var projectDictionaries = bson.M{
		"$project": bson.M{
			"dictionaries": bson.M{
				"quizzes": 0,
				"lessons": 0,
			},
		},
	}

	var match = bson.M{
		"$match": bson.M{
			"dictionary_id": bson.M{
				"$exists": true,
			},
		},
	}

	cursor, err := GetQuizCollection(qs).Aggregate(context.TODO(), []bson.M{
		lookupQuestions,
		lookupDictionaries,
		projectDictionaries,
		match,
		// project,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &quizzes); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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

	var quiz []model.Quiz

	//** Get a lesson with dictionary populated **//
	var lookupDictionaries = bson.M{
		"$lookup": bson.M{
			"from":         "dictionaries",
			"localField":   "dictionary_id",
			"foreignField": "_id",
			"as":           "dictionaries",
		},
	}
	var lookupQuestions = bson.M{
		"$lookup": bson.M{
			"from":         "questions",
			"localField":   "_id",
			"foreignField": "quiz_id",
			"as":           "questions",
		},
	}
	var project = bson.M{
		"$project": bson.M{
			"dictionaries": bson.M{
				"lessons": 0,
				"quizzes": 0,
			},
		},
	}
	var match = bson.M{
		"$match": filter,
	}
	cursor, err := GetQuizCollection(qs).Aggregate(ctx, []bson.M{
		lookupDictionaries,
		lookupQuestions,
		project,
		match,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &quiz); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get a quiz"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(quiz[0], responseMessage))
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

	updateDictionary, err := utils.GetDatabaseCollection(utils.DbCollectionConstant.DictionaryCollection, qs.Db).UpdateOne(ctx, filter, bson.M{"$push": bson.M{"quizzes": quiz.ID}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updateDictionary.MatchedCount)

	responseMessage := "Successfully create a quiz"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "quiz": quiz}, responseMessage))
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

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "quiz": quiz}, responseMessage))
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

	pull := bson.M{"$pull": bson.M{"quizzes": id}}
	_, err = utils.GetDatabaseCollection(utils.DbCollectionConstant.DictionaryCollection, qs.Db).UpdateOne(ctx, filter, pull)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = utils.GetDatabaseCollection(utils.DbCollectionConstant.QuestionCollection, qs.Db).DeleteMany(ctx, bson.M{"quiz_id": id})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result, err := GetQuizCollection(qs).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a quiz"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))

}
