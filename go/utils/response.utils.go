package utils

type Response struct {
	Data    interface{} `json:"data" bson:"data"`
	Message string      `json:"message" bson:"message"`
}

func SuccessfulResponse(data interface{}, message string) *Response {
	return &Response{
		Data:    data,
		Message: message,
	}
}

func FailedResponse(message string) *Response {
	return &Response{
		Data:    nil,
		Message: message,
	}
}
