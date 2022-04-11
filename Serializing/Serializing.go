package serial

import (
	"strings"

	user "heimdall.com/app/User"
)

func Deserialize(serializedValues []string, delimiter string) string {
	return (delimiter + strings.Join(serializedValues, delimiter))
}

func Serialize(row string) user.UC {
	serializedValues := splitDelimitedString(row)
	serializedObject := user.UC{Key: serializedValues[0], Username: serializedValues[1], Password: serializedValues[2]}
	return serializedObject
}

func splitDelimitedString(delimitedRow string) []string {
	var delimiter string
	var element string
	var delimitedValues []string
	var rowLength = len(delimitedRow)
	for position, character := range delimitedRow {
		if (string(character) == delimiter) && (position != 0) {
			delimitedValues = append(delimitedValues, element)
			element = ""
		} else if (position > 0) && (position < rowLength-1) {
			element += string(character)
		} else if position == rowLength-1 {
			delimitedValues = append(delimitedValues, element+string(character))
		} else {
			delimiter = string(character)
		}
	}
	return delimitedValues
}

func GetFirstValue(delimitedRow string) string {
	var element string
	var delimiter string
	for position, character := range delimitedRow {
		if (string(character) == delimiter) && (position != 0) {
			return element
		} else if position != 0 {
			element += string(character)
		} else {
			delimiter = string(character)
		}
	}
	return element
}
