package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func predictNext(history []int) (int, int) {
	curr := history
	sequences := [][]int{}
	sequences = append(sequences, curr)

	for {
		nextseq := []int{}
		for i := 0; i < len(curr)-1; i++ {
			nextseq = append(nextseq, curr[i+1]-curr[i])
		}
		sequences = append(sequences, nextseq)
		curr = nextseq
		if completes(nextseq) {
			break
		}
	}

	lenseqs := len(sequences) - 1
	next := sequences[lenseqs][len(sequences[lenseqs])-1]
	prev := sequences[lenseqs][0]

	for i := len(sequences) - 2; i >= 0; i-- {
		next = sequences[i][len(sequences[i])-1] + next
		prev = sequences[i][0] - prev
	}

	return next, prev
}

func completes(seq []int) bool {
	for _, s := range seq {
		if s != 0 {
			return false
		}
	}
	return true
}

func solve(file string) {
	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	sol1 := 0
	sol2 := 0

	for scanner.Scan() {
		fs := strings.Fields(scanner.Text())

		history := []int{}
		for _, f := range fs {
			fi, _ := strconv.Atoi(f)
			history = append(history, fi)
		}

		next, prev := predictNext(history)
		sol1 += next
		sol2 += prev
	}

	fmt.Println("1: ", sol1)
	fmt.Println("2: ", sol2)
}

func main() {
	solve("./2023/day09/input.txt")
}
