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

type SavedServices struct {
	Db *mongo.Client
}

func GetSavedCollection(ss *SavedServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.SavedCollection, ss.Db)
}

func (ss *SavedServices) GetSaveds(c *gin.Context) {
	ctx := c.Request.Context()
	var saved []model.Saved

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
			"localField":   "uid",
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

	cursor, err := GetSavedCollection(ss).Aggregate(context.TODO(), []bson.M{
		lookupUser,
		lookupDictionaries,
		projectDictionaries,
		match,
		// project,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &saved); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	responseMessage := "Successfully get all saveds"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(saved, responseMessage))
}

func (ss *SavedServices) GetASaved(c *gin.Context) {
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

	var saved []model.Saved

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
			"localField":   "uid",
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
	cursor, err := GetSavedCollection(ss).Aggregate(ctx, []bson.M{
		lookupDictionaries,
		lookupUser,
		project,
		match,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &saved); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get a saved"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(saved[0], responseMessage))
}

func (ss *SavedServices) GetSavedsByUserId(c *gin.Context) {
	ctx := c.Request.Context()
	//** Get param of the request uri **//
	param := c.Param("uid")
	//** Convert id to mongodb object id **//
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"uid": id}

	var saved []model.Saved

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
			"localField":   "uid",
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
	cursor, err := GetSavedCollection(ss).Aggregate(ctx, []bson.M{
		lookupDictionaries,
		lookupUser,
		project,
		match,
	})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &saved); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get saveds by user id " + param

	c.JSON(http.StatusOK, utils.SuccessfulResponse(saved, responseMessage))
}

func (ss *SavedServices) GetSavedsOfUser(c *gin.Context) {
	ctx := c.Request.Context()

	user := GetCurrentUser(c)

	var saved []model.Saved

	filter := bson.M{"uid": user.UID}
	cursor, err := GetSavedCollection(ss).Find(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err = cursor.All(ctx, &saved); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get saveds of user " + user.UID

	c.JSON(http.StatusOK, utils.SuccessfulResponse(saved, responseMessage))

	// var saved []model.Saved

	// //** Get a lesson with dictionary populated **//
	// var lookupDictionaries = bson.M{
	// 	"$lookup": bson.M{
	// 		"from":         "dictionaries",
	// 		"localField":   "dictionary_id",
	// 		"foreignField": "_id",
	// 		"as":           "dictionaries",
	// 	},
	// }
	// var lookupUser = bson.M{
	// 	"$lookup": bson.M{
	// 		"from":         "users",
	// 		"localField":   "uid",
	// 		"foreignField": "_id",
	// 		"as":           "user",
	// 	},
	// }
	// var project = bson.M{
	// 	"$project": bson.M{
	// 		"dictionaries": bson.M{
	// 			"display_image":     0,
	// 			"description":       0,
	// 			"short_description": 0,
	// 		},
	// 	},
	// }
	// var match = bson.M{
	// 	"$match": filter,
	// }
	// cursor, err := GetSavedCollection(ss).Aggregate(ctx, []bson.M{
	// 	lookupDictionaries,
	// 	lookupUser,
	// 	project,
	// 	match,
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// if err = cursor.All(ctx, &saved); err != nil {
	// 	fmt.Println(err.Error())
	// 	c.JSON(http.StatusBadRequest, err)
	// 	return
	// }

	// responseMessage := "Successfully get saveds"

	// c.JSON(http.StatusOK, utils.SuccessfulResponse(saved, responseMessage))
}

func (ss *SavedServices) CreateASaved(c *gin.Context) {
	ctx := c.Request.Context()
	var saved model.Saved
	if err := c.ShouldBindJSON(&saved); err != nil {
		return
	}
	if saved.DictionaryID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, "Dictionary id is required")
		return
	}
	if saved.UserID == "" {
		c.JSON(http.StatusBadRequest, "User id is required")
		return
	}
	saved.ID = primitive.NewObjectID()
	//** Insert a saved to the database **//
	result, err := GetSavedCollection(ss).InsertOne(ctx, saved)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully create a saved"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "saved": saved}, responseMessage))
}

func (ss *SavedServices) UpdateASaved(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var saved model.Saved

	err = GetSavedCollection(ss).FindOne(ctx, filter).Decode(&saved)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&saved); err != nil {
		return
	}

	update := bson.M{"$set": saved}

	result, err := GetSavedCollection(ss).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update a saved"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "saved": saved}, responseMessage))
}

func (ss *SavedServices) DeleteASaved(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetSavedCollection(ss).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a saved"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))

}

func (ss *SavedServices) DeleteASavedByDictionaryId(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("dictionary_id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"dictionary_id": id}

	result, err := GetSavedCollection(ss).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete saveds by dictionary_id " + param

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))

}
