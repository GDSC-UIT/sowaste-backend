package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

const maxObjectIDLength = 24

func validEmailLength(email string) bool {
	return len(email) == 24
}

func EmailProcess(email string) string {
	if validEmailLength(email) {
		return email
	}
	for i := 0; i < maxObjectIDLength-len(email); i++ {
		email = email + "&"
	}
	return email
}

func EmailToMongoID(email string) []byte {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	hash := md5.Sum([]byte(email))
	mongoID := make([]byte, 12)
	hex.Decode(mongoID, []byte(hash[:]))

	return mongoID
}

func EmailToID(email string) (primitive.ObjectID, error) {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	hash := md5.Sum([]byte(email))
	mongoID := make([]byte, 12)
	hex.Decode(mongoID, []byte(hash[:]))
	fmt.Println(mongoID)
	return primitive.ObjectIDFromHex(string(mongoID))
}
