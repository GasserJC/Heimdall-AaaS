package typeConversion

import (
	"fmt"
	"testing"
)

func TestByteArrayToInt(t *testing.T) {
	byteArr := [64]byte{213, 188, 168, 38, 20, 206, 182, 78, 110, 95, 81, 88, 231, 233, 184, 33, 128, 19, 199, 147, 150, 167, 26, 104, 217, 126, 150, 79, 116, 0, 18, 201, 213, 188, 168, 38, 20, 206, 182, 78, 110, 95, 81, 88, 231, 233, 184, 33, 128, 19, 199, 147, 150, 167, 26, 104, 217, 126, 150, 79, 116, 0, 18, 201}
	fmt.Println()
	integerValue := ByteArrayToInt(byteArr)
	fmt.Println(integerValue)
	if integerValue != 5671947366662978773 {
		t.Errorf("Byte Array to Integer conversion failed.")
	}
}
