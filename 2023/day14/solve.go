package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

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

func parse(file string) [][]rune {
	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	runes := [][]rune{}

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		runes = append(runes, make([]rune, len(line)))
		for j, r := range line {
			runes[i][j] = r
		}
	}
	return runes
}

func part1(file string) int {

	sol := 0

	runes := transpose(parse(file))

	for i := 0; i < len(runes); i++ {
		lineLoad := 0
		val := len(runes[i])

		for j := 0; j < len(runes[i]); j++ {
			// fmt.Printf("%c", runes[i][j])
			switch runes[i][j] {
			case 'O':
				lineLoad += val
				val--
			case '#':
				val = len(runes[i]) - j - 1
			case '.':
				// do nothing
			}
		}
		// fmt.Printf("  -> Load: %d\n", lineLoad)
		sol += lineLoad
	}

	return sol
}

func shift(runes []rune) {

	//runes := *runesPt
	head := 0
	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case 'O':
			runes[i] = '.'
			runes[head] = 'O'
			head++
		case '#':
			head = i + 1
		case '.':
			// do nothing
		}
	}
}

func shiftR(runes []rune) {

	//runes := *runesPt
	head := len(runes) - 1
	for i := head; i >= 0; i-- {
		switch runes[i] {
		case 'O':
			runes[i] = '.'
			runes[head] = 'O'
			head--
		case '#':
			head = i - 1
		case '.':
			// do nothing
		}
	}
}

func tiltR(runes [][]rune) {
	for i := 0; i < len(runes); i++ {
		shiftR(runes[i])
	}
}

func tilt(runes [][]rune) {
	for i := 0; i < len(runes); i++ {
		shift(runes[i])
	}
}

func cycle(runes [][]rune) [][]rune {
	runes = transpose(runes)
	tilt(runes) // north
	runes = transpose(runes)
	tilt(runes) // west
	runes = transpose(runes)
	tiltR(runes) // south
	runes = transpose(runes)
	tiltR(runes) // east

	return runes
}

func load(runes [][]rune) int {
	sol := 0

	for i := 0; i < len(runes); i++ {
		lineLoad := 0
		for j := 0; j < len(runes[i]); j++ {
			//fmt.Printf("%c", runes[i][j])
			switch runes[i][j] {
			case 'O':
				lineLoad += len(runes[i]) - j
			case '#':
				// do nothing
			case '.':
				// do nothing
			}
		}
		//fmt.Printf("  -> Load: %d\n", lineLoad)
		sol += lineLoad
	}
	return sol
}

func RunesAsString(m [][]rune) string {
	s := strings.Builder{}
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[r]); c++ {
			s.WriteRune(m[r][c])
		}
	}
	return s.String()
}

func part2(file string) int {
	runes := parse(file)

	visited := map[string]int{}

	// var firstSeen int
	// var epoch int

	goal := 1000000000

	i := 0

	for i < goal {

		runes = cycle(runes)
		key := RunesAsString(runes)

		//fmt.Println(i, key, load(runes))

		if seen, ok := visited[key]; ok {

			epochLen := i - seen
			epochNum := (goal - i) / epochLen

			//fmt.Println(i, "seen", seen, "epochLen", epochLen, epochNum, ((epochNum * epochLen) + 1 + i))

			skip := (epochNum * epochLen) + 1
			if i+skip < goal {
				i += skip
			} else {
				i++
			}

		} else {
			visited[key] = i
			i++
		}
	}

	// last transpose for the north
	runes = transpose(runes)

	return load(runes)
}

func solve(file string) {
	fmt.Println("1: ", part1(file))
	fmt.Println("2: ", part2(file))
}

func main() {
	solve("./2023/day14/input.txt")
}
