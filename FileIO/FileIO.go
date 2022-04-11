package file

import (
	"bufio"
	"fmt"
	"log"
	"os"

	serial "heimdall.com/app/Serializing"
	user "heimdall.com/app/User"
)

func Read(filePath string, fileName string) []string {
	var path string = "./" + filePath + "/" + fileName + ".txt"
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	var rows []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	file.Close()
	return rows
}

func Write(filePath string, fileName string, usr user.UC) {
	// if the file does not exist, then create it.
	if !FileExists(fileName) {
		createFile(filePath, fileName)
	}

	usrKey := usr.Slice()[0]

	if !PrimaryKeyIsUnique(usrKey, filePath, fileName, false) {
		log.Fatal("Duplicate Primary Key")
	}
	usrString := serial.Deserialize(usr.Slice(), "|")

	// write the value to the file.
	filePointer, err := os.OpenFile("./"+filePath+"/"+fileName+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	if _, err = filePointer.WriteString(usrString + "\n"); err != nil {
		panic(err)
	}
	filePointer.Close()
}

func FileExists(fileName string) bool {
	fPointer, _ := os.Stat("./data/" + fileName + ".txt")
	return fPointer != nil
}

func createFile(filePath string, fileName string) {
	file, err := os.Create("./" + filePath + "/" + fileName + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func PrimaryKeyIsUnique(newKey string, filePath string, fileName string, useLowMemory bool) bool {
	if !useLowMemory {
		rows := Read(filePath, fileName)
		for _, rowValue := range rows {
			fmt.Println(len(newKey))
			fmt.Println(len(serial.GetFirstValue(rowValue)))
			if newKey == serial.GetFirstValue(rowValue) {
				return false
			}
		}
	}
	return true
}
