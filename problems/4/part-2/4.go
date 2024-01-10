package problems_4_2

import (
	"fmt"
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
	var cards []string
	var cardsMemo = make(map[int]int) // We use it for map[index card] = number of cards/copies

	// Fill the maps first
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		cards = append(cards, line)
		cardsMemo[i] = 1 // Original there is 1 card always
	}

	for i, line := range cards {
		// Extra work to extract the numbers for each card
		firstIndexSeparator := strings.Index(line, ":")
		secondIndexSeparator := strings.Index(line, "|")

		winningNumbers := strings.Fields(line[firstIndexSeparator+1 : secondIndexSeparator])
		myNumbers := strings.Fields(line[secondIndexSeparator+1:])

		ocurrencies := computeScore(winningNumbers, myNumbers)

		if ocurrencies > 0 {
			// We take the next card [j + n] and we add the ammount of the current card [i]
			for j := i + 1; j <= i+ocurrencies && j < len(cardsMemo); j++ {
				cardsMemo[j] += cardsMemo[i]
			}
		}
	}

	for _, v := range cardsMemo {
		answer += v
	}

	return strconv.Itoa(answer)
}

func computeScore(winningNumbers []string, myNumbers []string) int {
	// Convert the arr string to arr integer, because it's used for Binary Search
	integersForWN := common_functions.GetIntegersArr(winningNumbers, true)
	integersForMN := common_functions.GetIntegersArr(myNumbers, true)
	n := 0

	for _, v := range integersForMN {
		i := sort.Search(len(integersForWN), func(i int) bool {
			return integersForWN[i] >= v
		})

		// This validation is if the number was founded
		if i < len(integersForWN) && integersForWN[i] == v {
			n += 1
		}
	}

	// Return the occurrences for this card
	return n
}
