package file

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func Read(directory string, filename string) []string {
	relativeFilePath := baseDir + directory + "/" + filename + baseFileType
	return readRelativePath(relativeFilePath)
}

func readRelativePath(relativePath string) []string {
	filePointer := getFilePointer(relativePath, true)
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
	fullRelativeFilePath := baseDir + directory + "/" + filename + baseFileType
	filePointer := getFilePointer(fullRelativeFilePath, false)
	success := write(fullRelativeFilePath, rowValue, filePointer)
	filePointer.Close()
	return success
}

func FileExists(fileName string) bool {
	fPointer, _ := os.Stat(baseDir + fileName + baseFileType)
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

func IncreaseAppFileGranularity(application string) {
	files, _ := ioutil.ReadDir(baseDir + application + "/" + baseApplicationConfigFile)
	averageFileSize := getAverageSizePerFile(files)
	if averageFileSize > baseMaximumBytesPerFile {
		os.RemoveAll(baseDir + application + "/" + baseApplicationConfigFile)
		newGranularity := GetAppFileGranularity(application) + 1
		WriteLine(application, baseApplicationConfigFile, fmt.Sprint(newGranularity))
		increaseFileGranularity(application, newGranularity)
	}
}

func getAverageSizePerFile(files []os.FileInfo) int {
	var count int64 = 0
	var totalSize int64 = 0
	for _, file := range files {
		if !file.IsDir() {
			count++
			totalSize += file.Size()
		}
	}
	return int(totalSize / count)
}

func GetAppFileGranularity(application string) int {
	rows := Read(application, baseApplicationConfigFile)
	for _, rowValue := range rows {
		if strings.Contains(rowValue, baseGranularityString) {
			granularityString, _ := strconv.Atoi(rowValue[len(baseGranularityString):])
			return granularityString
		}
	}
	return baseGranularity
}

func ensureFolderPath(directory string) {
	// if the file does not exist, then create it.
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		createFolder(baseDir, "")
	}
	if _, err := os.Stat(baseDir + directory); os.IsNotExist(err) {
		createFolder(baseDir+directory, "")
	}
}

func getFilePointer(file string, readOnly bool) *os.File {
	var filePointer *os.File
	var err error
	attemptsLeft := baseAttemptsLeft
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		filePointer, _ = os.Create(file)
		return filePointer
	}

	for {
		if readOnly {
			filePointer, err = os.Open(file)
		} else {
			filePointer, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		}

		if (err == nil) || (attemptsLeft < 0) {
			break
		}
		attemptsLeft = resequence(attemptsLeft)
	}

	return filePointer
}

func resequence(attemptsLeft int) int {
	if attemptsLeft > 0 {
		attemptsLeft--
		time.Sleep(baseResequenceTime * time.Millisecond)
	}
	return attemptsLeft
}

func createFolder(filePath string, folderName string) {
	os.Mkdir("./"+filePath+"/"+folderName, os.ModePerm)
}

func write(fullRelativeFilePath string, rowValue string, filePointer *os.File) bool {
	if !IsRowUnique(rowValue, fullRelativeFilePath, false) {
		return false
	}
	attemptsLeft := baseAttemptsLeft
	var err error
	for {
		if _, err = filePointer.WriteString(rowValue + "\n"); err == nil {
			break
		}
		attemptsLeft = resequence(attemptsLeft)
		if (attemptsLeft == 0) && err != nil {
			panic(err)
		}
	}
	return true
}
