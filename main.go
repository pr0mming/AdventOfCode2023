package main

import (
	"fmt"
	"os"
	"strings"

	"aoc.2023/lib/functions"
)

func main() {
	argsWithProg := os.Args

	if len(argsWithProg) >= 2 {
		problemFlag := strings.Join(argsWithProg[1:], "")

		answer, err := functions.SolveProblemByKey(problemFlag)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(answer)
		}

	} else {
		fmt.Println("You should give a valid args input: [problem] [part]")
	}
}
