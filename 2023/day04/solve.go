package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"regexp"
)

// numbers that exist in both slices
func intersection(arg0 []string, arg1 []string) []string {
	out := []string{}
	for _, el0 := range arg0 {
		for _, el1 := range arg1 {
			if el0 == el1 {
				out = append(out, el0)
			}
		}
	}
	return out
}

func solve(file string) {
	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	sol1 := 0
	sol2 := 0

	r := regexp.MustCompile(`^Card\s*[0-9]+: (.+)\|(.+)$`)
	r2 := regexp.MustCompile(`[0-9]+`)

	scratchcardsScores := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		matches := r.FindStringSubmatch(line)
		//fmt.Println(matches)
		//fmt.Println("Matches Wining", matches[1], "Your", matches[2])
		winningNumbers := r2.FindAllString(matches[1], -1)
		yourNumbers := r2.FindAllString(matches[2], -1)
		//fmt.Println("Numbers Wining", winningNumbers, "Your", yourNumbers)
		myWinningNumbers := intersection(winningNumbers, yourNumbers)
		//fmt.Println("Numbers Wining", winningNumbers, "Your", yourNumbers, "=> MyWinning", myWinningNumbers)
		sol1 += int(math.Pow(2, float64(len(myWinningNumbers))-1))

		scratchcardsScores = append(scratchcardsScores, len(myWinningNumbers))
	}

	// Part 2
	copies := []int{}

	for range scratchcardsScores {
		copies = append(copies, 1)
	}

	for i, score := range scratchcardsScores {
		//fmt.Println("card", i, "scored", score)
		max := i + score + 1
		if i+score > len(scratchcardsScores) {
			max = len(scratchcardsScores)
		}

		for j := i + 1; j < max; j++ {
			copies[j] += copies[i]
		}
	}

	for _, c := range copies {
		sol2 += c
	}

	fmt.Println("1: ", sol1)
	fmt.Println("2: ", sol2)
}

func main() {
	solve("./2023/day04/input.txt")
}
