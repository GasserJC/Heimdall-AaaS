package heimdall

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	cryptography "heimdall.com/app/Cryptography"
	file "heimdall.com/app/FileIO"
)

func AddUser(application string, username string, password string) {
	go addUserHelper(application, username, password)
}

func AuthUser(application string, username string, password string) {
	go authUserHelper(application, username, password)
}

func addUserHelper(application string, username string, password string) bool {
	userHash := generateUserHash(username, password)
	userPrex := getUserHashPrex(userHash)
	return file.WriteLine(application, "output", userHash, userPrex)
}

func authUserHelper(application string, username string, password string) string {
	userHash := generateUserHash(username, password)
	userPrex := getUserHashPrex(userHash)
	rows := file.Read(application, "output", userPrex)
	for _, row := range rows {
		if row == userHash {
			return userHash
		}
	}
	return ""
}

func generateArrayOfRandomStrings(size int) []string {
	var randomStrings []string
	for i := 1; i <= size; i++ {
		randomStrings = append(randomStrings, randomString(10, i))
	}
	return randomStrings
}

func randomString(length int, mutex int) string {
	rand.Seed(time.Now().UnixNano() + int64(mutex))
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func getUserHashPrex(userHash string) string {
	if userHash[0] == '-' {
		return userHash[0:4]
	} else {
		return userHash[0:3]
	}
}

func generateUserHash(username string, password string) string {
	var hashString = username + password + password + username
	return cryptography.Hash(hashString)
}

func buffer() {
	if runtime.NumGoroutine() > 10000 {
		time.Sleep(time.Microsecond)
	}
}
