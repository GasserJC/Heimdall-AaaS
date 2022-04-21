package file

import (
	"bufio"
	"errors"
	"os"
	"time"
)

func Read(directory string, filename string) []string {
	relativeFilePath := "./data/" + directory + "/" + filename + ".txt"
	return readRelativePath(relativeFilePath)
}

func readRelativePath(relativePath string) []string {
	filePointer := getFilePointer(relativePath)
	var rows []string
	scanner := bufio.NewScanner(filePointer)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	filePointer.Close()
	return rows
}

func WriteLine(directory string, filename string, rowValue string) bool {
	ensureFolderPath(directory)
	fullRelativeFilePath := "./data/" + directory + "/" + filename + ".txt"
	filePointer := getFilePointer(fullRelativeFilePath)
	if write(fullRelativeFilePath, rowValue, filePointer) {
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

func IsRowUnique(rowToWrite string, fullRelativeFilePath string, useLowMemory bool) bool {
	if !useLowMemory {
		rows := readRelativePath(fullRelativeFilePath)
		for _, rowValue := range rows {
			if rowValue == rowToWrite {
				return false
			}
		}
	}
	return true
}

func ensureFolderPath(directory string) {
	// if the file does not exist, then create it.
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		createFolder("./data", "")
	}
	if _, err := os.Stat("./data/" + directory); os.IsNotExist(err) {
		createFolder("./data/"+directory, "")
	}
}

func getFilePointer(file string) *os.File {
	var filePointer *os.File
	var err error
	var attemptsLeft int = 100

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		filePointer, _ = os.Create(file)
		return filePointer
	}

	for {
		filePointer, err = os.Open(file)
		if (err == nil) || (attemptsLeft < 0) {
			break
		}
		attemptsLeft -= 1
		time.Sleep(5 * time.Millisecond)
	}

	return filePointer
}

func createFolder(filePath string, folderName string) {
	os.Mkdir("./"+filePath+"/"+folderName, os.ModePerm)
}

func write(fullRelativeFilePath string, rowValue string, filePointer *os.File) bool {
	if !IsRowUnique(rowValue, fullRelativeFilePath, false) {
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

func generateUserHash(userHash string) string {
	if userHash[0] == '-' {
		return userHash[0:4]
	} else {
		return userHash[0:3]
	}
}
