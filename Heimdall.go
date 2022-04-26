package main

import (
	"time"

	cryptography "heimdall.com/app/Cryptography"
	file "heimdall.com/app/FileIO"
)

const maximumAcceptableTime int64 = 1000

func main() {
	AddUser("testApp", "tester", "password")
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
	now := time.Now()
	userHash := cryptography.GenerateUserHash(username, password)
	prefixLength := file.GetAppFileGranularity(application)
	userPrex := cryptography.GetUserHashPrex(application, userHash, prefixLength)
	out <- file.WriteLine(application, userPrex, userHash)
	then := time.Now()
	if then.Sub(now).Microseconds() < file.BaseMaximumAddUserTime {
		file.IncreaseAppFileGranularity(application)
	}
}

func authUserHelper(application string, username string, password string, out chan string) {
	userHash := cryptography.GenerateUserHash(username, password)
	prefixLength := file.GetAppFileGranularity(application)
	userPrex := cryptography.GetUserHashPrex(application, userHash, prefixLength)
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
