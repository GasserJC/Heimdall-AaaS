package file

import (
	"fmt"

	cryptography "heimdall.com/app/Cryptography"
)

func increaseFileGranularity(application string, newGranularity int) {
	var endBound int = power(10, newGranularity)
	startBound := -1 * endBound
	for filename := startBound; filename < endBound; filename++ {
		rows := Read(application, fmt.Sprint(filename))
		for _, row := range rows {
			WriteLine(application, cryptography.GetUserHashPrex(application, row, newGranularity), row)
		}
	}
}

// Integer only exponents
func power(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	subpower := power(base, exponent/2)
	if exponent%2 == 0 {
		return subpower * subpower
	}
	return base * subpower * subpower
}
