package problems

import (
	"math"
	"sort"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
)

func solveChallenge() int {
	// Process the input
	scanner := common_functions.CreateInputScanner("../input.txt")
	defer scanner.File.Close()

	var answer = 0

	for scanner.Scan() {
		line := scanner.Text()

		firstIndexSeparator := strings.Index(line, ":")
		secondIndexSeparator := strings.Index(line, "|")

		winningNumbers := strings.Fields(line[firstIndexSeparator+1 : secondIndexSeparator])
		myNumbers := strings.Fields(line[secondIndexSeparator+1:])

		answer += computeScore(winningNumbers, myNumbers)
	}

	return answer
}

func computeScore(winningNumbers []string, myNumbers []string) int {
	integersForWN := getIntegersSlice(winningNumbers)
	integersForMN := getIntegersSlice(myNumbers)
	n := 0

	for _, v := range integersForMN {
		i := sort.Search(len(integersForWN), func(i int) bool {
			return integersForWN[i] >= v
		})

		if i < len(integersForWN) && integersForWN[i] == v {
			n += 1
		}
	}

	if n > 0 {
		return int(math.Pow(2, float64(n-1)))
	}

	return 0
}

func getIntegersSlice(arr []string) []int {
	integers := make([]int, len(arr))

	for i, v := range arr {
		n, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		integers[i] = n
	}

	sort.Ints(integers)

	return integers
}
