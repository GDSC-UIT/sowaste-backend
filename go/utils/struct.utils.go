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
}

type CacheConstants struct {
	Dictionaries string
	Dictionary   string
	Lessons      string
	Lesson       string
	Questions    string
	Question     string
	Options      string
	Option       string
	BottleShops  string
	BottleShop   string
	Articles     string
	Article      string
}
