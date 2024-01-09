package functions

import (
	"errors"

	problems_1_1 "aoc.2023/problems/1/part-1"
	problems_1_2 "aoc.2023/problems/1/part-2"
)

func SolveProblemByKey(key string) (string, error) {
	var answer string

	switch key {
	case "11":
		answer = problems_1_1.SolveChallenge()
	case "12":
		answer = problems_1_2.SolveChallenge()
	default:
		return "", errors.New("The given input is not in a valid range, try something like: [1 1]")
	}

	return answer, nil
}
