package file

import (
	"bufio"
	"os"
	"time"
)

func Read(filePath string, fileName string, prex string) []string {
	var path string = "./" + filePath + "/" + prex + "/" + fileName + ".txt"
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

func WriteLine(filePath string, fileName string, rowValue string, rowPrex string) bool {
	// if the file does not exist, then create it.
	filePath = "data/" + filePath
	if _, err := os.Stat("./" + filePath); os.IsNotExist(err) {
		createFolder(filePath, "")
	}
	if _, err := os.Stat("./" + filePath + "/" + rowPrex); os.IsNotExist(err) {
		createFolder(filePath, rowPrex)
	}
	filePath = filePath + "/" + rowPrex

	filePointer := getFilePointer("./" + filePath + "/" + fileName + ".txt")
	if write(filePath, fileName, rowValue, filePointer) {
		filePointer.Close()
		return true
	}
	filePointer.Close()
	return false
}

func FileExists(fileName string) bool {
	fPointer, _ := os.Stat("./data/" + fileName + ".txt")
	return fPointer != nil
}

func IsRowUnique(rowToWrite string, filePath string, fileName string, useLowMemory bool) bool {
	if !useLowMemory {
		rows := Read(filePath, fileName, rowToWrite[0:3])
		for _, rowValue := range rows {
			if rowValue == rowToWrite {
				return false
			}
		}
	}
	return true
}

func getFilePointer(file string) *os.File {
	var filePointer *os.File
	var err error
	for {
		filePointer, err = os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err == nil {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	return filePointer
}

func createFolder(filePath string, folderName string) {
	os.Mkdir("./"+filePath+"/"+folderName, os.ModePerm)
}

func write(filePath string, fileName string, rowValue string, filePointer *os.File) bool {
	if !IsRowUnique(rowValue, filePath, fileName, false) {
		return false
	}
	var err error
	if _, err = filePointer.WriteString(rowValue + "\n"); err != nil {
		panic(err)
	}
	return true
}

func delete(slice []string, element string) []string {
	index := indexOf(slice, element)
	if index >= 0 {
		slice = removeIndex(slice, index)
	}
	return slice
}

func removeIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func indexOf(slice []string, element string) int {
	for i := range slice {
		if slice[i] == element {
			return i
		}
	}
	return -1
}
