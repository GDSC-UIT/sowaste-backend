package utils

type Response struct {
	Data    interface{} `json:"data" bson:"data"`
	Message string      `json:"message" bson:"message"`
}

type DbCollectionConstants struct {
	DictionaryCollection       string
	QuestionCollection         string
	OptionCollection           string
	ArticleCollection          string
	CategoryCollection         string
	UserCollection             string
	SavedCollection            string
	DIYCollection              string
	RewardCollection           string
	BadgeCollection            string
	Badge_CollectionCollection string
	QuestionResultCollection   string
	QuizResultCollection       string
	ExchangedCollection        string
}
