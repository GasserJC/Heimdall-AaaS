package main

import (
	"crypto/sha256"
	"fmt"
	"unsafe"

	file "heimdall.com/app/FileIO"
	serial "heimdall.com/app/Serializing"
	user "heimdall.com/app/User"
)

func main() {
	AddUser("jgasser", "strongPassword")
	fmt.Println(AuthUser("jgasser", "strongPassword"))
}

func AddUser(username string, password string) {
	key := generateKey(username, password)
	usr := user.UC{Key: key, Username: username, Password: password}
	file.Write("data", "output", usr)
}

func AuthUser(username string, password string) bool {
	var rows = file.Read("data", "output")
	for _, row := range rows {
		user := serial.Serialize(row)
		if (user.Username == username) && (user.Password == password) {
			return true
		}
	}
	return false
}

func generateKey(username string, password string) string {
	// using a string1_string2_string2_string1 input string for a hash means it does not mean
	// a user with a username of someone else's password which has their password as their username is
	// not going to produce the same hash
	var hashString = username + password + password + username
	return hash(hashString)
}

func hash(inputString string) string {
	hashBytes := sha256.Sum256([]byte(inputString))
	hashNumber := byteArrayToInt(hashBytes)
	return fmt.Sprint(hashNumber)
}

func byteArrayToInt(arr [32]byte) int64 {
	val := int64(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}
