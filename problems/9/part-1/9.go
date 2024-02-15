package problems_9_1

import (
	"fmt"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	answer := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Extract records (array) from the input
		historyRecord := common_functions.GetIntegersArr(strings.Fields(line), false)

		answer += computeHistoryRecord(historyRecord)
	}

	return strconv.Itoa(answer)
}

func computeHistoryRecord(historyRecord []int) int {
	var (
		zeros       int = 0                      // Flag to check if we have all the values of the history record are zeros
		offsetIndex int = len(historyRecord) - 1 // Each iteration the history records is len - 1, we use it to control the new record and access the last item
		sum         int = historyRecord[offsetIndex]
	)

	for zeros < len(historyRecord) {
		// New empty copy of the history record
		var newHistoryRecord = make([]int, offsetIndex)

		// Fill the array with differences
		for i := 1; i < len(historyRecord); i++ {
			diff := historyRecord[i] - historyRecord[i-1]
			newHistoryRecord[i-1] = diff

			if diff == 0 {
				zeros++
			}
		}

		historyRecord = newHistoryRecord

		// App up last item
		sum += newHistoryRecord[offsetIndex-1]
		offsetIndex-- // The next array will be fewer one
	}

	return sum
}
