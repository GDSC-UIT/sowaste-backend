package utils

import (
	"encoding/json"
)

func JSONStringify(data interface{}) string {
	json, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(json)
}

func JSONParse(data string) interface{} {
	var result interface{}
	json.Unmarshal([]byte(data), &result)
	return result
}
