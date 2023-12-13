package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

type P struct {
	i int
	j int
}

type Range struct {
	from int
	to   int
}

func NewRange(a int, b int) Range {
	if a < b {
		return Range{from: a, to: b}
	} else {
		return Range{from: b, to: a}
	}
}

func (r *Range) Contains(arg int) bool {
	return r.from < arg && r.to > arg
}

func parse(file string) ([]P, [][]rune) {
	tiles := [][]rune{}
	galaxies := []P{}

	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	for i := 0; scanner.Scan(); i++ {
		tiles = append(tiles, []rune{})
		for j, c := range scanner.Text() {
			tiles[len(tiles)-1] = append(tiles[len(tiles)-1], c)
			if c == '#' {
				galaxies = append(galaxies, P{i: i, j: j})
			}
		}
	}

	return galaxies, tiles
}

func findEmptyRows(tiles [][]rune) []int {
	ans := []int{}
	for i := 0; i < len(tiles); i++ {
		numgal := 0
		for j := 0; j < len(tiles[i]); j++ {
			if tiles[i][j] == '#' {
				numgal++
			}
		}
		if numgal == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func findEmptyCols(tiles [][]rune) []int {
	ans := []int{}
	for j := 0; j < len(tiles[0]); j++ {
		numgal := 0
		for i := 0; i < len(tiles); i++ {
			if tiles[i][j] == '#' {
				numgal++
			}
		}
		if numgal == 0 {
			ans = append(ans, j)
		}
	}
	return ans
}

func findDistances(file string, expansion int) int {
	galaxies, tiles := parse(file)

	emptyCols, emptyRows := findEmptyCols(tiles), findEmptyRows(tiles)

	sol := 0

	for g := 0; g < len(galaxies); g++ {
		for h := 0; h < len(galaxies); h++ {
			if h > g {

				src, dest := galaxies[g], galaxies[h]

				vrange, hrange := NewRange(src.i, dest.i), NewRange(src.j, dest.j)

				doubled := 0

				for _, e := range emptyRows {
					if vrange.Contains(e) {
						doubled += expansion - 1
					}
				}

				for _, e := range emptyCols {
					if hrange.Contains(e) {
						doubled += expansion - 1
					}
				}

				dist := int64(math.Abs(float64(src.i-dest.i))+
					math.Abs(float64(src.j-dest.j))) + int64(doubled)

				sol += int(dist)
			}
		}
	}

	return sol
}

func solve(file string) {
	fmt.Println("1: ", findDistances(file, 2))
	fmt.Println("2: ", findDistances(file, 1000000))
}

func main() {
	solve("./2023/day11/input.txt")
}
