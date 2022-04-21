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
