package main

import (
	cryptography "heimdall.com/app/Cryptography"
	file "heimdall.com/app/FileIO"
)

func main() {
	reportPerformance()
}

func AddUser(application string, username string, password string) bool {
	out := make(chan bool)
	go addUserHelper(application, username, password, out)
	return <-out
}

func AuthUser(application string, username string, password string) string {
	out := make(chan string)
	go authUserHelper(application, username, password, out)
	return <-out
}

func addUserHelper(application string, username string, password string, out chan bool) {
	userHash := GenerateUserHash(username, password)
	userPrex := GetUserHashPrex(userHash)
	out <- file.WriteLine(application, userPrex, userHash)
}

func authUserHelper(application string, username string, password string, out chan string) {
	userHash := GenerateUserHash(username, password)
	userPrex := GetUserHashPrex(userHash)
	authUserHash := ""
	rows := file.Read(application, userPrex)
	for _, row := range rows {
		if row == userHash {
			authUserHash = userHash
			break
		}
	}
	out <- authUserHash
}

func GetUserHashPrex(userHash string) string {
	if userHash[0] == '-' {
		return userHash[0:4]
	} else {
		return userHash[0:3]
	}
}

func GenerateUserHash(username string, password string) string {
	var hashString = username + password + password + username
	return cryptography.Hash(hashString)
}
