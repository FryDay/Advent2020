package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/FryDay/advent/lib"
)

type password struct {
	minRule  int
	maxRule  int
	letter   string
	password string
}

func (p *password) partOne() bool {
	count := strings.Count(p.password, p.letter)
	if count >= p.minRule && count <= p.maxRule {
		return true
	}
	return false
}

func (p *password) partTwo() bool {
	charOne := string(p.password[p.minRule-1])
	charTwo := string(p.password[p.maxRule-1])
	if (charOne == p.letter || charTwo == p.letter) && charOne != charTwo {
		return true
	}
	return false
}

func main() {
	passwords := parseInput()
	valid := 0

	for _, p := range passwords {
		if p.partOne() {
			valid++
		}
	}
	log.Println(valid)

	valid = 0
	for _, p := range passwords {
		if p.partTwo() {
			valid++
		}
	}
	log.Println(valid)
}

func parseInput() []*password {
	input := lib.ParseInput("input")
	passwords := []*password{}

	for _, line := range input.Raw {
		var p = new(password)

		p.minRule, _ = strconv.Atoi(strings.Split(line, "-")[0])
		p.maxRule, _ = strconv.Atoi(strings.Split(strings.SplitAfter(line, "-")[1], " ")[0])
		p.letter = strings.Split(strings.SplitAfter(line, " ")[1], ":")[0]
		p.password = strings.SplitAfter(line, " ")[2]

		passwords = append(passwords, p)
	}

	return passwords
}
