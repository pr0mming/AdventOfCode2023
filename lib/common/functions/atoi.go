package common_functions

import "strconv"

// Encapsulation for the Atoi util
func Atoi(numberStrTmp string) int {
	numberIntTmp, err := strconv.Atoi(numberStrTmp)
	if err != nil {
		panic(err)
	}

	return numberIntTmp
}
