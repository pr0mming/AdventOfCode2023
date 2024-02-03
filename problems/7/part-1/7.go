package problems_7_1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	common_functions "aoc.2023/lib/common/functions"
	. "aoc.2023/lib/common/types/problems/7"
)

// Kepp the alphabet equivalence to check which letter is higher
// (runes approach (ASCII Code) doesn't work)
var ALPHABET_MAP = map[string]int{
	"A": 1,
	"K": 2,
	"Q": 3,
	"J": 4,
	"T": 5,
	"9": 6,
	"8": 7,
	"7": 8,
	"6": 9,
	"5": 10,
	"4": 11,
	"3": 12,
	"2": 13,
}

func SolveChallenge(problemId string) string {
	// Process the input
	inputFilePath := fmt.Sprintf("problems/%s/input.txt", problemId)
	scanner := common_functions.CreateInputScanner(inputFilePath)
	defer scanner.File.Close()

	var camelCards []CamelCard

	// To avoid extra convertions (int -> string -> int) and ease sorting I decided use a struct to keep different values
	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Fields(line)

		camelCards = append(camelCards, CamelCard{
			Card:     values[0],
			Bid:      common_functions.Atoi(values[1]),
			CardType: -1,
		})
	}

	// Get the card type before the sorting operations (avoid extra iterations withiin slices.SortFunc)
	for i := 0; i < len(camelCards); i++ {
		camelCards[i].CardType = computeCamelCard(camelCards[i].Card)

		//fmt.Println(camelCards[i].Card, camelCards[i].CardType)
	}

	sortCamelCards(&camelCards)

	answer := 0

	// Make final operation with sorted cards
	for i, v := range camelCards {
		answer += ((i + 1) * v.Bid)
	}

	return strconv.Itoa(answer)
}

func sortCamelCards(camelCards *[]CamelCard) {
	slices.SortFunc(*camelCards, func(a, b CamelCard) int {
		// Final sorting logic is here ...
		comparision := cmpCamelCards(a, b)

		return comparision
	})
}

func cmpCamelCards(a CamelCard, b CamelCard) int {
	// I decided to assign cards codes using integers (1, 2, 3...)
	// But 1 is actually the strongest card, here we do a DESC sorting using the CardType
	if a.CardType < b.CardType {
		return 1
	}

	if a.CardType > b.CardType {
		return -1
	}

	// If both are the same type, we do the same logic above but this time comparing each char
	for i := 0; i < len(a.Card); i++ {
		if getCharEqByRune(a.Card[i]) < getCharEqByRune(b.Card[i]) {
			return 1
		}

		if getCharEqByRune(a.Card[i]) > getCharEqByRune(b.Card[i]) {
			return -1
		}
	}

	return 0
}

func computeCamelCard(camelCard string) int {
	// I think is easy keep the occurs by each letter using a map
	var camelCardMap = make(map[rune]int)

	for _, v := range camelCard {
		if _, ok := camelCardMap[v]; ok {
			camelCardMap[v]++
		} else {
			camelCardMap[v] = 1
		}
	}

	// We extract the values of the map...
	var cardsAmmount []int

	for _, v := range camelCardMap {
		cardsAmmount = append(cardsAmmount, v)
	}

	// We extract the highest number from the map to check the Camel Card type
	maxCardAmmount := slices.Max(cardsAmmount)

	// 1 is Five of a kind
	// 2 is Four of a kind
	// And so on ...

	switch len(cardsAmmount) {
	case 1:
		return 1
	case 2:
		if maxCardAmmount == 4 {
			return 2
		} else {
			return 3
		}
	case 3:
		if maxCardAmmount == 3 {
			return 4
		} else {
			return 5
		}
	case 4:
		if maxCardAmmount == 2 {
			return 6
		}
	case 5:
		return 7
	}

	return -1
}

func getCharEqByRune(char byte) int {
	return ALPHABET_MAP[string(char)]
}
