package utils

type Response struct {
	Data    interface{} `json:"data" bson:"data"`
	Message string      `json:"message" bson:"message"`
}

type DbCollectionConstants struct {
	DictionaryCollection string
	LessonCollection     string
	QuestionCollection   string
	OptionCollection     string
	BottleShopCollection string
	ArticleCollection    string
	CategoryCollection   string
	UserCollection       string
	SavedCollection      string
	DIYCollection        string
	RewardCollection     string
}
