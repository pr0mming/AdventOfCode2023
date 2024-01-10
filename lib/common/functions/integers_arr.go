package common_functions

import (
	"slices"
	"strconv"
)

func GetIntegersArr(arr []string, addAscOrder bool) []int {
	integers := make([]int, len(arr))

	for i, v := range arr {
		n, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		integers[i] = n
	}

	if addAscOrder {
		slices.Sort(integers)
	}

	return integers
}
