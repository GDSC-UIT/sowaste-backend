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

type QuestionServices struct {
	Db *mongo.Client
}

func GetQuestionCollection(qs *QuestionServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.QuestionCollection, qs.Db)
}

func (qs *QuestionServices) GetQuestions(c *gin.Context) {
	ctx := c.Request.Context()
	var questions []model.Question
	cursor, err := GetQuestionCollection(qs).Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &questions); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all quizzes"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(questions, responseMessage))
}

func (qs *QuestionServices) GetAQuestion(c *gin.Context) {
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

	var question model.Question

	err = GetQuestionCollection(qs).FindOne(ctx, filter).Decode(&question)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get a question"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(question, responseMessage))
}

func (qs *QuestionServices) CreateAQuestion(c *gin.Context) {
	ctx := c.Request.Context()
	var question model.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		return
	}
	if question.QuizID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, "Quiz id is required")
		return
	}

	if question.DictionaryID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, "Dictionary id is required")
		return
	}

	question.ID = primitive.NewObjectID()
	//** Insert a quiz to the database **//
	result, err := GetQuestionCollection(qs).InsertOne(ctx, question)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	filter := bson.M{"_id": question.QuizID}

	updateQuiz, err := utils.GetDatabaseCollection(utils.DbCollectionConstant.QuizCollection, qs.Db).UpdateOne(ctx, filter, bson.M{"$push": bson.M{"questions": question.ID}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updateQuiz.MatchedCount)

	responseMessage := "Successfully create a question"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}

func (qs *QuestionServices) UpdateAQuestion(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var question model.Question

	err = GetQuestionCollection(qs).FindOne(ctx, filter).Decode(&question)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&question); err != nil {
		return
	}

	update := bson.M{"$set": question}

	result, err := GetQuestionCollection(qs).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update a question"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}

func (qs *QuestionServices) DeleteAQuestion(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	pull := bson.M{"$pull": bson.M{"questions": id}}
	_, err = utils.GetDatabaseCollection(utils.DbCollectionConstant.QuizCollection, qs.Db).UpdateOne(ctx, filter, pull)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result, err := GetQuestionCollection(qs).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a question"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))

}
