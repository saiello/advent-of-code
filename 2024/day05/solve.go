package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	regOrder    = regexp.MustCompile(`^([0-9]+)\|([0-9]+)$`)
	regPrintSeq = regexp.MustCompile(`([0-9]+,?)+`)
)

type rule struct {
	before int
	after  int
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func isValid(rules map[rule]bool, seq []int) bool {
	for i, before := range seq {
		tail := seq[i+1:]

		for _, after := range tail {
			if !rules[rule{before: before, after: after}] {
				return false
			}
		}
	}
	return true
}

func main() {

	var sum, sum2 int

	rules := map[rule]bool{}
	printSequences := [][]int{}

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		line := scanner.Text()

		if regOrder.MatchString(line) {
			m := regOrder.FindStringSubmatch(line)
			rules[rule{before: atoi(m[1]), after: atoi(m[2])}] = true
		} else if regPrintSeq.MatchString(line) {
			pages := []int{}

			for _, el := range strings.Split(line, ",") {
				pages = append(pages, atoi(el))
			}

			printSequences = append(printSequences, pages)
		}
	}

	// Part-1
	for _, seq := range printSequences {
		if isValid(rules, seq) {
			sum += seq[len(seq)/2]
		}
	}

	// Part-2
	for _, seq := range printSequences {
		if !isValid(rules, seq) {

			slices.SortFunc(seq, func(a int, b int) int {
				if a == b {
					return 0
				} else if rules[rule{before: a, after: b}] {
					return -1
				} else {
					return 1
				}
			})

			sum2 += seq[len(seq)/2]
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
