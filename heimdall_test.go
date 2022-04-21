package main

import (
	"os"
	"testing"
)

func TestGenerateUserHash(t *testing.T) {
	userHash := GenerateUserHash("testUserName", "testPassword")
	if userHash != "5553773136701468747" {
		t.Errorf("Generated user hash function is improper.")
	}
}

func TestGetUserHashPrex(t *testing.T) {
	userHashPrefix := GetUserHashPrex("7311242285026203200")
	if userHashPrefix != "731" {
		t.Errorf("Generated user hash prefix function is improper for positive numbers.")
	}
	userHashPrefix = GetUserHashPrex("-7311242285026203200")
	if userHashPrefix != "-731" {
		t.Errorf("Generated user hash prefix function is improper for negative numbers.")
	}
}

func TestAddUserHelper(t *testing.T) {
	os.RemoveAll("./data/testingRoutine")
	firstAdd := AddUser("testingRoutine", "test", "test")
	if firstAdd == false {
		t.Errorf("Adding unique user failed.")
	}
	secondAdd := AddUser("testingRoutine", "test", "test")
	if secondAdd == true {
		t.Errorf("Adding non-unique user was unexpectedly successful.")
	}
	os.RemoveAll("./data/testingRoutine")
}

func TestAuthUserHelper(t *testing.T) {
	os.RemoveAll("./data/testingRoutine")
	AddUser("testingRoutine", "test", "test")
	firstAuth := AuthUser("testingRoutine", "test", "test")
	if firstAuth == "" {
		t.Errorf("Authorizing existing user failed.")
	}
	secondAuth := AuthUser("testingRoutine", "notauserwhoexists", "notauserwhoexists")
	if secondAuth != "" {
		t.Errorf("Authorizing non-existing user was unexpectedly successful.")
	}
	os.RemoveAll("./data/testingRoutine")
}
