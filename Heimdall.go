package main

import (
	"fmt"
	"math/rand"
	"time"

	cryptography "heimdall.com/app/Cryptography"
	file "heimdall.com/app/FileIO"
)

func main() {
	users := generateArrayOfRandomStrings(5000)
	passwords := generateArrayOfRandomStrings(5000)
	start := time.Now()
	for i := 1; i < 5000; i++ {
		AddUser(users[i], passwords[i])
	}
	duration := time.Since(start)
	fmt.Print("Seconds for 500 writes")
	fmt.Println(duration)
	fmt.Print("Seconds per avg write/delete: ")
	fmt.Println(duration / 5000)
}
func generateArrayOfRandomStrings(size int) []string {
	var randomStrings []string
	for i := 1; i <= size; i++ {
		randomStrings = append(randomStrings, randomString(10))
	}
	return randomStrings
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func AddUser(username string, password string) bool {
	userHash := generateUserHash(username, password)
	return file.WriteLine("data", "output", userHash)
}

func AuthUser(username string, password string) string {
	userHash := generateUserHash(username, password)
	var rows = file.Read("data", "output")
	for _, row := range rows {
		if row == userHash {
			return userHash
		}
	}
	return ""
}

func RemoveUser(username string, password string) {
	userHash := AuthUser(username, password)
	if userHash != "" {
		file.DeleteRow("data", "output", userHash)
	}
}

func generateUserHash(username string, password string) string {
	var hashString = username + password + password + username
	return cryptography.Hash(hashString)
}
