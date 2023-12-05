package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// ------------
// Model
// ------------

var maps = []CMap{}

type CMap struct {
	name   string
	ranges Ranges
}

func (m *CMap) getDestinationCategory(sourceCategory int) int {
	for _, r := range m.ranges {
		if sourceCategory >= r.source && sourceCategory < r.source+r.lenght {
			return sourceCategory - r.source + r.destination
		}
	}
	return sourceCategory
}

type Ranges []Range

type Range struct {
	destination int
	source      int
	lenght      int
}

// ------------
// Util
// ------------

func mustBeInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic("Unexpected not integer string")
	}

	return i
}

func solve(file string) {
	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	sol1 := -1
	sol2 := -1

	seeds := []int{}

	// ------------
	// Parsing
	// ------------
	for scanner.Scan() {
		line := scanner.Text()

		if reg := regexp.MustCompile(`^seeds: ([0-9]*\s*)*`); reg.MatchString(line) {
			submatches := regexp.MustCompile(`[0-9]+`).FindAllStringSubmatch(line, -1)
			for _, matches := range submatches {
				seed, _ := strconv.Atoi(matches[0])
				seeds = append(seeds, seed)
			}

		} else if reg := regexp.MustCompile(`^(.+) map:`); reg.MatchString(line) {
			m := reg.FindAllStringSubmatch(line, 1)
			maps = append(maps, CMap{name: m[0][1]})

		} else if reg := regexp.MustCompile(`^([0-9]+)\s+([0-9]+)\s+([0-9]+)$`); reg.MatchString(line) {
			m := reg.FindStringSubmatch(line)

			maps[len(maps)-1].ranges = append(maps[len(maps)-1].ranges, Range{
				destination: mustBeInt(m[1]),
				source:      mustBeInt(m[2]),
				lenght:      mustBeInt(m[3]),
			})
		}

	}

	// ------------
	// Solving
	// ------------
	fmt.Println("Seeds", seeds, "Maps", maps)

	// Part 1
	for _, sourceCategory := range seeds {
		// convertions := []string{}
		category := sourceCategory
		// convertions = append(convertions, fmt.Sprintf("%s:%d", "Seed", category))
		for _, rangeMap := range maps {
			category = rangeMap.getDestinationCategory(category)
			// convertions = append(convertions, fmt.Sprintf("%s:%d", rangeMap.name, category))
		}
		// fmt.Println(convertions)
		// fmt.Println("Seed", sourceCategory, "Location", category)
		if sol1 == -1 || category < sol1 {
			sol1 = category
		}
	}

	// Part 2
	for i := 0; i < len(seeds)-1; i += 2 {
		startSeed := seeds[i]
		endSeed := startSeed + seeds[i+1]

		for sourceCategory := startSeed; sourceCategory < endSeed; sourceCategory++ {

			category := sourceCategory
			for _, rangeMap := range maps {
				category = rangeMap.getDestinationCategory(category)
			}
			if sol2 == -1 || category < sol2 {
				sol2 = category
			}
		}
	}

	fmt.Println("1: ", sol1)
	fmt.Println("2: ", sol2)
}

func main() {
	solve("./2023/day05/input.txt")
}
