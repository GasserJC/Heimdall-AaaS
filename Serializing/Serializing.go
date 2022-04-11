package serial

import (
	"strings"
	"heimdall.com/app/User"
	"fmt"
)

func Deserialize(listToSave []string, delimiter string) string{
	return (delimiter + strings.Join(listToSave, delimiter))
}

func Serialize(row string) user.UC{
	usrArr := splitDelimitedString(row)
	usr := user.UC{ Key:usrArr[0], Username:usrArr[1], Password:usrArr[2] }
	return usr
}

func splitDelimitedString(row string) []string {
	var delimiter = row[0:1]
	var element string
	var delimitedString []string
	var rowLength = len(row)
	for position, character := range row {
		if (string(character) == delimiter) && (position != 0) {
			
			delimitedString = append(delimitedString, element)
			element = ""
		} else if position != 0 {
			element += string(character)
		} else if position == rowLength - 1 {
			delimitedString = append(delimitedString, element+character)
		}
	}
	fmt.Println(delimitedString)
	return delimitedString
}