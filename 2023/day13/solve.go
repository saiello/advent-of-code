package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type Pattern [][]rune

func parse(file string) []Pattern {
	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	patterns := []Pattern{}
	p := Pattern{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			patterns = append(patterns, p)
			p = Pattern{}
			continue
		}

		p = append(p, []rune{})
		for _, c := range line {
			p[len(p)-1] = append(p[len(p)-1], c)
		}
	}

	patterns = append(patterns, p)

	return patterns
}

func solve(file string) {

	patterns := parse(file)

	fmt.Println("1: ", sumReflections(patterns, 0))
	fmt.Println("2: ", sumReflections(patterns, 1))
}

func sumReflections(patterns []Pattern, smudge int) int {
	sol := 0
	for _, p := range patterns {
		if ok, row := reflectionPoint(p, smudge); ok {
			//fmt.Println(idx, "Horizontal, ", row)
			sol += row * 100
		}

		if ok, col := reflectionPoint(transpose(p), smudge); ok {
			//fmt.Println(idx, "Vertical, ", col)
			sol += col
		}
	}
	return sol
}

func transpose(slice [][]rune) [][]rune {

	xl := len(slice[0])
	yl := len(slice)

	result := make([][]rune, xl)
	for i := range result {
		result[i] = make([]rune, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func reflectionPoint(pattern Pattern, smudge int) (bool, int) {

	for i := 0; i < len(pattern); i++ {
		if isReflection(pattern, i, smudge) {
			return true, i + 1
		}
	}

	return false, -1
}

func isReflection(pattern Pattern, row int, smudge int) bool {
	if row+1 >= len(pattern) {
		return false
	}

	ds := 0

	for d := 0; ; d++ {
		if row-d < 0 || row+d+1 >= len(pattern) {
			break
		}
		ds += distance(pattern[row-d], pattern[row+d+1])
	}

	return ds == smudge
}

func distance(a []rune, b []rune) int {
	ds := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			ds++
		}
	}
	return ds
}

func main() {
	solve("./2023/day13/input.txt")
}
