package file

import (
	"os"
	"log"
    "bufio"
)

func Read(fileName string) []string {
	var path string = "./data/"+fileName+".txt"
    file, err := os.Open(path)
    if err != nil {
        return nil
    }

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
	file.Close()
    return lines
}

func Write(filePath string, fileName string, value string){
	// if the file does not exist, then create it.
	if(!FileExists(fileName)){createFile(filePath,fileName)}

	// write the value to the file.
	filePointer, err := os.OpenFile("./data/"+fileName+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
 	   panic(err)
	}

	if _, err = filePointer.WriteString(value+"\n"); err != nil {
	    panic(err)
	}
	filePointer.Close()
}

func createFile(filePath string,fileName string)  {
	file, err := os.Create("./"+filePath+"/"+fileName+".txt");
	if err != nil {
		log.Fatal(err)
	}
	file.Close();
}

func FileExists(fileName string) bool {
	fpointer, _ := os.Stat("./data/"+fileName+".txt")
	if fpointer != nil {
		return true
	}
	return false
}