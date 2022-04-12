package file

import (
	"bufio"
	"log"
	"os"
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

func OverWriteLines(filePath string, fileName string, rows []string) bool {
	if !FileExists(fileName) {
		createFile(filePath, fileName)
	} else {
		os.Truncate("./"+filePath+"/"+fileName+".txt", 0)
	}
	// write the value to the file.
	filePointer, err := os.OpenFile("./"+filePath+"/"+fileName+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		if !write(filePath, fileName, row, filePointer) {
			log.Fatal("Error occured while writing, process killed.")
		}
	}
	filePointer.Close()
	return false
}

func WriteLine(filePath string, fileName string, rowValue string) bool {
	// if the file does not exist, then create it.
	if !FileExists(fileName) {
		createFile(filePath, fileName)
	}
	// write the value to the file.
	filePointer, err := os.OpenFile("./"+filePath+"/"+fileName+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	if write(filePath, fileName, rowValue, filePointer) {
		filePointer.Close()
		return true
	}
	filePointer.Close()
	return false
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

func DeleteRow(filePath string, fileName string, valueToDelete string) bool {
	rows := Read(filePath, fileName)
	rows = delete(rows, valueToDelete)
	return OverWriteLines(filePath, fileName, rows)
}

func FileExists(fileName string) bool {
	fPointer, _ := os.Stat("./data/" + fileName + ".txt")
	return fPointer != nil
}

func IsRowUnique(rowToWrite string, filePath string, fileName string, useLowMemory bool) bool {
	if !useLowMemory {
		rows := Read(filePath, fileName)
		for _, rowValue := range rows {
			if rowValue == rowToWrite {
				return false
			}
		}
	}
	return true
}

func createFile(filePath string, fileName string) {
	file, err := os.Create("./" + filePath + "/" + fileName + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
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
