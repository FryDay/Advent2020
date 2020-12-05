package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/FryDay/advent/lib"
)

type passport struct {
	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColor      string
	eyeColor       string
	passportID     string
	countryID      string
}

func (p *passport) partOne() bool {
	if p.birthYear == "" ||
		p.issueYear == "" ||
		p.expirationYear == "" ||
		p.height == "" ||
		p.hairColor == "" ||
		p.eyeColor == "" ||
		p.passportID == "" {
		return false
	}

	return true
}

func (p *passport) partTwo() bool {
	if !p.partOne() {
		return false
	}

	// (Birth Year) - four digits; at least 1920 and at most 2002.
	if !validYearRange(p.birthYear, 1920, 2002) {
		return false
	}

	// (Issue Year) - four digits; at least 2010 and at most 2020.
	if !validYearRange(p.issueYear, 2010, 2020) {
		return false
	}

	// (Expiration Year) - four digits; at least 2020 and at most 2030.
	if !validYearRange(p.expirationYear, 2020, 2030) {
		return false
	}

	// (Height) - a number followed by either cm or in:
	//   If cm, the number must be at least 150 and at most 193.
	//   If in, the number must be at least 59 and at most 76.
	switch {
	case strings.HasSuffix(p.height, "cm"):
		if !validRange(strings.Split(p.height, "cm")[0], 150, 193) {
			return false
		}
	case strings.HasSuffix(p.height, "in"):
		if !validRange(strings.Split(p.height, "in")[0], 59, 76) {
			return false
		}
	default:
		return false
	}

	// (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	matched, err := regexp.MatchString(`^#[a-zA-Z0-9]{6,}$`, p.hairColor)
	if err != nil || !matched {
		return false
	}

	// (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if p.eyeColor != "amb" &&
		p.eyeColor != "blu" &&
		p.eyeColor != "brn" &&
		p.eyeColor != "gry" &&
		p.eyeColor != "grn" &&
		p.eyeColor != "hzl" &&
		p.eyeColor != "oth" {
		return false
	}

	// (Passport ID) - a nine-digit number, including leading zeroes.
	if len(p.passportID) != 9 {
		return false
	}
	_, err = strconv.Atoi(p.passportID)
	if err != nil {
		return false
	}

	return true
}

func main() {
	passports := parseInput()
	valid := 0

	for _, p := range passports {
		if p.partOne() {
			valid++
		}
	}
	log.Println(valid)

	valid = 0
	for _, p := range passports {
		if p.partTwo() {
			valid++
		}
	}
	log.Println(valid)
}

func validYearRange(year string, min, max int) bool {
	if len(year) != 4 {
		return false
	}

	return validRange(year, min, max)
}

func validRange(value string, min, max int) bool {
	asInt, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	if asInt < min || asInt > max {
		return false
	}

	return true
}

func parseInput() []*passport {
	input := lib.ParseInput("input")
	passports := []*passport{}

	p := new(passport)
	for _, line := range input.Raw {
		if line == "" {
			passports = append(passports, p)
			p = new(passport)
			continue
		}

		kvs := strings.Split(line, " ")
		for _, i := range kvs {
			kv := strings.Split(i, ":")
			switch kv[0] {
			case "byr":
				p.birthYear = kv[1]
			case "iyr":
				p.issueYear = kv[1]
			case "eyr":
				p.expirationYear = kv[1]
			case "hgt":
				p.height = kv[1]
			case "hcl":
				p.hairColor = kv[1]
			case "ecl":
				p.eyeColor = kv[1]
			case "pid":
				p.passportID = kv[1]
			case "cid":
				p.countryID = kv[1]
			}
		}
	}
	passports = append(passports, p)

	return passports
}
