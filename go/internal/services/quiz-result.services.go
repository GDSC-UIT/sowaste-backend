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

type QuizResultServices struct {
	Db *mongo.Client
}

func GetQuizResultCollection(qs *QuizResultServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.QuizResultCollection, qs.Db)
}

func (qs *QuizResultServices) GetQuizResults(c *gin.Context) {
	ctx := c.Request.Context()
	var quizResult []model.QuizResult

	var lookupDictionaries = bson.M{
		"$lookup": bson.M{
			"from":         "dictionaries",
			"localField":   "dictionary_id",
			"foreignField": "_id",
			"as":           "dictionaries",
		},
	}
	var lookupUser = bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "user_id",
			"foreignField": "_id",
			"as":           "user",
		},
	}

	var projectDictionaries = bson.M{
		"$project": bson.M{
			"dictionaries": bson.M{
				"display_image":     0,
				"description":       0,
				"short_description": 0,
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

	cursor, err := GetQuizResultCollection(qs).Aggregate(context.TODO(), []bson.M{
		lookupUser,
		lookupDictionaries,
		projectDictionaries,
		match,
		// project,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &quizResult); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	responseMessage := "Successfully get all quiz results"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(quizResult, responseMessage))
}

func (qs *QuizResultServices) GetAQuizResult(c *gin.Context) {
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

	var quizResult []model.QuizResult

	//** Get a lesson with dictionary populated **//
	var lookupDictionaries = bson.M{
		"$lookup": bson.M{
			"from":         "dictionaries",
			"localField":   "dictionary_id",
			"foreignField": "_id",
			"as":           "dictionaries",
		},
	}
	var lookupUser = bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "user_id",
			"foreignField": "_id",
			"as":           "user",
		},
	}
	var project = bson.M{
		"$project": bson.M{
			"dictionaries": bson.M{
				"display_image":     0,
				"description":       0,
				"short_description": 0,
			},
		},
	}
	var match = bson.M{
		"$match": filter,
	}
	cursor, err := GetQuizResultCollection(qs).Aggregate(ctx, []bson.M{
		lookupDictionaries,
		lookupUser,
		project,
		match,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &quizResult); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get a quiz result"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(quizResult[0], responseMessage))
}

func (qs *QuizResultServices) GetQuizResultsByUserId(c *gin.Context) {
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

	var quizResult []model.QuizResult

	var lookupDictionaries = bson.M{
		"$lookup": bson.M{
			"from":         "dictionaries",
			"localField":   "dictionary_id",
			"foreignField": "_id",
			"as":           "dictionaries",
		},
	}
	var lookupUser = bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "user_id",
			"foreignField": "_id",
			"as":           "user",
		},
	}
	var project = bson.M{
		"$project": bson.M{
			"dictionaries": bson.M{
				"display_image":     0,
				"description":       0,
				"short_description": 0,
			},
		},
	}
	var match = bson.M{
		"$match": filter,
	}
	cursor, err := GetQuizResultCollection(qs).Aggregate(ctx, []bson.M{
		lookupDictionaries,
		lookupUser,
		project,
		match,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &quizResult); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get quiz results by user id"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(quizResult, responseMessage))
}

func (qs *QuizResultServices) CreateAQuizResult(c *gin.Context) {
	ctx := c.Request.Context()
	var quizResult model.QuizResult
	if err := c.ShouldBindJSON(&quizResult); err != nil {
		return
	}
	if quizResult.DictionaryID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, "Dictionary id is required")
		return
	}
	if quizResult.UserID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, "User id is required")
		return
	}
	quizResult.ID = primitive.NewObjectID()
	//** Insert a quizResult to the database **//
	result, err := GetQuizResultCollection(qs).InsertOne(ctx, quizResult)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create a quiz result"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "quiz_result": quizResult}, responseMessage))
}

func (qs *QuizResultServices) UpdateAQuizResult(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var quizResult model.QuizResult

	err = GetQuizResultCollection(qs).FindOne(ctx, filter).Decode(&quizResult)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&quizResult); err != nil {
		return
	}

	update := bson.M{"$set": quizResult}

	result, err := GetQuizResultCollection(qs).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update a quiz result"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "quiz_result": quizResult}, responseMessage))
}

func (qs *QuizResultServices) DeleteAQuizResult(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetQuizResultCollection(qs).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a quiz result"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))

}
