package main

import (
	"fmt"
	"os"

	"aoc.2023/lib/functions"
)

func main() {
	argsWithProg := os.Args

	if len(argsWithProg) >= 2 {

		answer, err := functions.SolveProblemByKey(argsWithProg[1:])

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Answer: %s", answer)
		}

	} else {
		fmt.Println("You should give a valid args input: [problem] [part]")
	}
}
