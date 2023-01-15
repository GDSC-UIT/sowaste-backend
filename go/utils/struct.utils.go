package utils

type Response struct {
	Data    interface{} `json:"data" bson:"data"`
	Message string      `json:"message" bson:"message"`
}

type DbCollectionConstants struct {
	DictionaryCollection string
	LessonCollection     string
	QuizCollection       string
	BottleShopCollection string
	QuestionCollection   string
}
