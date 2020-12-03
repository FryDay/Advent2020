package main

import (
	"github.com/FryDay/advent/lib"
)

func main() {
	inputInts := lib.ParseInput("input").ToInts()

	partOne(inputInts)
	partTwo(inputInts)
}

func partOne(inputs []int) {
	for x, y := range inputs {
		for i, j := range inputs {
			if x == i {
				continue
			}

			if y+j == 2020 {
				println(y * j)
				return
			}
		}
	}

}

func partTwo(inputs []int) {
	for x, y := range inputs {
		for i, j := range inputs {
			if x == i {
				continue
			}

			for a, b := range inputs {
				if x == a {
					continue
				}

				if y+j+b == 2020 {
					println(y * j * b)
					return
				}
			}
		}
	}

}
