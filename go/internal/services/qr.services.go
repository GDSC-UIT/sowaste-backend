package services

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/GDSC-UIT/sowaste-backend/go/internal/model"
	"github.com/GDSC-UIT/sowaste-backend/go/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QRServices struct {
	Db *mongo.Client
}

func GetQRCollection(qrs *QRServices) *mongo.Collection {
	return utils.GetDatabaseCollection(utils.DbCollectionConstant.QRCollection, qrs.Db)
}

func (qrs *QRServices) GetQRs(c *gin.Context) {
	ctx := c.Request.Context()
	var QRs []model.QR
	cursor, err := GetQRCollection(qrs).Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &QRs); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get all QRs"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(QRs, responseMessage))
}

func (qrs *QRServices) GetAQR(c *gin.Context) {
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

	var QR model.QR

	err = GetQRCollection(qrs).FindOne(ctx, filter).Decode(&QR)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully get an QR"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(QR, responseMessage))
}

func (qrs *QRServices) CreateQR(c *gin.Context) {
	ctx := c.Request.Context()
	var QR model.QR

	if err := c.ShouldBindJSON(&QR); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	QR.ID = primitive.NewObjectID()
	QR.UserIDs = []string{}
	QR.IssuedAt = primitive.NewDateTimeFromTime(time.Now())
	QR.ExpireAt = primitive.NewDateTimeFromTime(time.Now().Add(time.Minute * 5)) // 5 minutes
	_, err := GetQRCollection(qrs).InsertOne(ctx, QR)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := http.Get("https://api.qrserver.com/v1/create-qr-code/?size=512x512&data=" + QR.ID.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Error while generating QR code"))
		return
	} else {
		c.Writer.Header().Set("Content-Type", "image/png")
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		c.Writer.Write(body)

	}
	defer response.Body.Close()
}

func (qrs *QRServices) GenerateQRCode(c *gin.Context) {
	param := c.Query("id")
	if param == "" {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Missing id parameter"))
		return
	}

	var QR model.QR

	id, err := primitive.ObjectIDFromHex(param)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Invalid id parameter"))
		return
	}

	filter := bson.M{"_id": id}

	err = GetQRCollection(qrs).FindOne(c.Request.Context(), filter).Decode(&QR)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("QR not found"))
		return
	}

	response, err := http.Get("https://api.qrserver.com/v1/create-qr-code/?size=512x512&data=" + param)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Error while generating QR code"))
		return
	} else {
		c.Writer.Header().Set("Content-Type", "image/png")
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		c.Writer.Write(body)

	}
	defer response.Body.Close()
}

func (qrs *QRServices) ScanQR(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")

	if param == "" {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Missing id parameter"))
		return
	}

	user := GetCurrentUser(c)
	uid := user.UID

	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	filter := bson.M{"_id": id}

	var QR model.QR

	err = GetQRCollection(qrs).FindOne(ctx, filter).Decode(&QR)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if QR.UserIDs == nil {
		QR.UserIDs = []string{}
	}

	if utils.ContainsElement(QR.UserIDs, uid) {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("User already scanned this QR"))
		return
	} else {
		QR.UserIDs = append(QR.UserIDs, uid)

		update := bson.M{"$set": QR}

		result, err := GetQRCollection(qrs).UpdateOne(ctx, filter, update)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		user.RewardPoint += QR.Point
		update = bson.M{"$set": user}

		_, err = utils.GetDatabaseCollection(utils.DbCollectionConstant.UserCollection, qrs.Db).UpdateOne(ctx, bson.M{"uid": user.UID}, update)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		responseMessage := "Successfully update a QR"

		c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "QR": QR}, responseMessage))
	}

	c.JSON(http.StatusInternalServerError, utils.FailedResponse("Something went wrong"))

}

func (qrs *QRServices) UpdateQr(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	var QR model.QR

	err = GetQRCollection(qrs).FindOne(ctx, filter).Decode(&QR)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&QR); err != nil {
		return
	}

	update := bson.M{"$set": QR}

	result, err := GetQRCollection(qrs).UpdateOne(ctx, filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully update a QR"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(bson.M{"result": result, "QR": QR}, responseMessage))
}

func (qrs *QRServices) DeleteQR(c *gin.Context) {
	ctx := c.Request.Context()
	param := c.Param("id")
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	filter := bson.M{"_id": id}

	result, err := GetQRCollection(qrs).DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage := "Successfully delete a QR"

	c.JSON(http.StatusOK, utils.SuccessfulResponse(result, responseMessage))
}
