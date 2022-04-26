package cryptography

import (
	"crypto/sha512"
	"fmt"

	typeConversion "heimdall.com/app/TypeConversion"
)

func Hash(inputString string) string {
	hashBytes := sha512.Sum512([]byte(inputString))
	hashNumber := typeConversion.ByteArrayToInt(hashBytes)
	return fmt.Sprint(hashNumber)
}

func GetUserHashPrex(application string, userHash string, prefixLength int) string {
	if userHash[0] == '-' {
		return userHash[0 : prefixLength+1]
	} else {
		return userHash[0:prefixLength]
	}
}

func GenerateUserHash(username string, password string) string {
	var hashString = username + password + password + username
	return Hash(hashString)
}
