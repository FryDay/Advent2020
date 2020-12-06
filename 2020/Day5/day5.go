package main

import (
	"log"
	"sort"

	"github.com/FryDay/advent/lib"
)

const (
	rows    = 128
	columns = 8
)

func main() {
	seats := lib.ParseInput("input")
	asMap := toMap(seats.Raw)

	log.Println(partOne(asMap))
	log.Println(partTwo(asMap))
}

func partOne(seats map[string]int) int {
	highest := 0
	for _, v := range seats {
		if v > highest {
			highest = v
		}
	}

	return highest
}

func partTwo(seats map[string]int) int {
	seatIDs := make([]int, len(seats))
	for _, v := range seats {
		seatIDs = append(seatIDs, v)
	}
	sort.Ints(seatIDs)

	for i := 1; i < len(seatIDs); i++ {
		if seatIDs[i] == 0 {
			continue
		}
		if seatIDs[i+1]-seatIDs[i] > 1 {
			return seatIDs[i] + 1
		}
	}

	return 0
}

func toMap(lines []string) map[string]int {
	m := make(map[string]int)

	for _, line := range lines {
		m[line] = partition(line[:7], 0, rows)*8 + partition(line[7:], 0, columns)
	}

	return m
}

func partition(s string, min, max int) int {
	half := (max-min)/2 + min

	if len(s) == 0 {
		return half
	}

	if s[0] == 'F' || s[0] == 'L' {
		if max-min == 1 {
			return min
		}
		return partition(s[1:], min, half)
	} else if s[0] == 'B' || s[0] == 'R' {
		if max-min == 1 {
			return max
		}
		return partition(s[1:], half, max)
	}

	return 0
}
