package lib

import (
	"bufio"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Input struct {
	Raw []string
}

func ParseInput(path string) *Input {
	var input = new(Input)

	rawInput, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(rawInput)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input.Raw = append(input.Raw, scanner.Text())
	}

	return input
}

func (i *Input) ToInts() []int {
	var ints []int
	for _, line := range i.Raw {
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, i)
	}

	return ints
}
