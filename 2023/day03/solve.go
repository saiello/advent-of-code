package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solveFile(file string) {
	input, _ := os.ReadFile(file)
	sol1, sol2 := Solve(input)
	fmt.Println("1: ", sol1)
	fmt.Println("2: ", sol2)
}

func Solve(input []byte) (int, int) {

	scanner := bufio.NewScanner(bytes.NewReader(input))

	sol1, sol2 := 0, 0

	dr, sr, gr := regexp.MustCompile(`([0-9]+)`), regexp.MustCompile(`([^\.0-9])`), regexp.MustCompile(`([\*])`)

	dIdxs, sIdxs, gIdxs, ds := [][][]int{}, [][][]int{}, [][][]int{}, [][]string{}

	for scanner.Scan() {
		bytes := scanner.Bytes()

		// ................410....&..........972.......................$..305..683..743........551.338.&..................................*............
		// fmt.Printf("%s\n", bytes)

		di := dr.FindAllIndex(bytes, -1)
		d := dr.FindAllString(fmt.Sprintf("%s", bytes), -1)
		si := sr.FindAllIndex(bytes, -1)
		gi := gr.FindAllIndex(bytes, -1)

		// Digits:  [410 972 305 683 743 551 338]
		// Digits(indexes):  [[16 19] [34 37] [63 66] [68 71] [73 76] [84 87] [88 91]]
		// Symbols(indexes):  [[23 24] [60 61] [92 93] [127 128]]
		// fmt.Println("Digits: ", d)
		// fmt.Println("Digits(indexes): ", di)
		// fmt.Println("Symbols(indexes): ", si)

		ds = append(ds, d)
		dIdxs = append(dIdxs, di)
		sIdxs = append(sIdxs, si)
		gIdxs = append(gIdxs, gi)

	}

	// Part 1
	for i := 0; i < len(ds); i++ {
		// each number
		for j := 0; j < len(ds[i]); j++ {
			// each index aka range
			num, _ := strconv.Atoi(ds[i][j])
			if isRangeAdiacent(sIdxs, dIdxs[i][j], i) {
				sol1 += num
			}
		}
	}

	// Part 2
	for gearLineIndex, gearLine := range gIdxs {
		// each potential gear
		for _, gearPositions := range gearLine {
			// each index aka range
			partNumbers := getAdiacentNumbers(gearPositions, dIdxs, ds, gearLineIndex)
			if len(partNumbers) == 2 {
				//fmt.Println("Gear Pos", gearPositions, "Part", partNumbers)
				// is a Gear
				sol2 += (partNumbers[0] * partNumbers[1])
			}
		}
	}

	return sol1, sol2
}

func getAdiacentNumbers(gearPositions []int, numbersIndexes [][][]int, numbers [][]string, gearLineIndex int) []int {
	partNumbers := []int{}
	gearStart, gearEnd := gearPositions[0], gearPositions[1]

	lineOfNumberIndexes := numbersIndexes[gearLineIndex]

	for i, numberIndexes := range lineOfNumberIndexes {
		if gearStart == numberIndexes[1] || gearEnd == numberIndexes[0] {
			num, _ := strconv.Atoi(numbers[gearLineIndex][i])
			partNumbers = append(partNumbers, num)
		}
	}

	if gearLineIndex-1 >= 0 {
		lineOfNumberIndexes := numbersIndexes[gearLineIndex-1]
		for i, numberIndexes := range lineOfNumberIndexes {
			if gearStart >= numberIndexes[0]-1 && gearStart <= numberIndexes[1] {
				num, _ := strconv.Atoi(numbers[gearLineIndex-1][i])
				partNumbers = append(partNumbers, num)
			}
		}
	}

	if gearLineIndex+1 < len(numbers) {
		lineOfNumberIndexes := numbersIndexes[gearLineIndex+1]
		for i, numberIndexes := range lineOfNumberIndexes {
			if gearStart >= numberIndexes[0]-1 && gearStart <= numberIndexes[1] {
				num, _ := strconv.Atoi(numbers[gearLineIndex+1][i])
				partNumbers = append(partNumbers, num)
			}
		}
	}

	return partNumbers
}

func isRangeAdiacent(symbolsIndexes [][][]int, numberIndexes []int, i int) bool {
	for z := numberIndexes[0]; z <= numberIndexes[1]; z++ {
		if existsSymbol(symbolsIndexes[i], z) {
			return true
		}
		if i > 1 && existsSymbol(symbolsIndexes[i-1], z) {
			return true
		}
		if i+1 < len(symbolsIndexes)-1 && existsSymbol(symbolsIndexes[i+1], z) {
			return true
		}
	}
	return false
}

func existsSymbol(indexesOfLineOfSymbols [][]int, index int) bool {
	//fmt.Println("DigitIndex", index, indexesOfLineOfSymbols)
	for _, idx := range indexesOfLineOfSymbols {
		for j := idx[0]; j <= idx[1]; j++ {
			if j == index {
				return true
			}
		}
	}
	return false
}

func main() {
	solveFile("./2023/day03/input.txt")
}
