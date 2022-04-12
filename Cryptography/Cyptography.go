package cryptography

import (
	"crypto/sha256"
	"fmt"

	typeConversion "heimdall.com/app/TypeConversion"
)

func Hash(inputString string) string {
	hashBytes := sha256.Sum256([]byte(inputString))
	hashNumber := typeConversion.ByteArrayToInt(hashBytes)
	return fmt.Sprint(hashNumber)
}
