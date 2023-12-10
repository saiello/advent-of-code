package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

type P struct {
	i int
	j int
}

func parse(file string) (P, [][]rune) {
	tiles := [][]rune{}
	var start P

	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	for i := 0; scanner.Scan(); i++ {
		tiles = append(tiles, []rune{})
		for j, c := range scanner.Text() {
			tiles[len(tiles)-1] = append(tiles[len(tiles)-1], c)
			if c == 'S' {
				start = P{i: i, j: j}
			}
		}
	}

	return start, tiles
}

func part1(file string) int {
	start, tiles := parse(file)

	ao, bo := initmoves(start, tiles)

	a := P{i: start.i + ao.i, j: start.j + ao.j}
	b := P{i: start.i + bo.i, j: start.j + bo.j}

	sol1 := 1

	for {
		ao = nextmove(ao, tiles[a.i][a.j])
		bo = nextmove(bo, tiles[b.i][b.j])

		a = P{i: a.i + ao.i, j: a.j + ao.j}
		b = P{i: b.i + bo.i, j: b.j + bo.j}

		sol1 += 1

		if b.i == a.i && b.j == a.j {
			break
		}
	}

	return sol1
}

func part2(file string) int {

	start, tiles := parse(file)
	areas := [][]rune{}

	ao, bo := initmoves(start, tiles)
	a := P{i: start.i + ao.i, j: start.j + ao.j}
	b := P{i: start.i + bo.i, j: start.j + bo.j}

	for i := 0; i < len(tiles); i++ {
		areas = append(areas, []rune{})
		for j := 0; j < len(tiles[i]); j++ {
			areas[i] = append(areas[i], '.')
		}
	}

	areas[start.i][start.j] = tiles[start.i][start.j]
	areas[a.i][a.j] = tiles[a.i][a.j]
	areas[b.i][b.j] = tiles[b.i][b.j]

	for {
		ao = nextmove(ao, tiles[a.i][a.j])
		bo = nextmove(bo, tiles[b.i][b.j])

		a = P{i: a.i + ao.i, j: a.j + ao.j}
		b = P{i: b.i + bo.i, j: b.j + bo.j}

		areas[a.i][a.j] = tiles[a.i][a.j]
		areas[b.i][b.j] = tiles[b.i][b.j]

		if b.i == a.i && b.j == a.j {
			break
		}
	}

	sol2 := 0
	in := false
	last := '.'
	for i := 0; i < len(areas); i++ {
		for j := 0; j < len(areas[i]); j++ {
			switch areas[i][j] {
			case '.':
				if in {
					sol2 += 1
				}
			case '-':
				// do nothing
			case '|':
				in = !in
			case 'F':
				in = !in
				last = 'F'
			case 'L':
				in = !in
				last = 'L'
			case '7':
				if last != 'L' {
					in = !in
				}
			case 'J':
				if last != 'F' {
					in = !in
				}
			}
		}
	}

	return sol2

}

func initmoves(start P, tiles [][]rune) (P, P) {
	moves := []P{
		{i: 0, j: 1},  // -, J, 7
		{i: 1, j: 0},  // |, J, L
		{i: -1, j: 0}, // |, F, 7
		{i: 0, j: -1}, // -, L, F
	}

	allowed := [][]rune{
		{'-', 'J', '7'},
		{'|', 'J', 'L'},
		{'|', 'F', '7'},
		{'-', 'F', 'L'},
	}

	ans := []P{}

	for i := 0; i < len(moves); i++ {
		if slices.Contains(allowed[i], tiles[start.i+moves[i].i][start.j+moves[i].j]) {
			ans = append(ans, moves[i])
		}
	}

	return ans[0], ans[1]
}

func nextmove(prev P, curr rune) P {
	switch curr {
	case '|': // | is a vertical pipe connecting north and south.
		return prev
	case '-': // - is a horizontal pipe connecting east and west.
		return prev
	case 'L': // L is a 90-degree bend connecting north and east.
		return P{i: prev.j, j: prev.i}
	case 'J': // J is a 90-degree bend connecting north and west.
		return P{i: -prev.j, j: -prev.i}
	case '7': // 7 is a 90-degree bend connecting south and west.
		return P{i: prev.j, j: prev.i}
	case 'F': // F is a 90-degree bend connecting south and east.
		return P{i: -prev.j, j: -prev.i}
	}
	panic(1) // should never happen
}

func solve(file string) {
	fmt.Println("1: ", part1(file))
	fmt.Println("2: ", part2(file))
}

func main() {
	solve("./2023/day10/input.txt")
}
