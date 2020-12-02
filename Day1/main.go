package main

import (
	"bufio"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var inputInts []int

	rawInputs, _ := ioutil.ReadFile("input")
	scanner := bufio.NewScanner(strings.NewReader(string(rawInputs)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		inputInts = append(inputInts, i)
	}

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
