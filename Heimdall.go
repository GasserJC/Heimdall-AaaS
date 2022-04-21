package main

import (
	cryptography "heimdall.com/app/Cryptography"
	file "heimdall.com/app/FileIO"
)

func AddUser(application string, username string, password string) bool {
	return addUserHelper(application, username, password)
}

func AuthUser(application string, username string, password string) string {
	return authUserHelper(application, username, password)
}

func addUserHelper(application string, username string, password string) bool {
	userHash := GenerateUserHash(username, password)
	userPrex := GetUserHashPrex(userHash)
	return file.WriteLine(application, userPrex, userHash)
}

func authUserHelper(application string, username string, password string) string {
	userHash := GenerateUserHash(username, password)
	userPrex := GetUserHashPrex(userHash)
	rows := file.Read(application, userPrex)
	for _, row := range rows {
		if row == userHash {
			return userHash
		}
	}
	return ""
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
