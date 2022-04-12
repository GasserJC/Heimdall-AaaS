package main

import (
	"fmt"
	"time"

	cryptography "heimdall.com/app/Cryptography"
	file "heimdall.com/app/FileIO"
)

func main() {
	start := time.Now()
	for i := 1; i <= 5000; i++ {
		AddUser("jgasser", "strongPassword")
		AddUser("agasser", "strongPassword")
		AddUser("sgasser", "strongPassword")
		RemoveUser("jgasser", "goodPassword")
		RemoveUser("agasser", "strongPassword")
		RemoveUser("sgasser", "strongestPassword")
	}
	duration := time.Since(start)
	fmt.Print("Seconds for 15,000 writes & 15,000 deletes: ")
	fmt.Println(duration)
	fmt.Print("Seconds per avg write/delete: ")
	fmt.Println(duration / 30000)
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
