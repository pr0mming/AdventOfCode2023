package problems_6_2

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
	var paperSheet [2]uint64
	i := 0

	for scanner.Scan() {
		line := scanner.Text()

		separatorIndex := strings.Index(line, ":")
		valuesStr := strings.Fields(line[separatorIndex+1:])

		paperSheet[i] = common_functions.ParseUint(strings.Join(valuesStr, ""))
		i++
	}

	time := paperSheet[0]
	distance := paperSheet[1]

	// This is the max distance possible
	timeHalf := uint64(math.Floor(float64(time) / 2))

	// Use this loop to find the first min distanceand keep in this variable
	// I split the way to avoid unecessary iterations
	minTime := timeHalf

	for {
		multiplier := time - minTime

		if minTime*multiplier > distance {
			minTime /= 2
		} else {
			break
		}
	}

	// From the previous min value we use the loop to find the first max distance
	records := 0

	for i := minTime + 1; ; i++ {
		multiplier := time - i

		if i*multiplier > distance {
			records = int(timeHalf-i) + 1
			break
		}

	}

	// We add up the double of the records it's a mirror "pattern"
	answer := records * 2

	// With even numbers we need minus 1
	if time%2 == 0 {
		answer--
	}

	return strconv.Itoa(answer)
}
