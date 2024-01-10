package common_functions

import "strconv"

func ParseUint(numberStrTmp string) uint64 {
	numberIntTmp, err := strconv.ParseUint(numberStrTmp, 10, 64)
	if err != nil {
		panic(err)
	}

	return numberIntTmp
}
