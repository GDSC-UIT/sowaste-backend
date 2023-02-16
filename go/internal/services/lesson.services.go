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

type LessonServices struct {
	Db *mongo.Client
}

func GetLessonCollection(ls *LessonServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.LessonCollection, ls.Db)
}

func (ls *LessonServices) GetLessons(c *gin.Context) {
	ctx := c.Request.Context()
	var lessons []model.Lesson

	//** Get all lessons with dictionary populated **//
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
				"lessons":   0,
				"questions": 0,
			},
		},
	}
	cursor, err := GetLessonCollection(ls).Aggregate(context.TODO(), []bson.M{
		lookup,
		project,
	})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &lessons); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all lessons"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(lessons, responseMessage))
}

func (ls *LessonServices) GetALesson(c *gin.Context) {
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

	var lesson []model.Lesson

	// err = GetLessonCollection(ls).FindOne(ctx, filter).Decode(&lesson)

	//** Get a lesson with dictionary populated **//
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
				"lessons":   0,
				"questions": 0,
			},
		},
	}
	var match = bson.M{
		"$match": filter,
	}
	cursor, err := GetLessonCollection(ls).Aggregate(ctx, []bson.M{
		lookup,
		project,
		match,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &lesson); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get a lesson"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(lesson[0], responseMessage))
}

func (ls *LessonServices) CreateALesson(c *gin.Context) {
	ctx := c.Request.Context()
	var lesson model.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		return
	}
	if lesson.DictionaryID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, "Dictionary id is required")
		return
	}
	lesson.ID = primitive.NewObjectID()
	//** Insert a lesson to the database **//
	result, err := GetLessonCollection(ls).InsertOne(ctx, lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	filter := bson.M{"_id": lesson.DictionaryID}

	updateDictionary, err := utils.GetDatabaseCollection(utils.DbCollectionConstant.DictionaryCollection, ls.Db).UpdateOne(ctx, filter, bson.M{"$push": bson.M{"lessons": lesson.ID}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updateDictionary.MatchedCount)

	responseMessage := "Successfully create a lesson"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "lesson": lesson}, responseMessage))
}

func (ls *LessonServices) UpdateALesson(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var lesson model.Lesson

	err = GetLessonCollection(ls).FindOne(ctx, filter).Decode(&lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&lesson); err != nil {
		return
	}

	update := bson.M{"$set": lesson}

	result, err := GetLessonCollection(ls).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update a lesson"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "lesson": lesson}, responseMessage))
}

func (ls *LessonServices) DeleteALesson(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	pull := bson.M{"$pull": bson.M{"lessons": id}}
	_, err = utils.GetDatabaseCollection(utils.DbCollectionConstant.DictionaryCollection, ls.Db).UpdateOne(ctx, filter, pull)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result, err := GetLessonCollection(ls).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a lesson"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))

}
