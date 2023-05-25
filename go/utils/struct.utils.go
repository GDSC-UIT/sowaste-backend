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

type UserClaims struct {
	ISS           string `json:"uuid"`
	AZP           string `json:"azp"`
	AUD           string `json:"aud"`
	SUB           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	IAT           int64  `json:"iat"`
	EXP           int64  `json:"exp"`
}
