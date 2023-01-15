package utils

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
