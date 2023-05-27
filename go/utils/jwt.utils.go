package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func CheckJwtIsValidateAndGetJwtDecoded(jwt string) (UserClaims, bool) {
	var claims UserClaims

	parts := strings.Split(jwt, ".")
	encodedClaims := parts[1]
	decodedClaims, err := base64.RawURLEncoding.DecodeString(encodedClaims)
	if err != nil {
		fmt.Println("Error decoding claims:", err)
		return claims, false
	}
	err = json.Unmarshal(decodedClaims, &claims)
	if err != nil {
		fmt.Println("Error unmarshaling claims:", err)
		return claims, false
	}
	if len(parts) != 3 {
		return claims, false

	}
	return claims, true
}
