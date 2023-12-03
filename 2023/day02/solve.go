package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// only 12 red cubes, 13 green cubes, and 14 blue cubes
var bag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func isPossible(drafts [][]string) bool {
	for _, draft := range drafts {
		qty, _ := strconv.Atoi(draft[1])
		color := draft[2]

		if qty > bag[color] {
			return false
		}
	}

	return true
}

func minimumSet(drafts [][]string) map[string]int {

	minSet := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, draft := range drafts {
		qty, _ := strconv.Atoi(draft[1])
		color := draft[2]
		if minSet[color] < qty {
			minSet[color] = qty
		}
	}

	return minSet
}

func powerOfSet(set map[string]int) int {
	power := 1
	for _, qty := range set {
		power *= qty
	}
	return power
}

func main() {

	input, _ := os.ReadFile("./2023/day02/input.txt")
	scanner := bufio.NewScanner(bytes.NewReader(input))

	r := regexp.MustCompile(`^Game (?P<gamenum>[0-9]+):(?P<gameconf>.+)$`)
	r2 := regexp.MustCompile(`(?P<qty>[0-9]+) (?P<color>(red|blue|green))`)

	sum1 := 0
	sum2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		sm := r.FindStringSubmatch(line)

		game, _ := strconv.Atoi(sm[1])
		drafts := r2.FindAllStringSubmatch(sm[2], -1)

		// part 1
		if isPossible(drafts) {
			sum1 += game
		}

		// part 2
		sum2 += powerOfSet(minimumSet(drafts))

	}

	fmt.Println("1: ", sum1)
	fmt.Println("2: ", sum2)
}
