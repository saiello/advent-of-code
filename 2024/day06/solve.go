package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Pos struct {
	r int
	c int
}

type Step struct {
	p Pos
	d Direction
}

type Offset = Pos

var Moves = map[Direction]Offset{
	Up:    {r: -1, c: 0},
	Right: {r: 0, c: 1},
	Down:  {r: 1, c: 0},
	Left:  {r: 0, c: -1},
}

func start(tiles [][]rune) Pos {
	for r := 0; r < len(tiles); r++ {
		for c := 0; c < len(tiles[r]); c++ {
			if tiles[r][c] == '^' {
				return Pos{r: r, c: c}
			}
		}
	}

	panic("No starting position found")
}

func main() {

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	tiles := [][]rune{}

	for scanner.Scan() {
		tiles = append(tiles, []rune(scanner.Text()))
	}

	startPos := start(tiles)

	// Part-1
	var (
		direction = Up
		pos       = startPos
		visited   = map[Pos]bool{}
	)

	for {
		visited[pos] = true

		newPos := Pos{r: pos.r + Moves[direction].r, c: pos.c + Moves[direction].c}

		if newPos.r < 0 || newPos.r >= len(tiles) || newPos.c < 0 || newPos.c >= len(tiles[0]) {
			break
		}

		if tiles[newPos.r][newPos.c] == '#' {
			direction = (direction + 1) % 4
		} else {
			pos = newPos
		}

	}

	fmt.Println(len(visited))

	// Part-2
	var sum int

	for obstacle := range visited {

		var (
			direction = Up
			pos       = startPos
			steps     = map[Step]bool{}
		)

		for {
			s := Step{p: pos, d: direction}

			if steps[s] {
				sum++
				break
			}

			steps[s] = true

			newPos := Pos{r: pos.r + Moves[direction].r, c: pos.c + Moves[direction].c}

			if newPos.r < 0 || newPos.r >= len(tiles) || newPos.c < 0 || newPos.c >= len(tiles[0]) {
				break
			}

			if tiles[newPos.r][newPos.c] == '#' || newPos == obstacle {
				direction = (direction + 1) % 4
			} else {
				pos = newPos
			}
		}
	}

	fmt.Println(sum)
}
