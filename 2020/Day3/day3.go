package main

import (
	"log"

	"github.com/FryDay/advent/lib"
)

const tree = byte('#')

func main() {
	input := lib.ParseInput("input")

	partOne(input)
	partTwo(input)
}

func partOne(input *lib.Input) {
	log.Println(traverse(input, 3, 1))
}

func partTwo(input *lib.Input) {
	log.Println(traverse(input, 1, 1) * traverse(input, 3, 1) * traverse(input, 5, 1) * traverse(input, 7, 1) * traverse(input, 1, 2))
}

func traverse(input *lib.Input, right, down int) int {
	count := 0
	startChar := right
	lineLen := len(input.Raw[0])

	for i := down; i < len(input.Raw); i += down {
		if input.Raw[i][startChar] == tree {
			count++
		}

		startChar += right
		if startChar >= lineLen {
			startChar -= lineLen
		}
	}

	return count
}
