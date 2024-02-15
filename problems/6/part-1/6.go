package problems_6_1

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	// Keep the values in the same format (time-distance) in an array (2 rows)
	var (
		paperSheet [2][]int
		i          int = 0
	)

	for scanner.Scan() {
		line := scanner.Text()

		separatorIndex := strings.Index(line, ":")
		valuesStr := strings.Fields(line[separatorIndex+1:])

		paperSheet[i] = common_functions.GetIntegersArr(valuesStr, false)
		i++
	}

	// Use it to keep the final multiplication
	answer := 1

	// Iterate over the columns
	for i := 0; i < len(paperSheet[0]); i++ {
		time := paperSheet[0][i]
		distance := paperSheet[1][i]

		// With the half we start with the max distance instead the beginning
		timeHalf := int(math.Floor(float64(time) / 2))

		// Use it to keep the number of max distances
		records := 0

		// Do the sequence from the max until the min distance (> distance)
		for x := timeHalf; ; x-- {
			multiplier := time - x

			if x*multiplier > distance {
				records++
			} else {
				break
			}
		}

		// We add up the double of the records it's a mirror "pattern"
		records *= 2

		// With even numbers we need minus 1
		if time%2 == 0 {
			records--
		}

		answer *= records
	}

	return strconv.Itoa(answer)
}
