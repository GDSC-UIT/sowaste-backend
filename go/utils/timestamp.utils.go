package utils

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCurrentTime() primitive.DateTime {
	return primitive.NewDateTimeFromTime(time.Now())
}