package problems_4_1

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var answer = 0

	for scanner.Scan() {
		line := scanner.Text()

		// Extra work to extract the numbers for each card
		firstIndexSeparator := strings.Index(line, ":")
		secondIndexSeparator := strings.Index(line, "|")

		winningNumbers := strings.Fields(line[firstIndexSeparator+1 : secondIndexSeparator])
		myNumbers := strings.Fields(line[secondIndexSeparator+1:])

		answer += computeScore(winningNumbers, myNumbers)
	}

	return strconv.Itoa(answer)
}

func computeScore(winningNumbers []string, myNumbers []string) int {
	// Convert the arr string to arr integer, because it's used for Binary Search
	var (
		integersForWN []int = common_functions.GetIntegersArr(winningNumbers, true)
		integersForMN []int = common_functions.GetIntegersArr(myNumbers, true)
		n             int   = 0
	)

	for _, v := range integersForMN {
		i := sort.Search(len(integersForWN), func(i int) bool {
			return integersForWN[i] >= v
		})

		// This validation is if the number was founded
		if i < len(integersForWN) && integersForWN[i] == v {
			n += 1
		}
	}

	if n > 0 {
		return int(math.Pow(2, float64(n-1)))
	}

	return 0
}
