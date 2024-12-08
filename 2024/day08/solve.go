package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	r int
	c int
}

type Offset Pos

type Antenna struct {
	frequency rune
	pos       Pos
}

func delta(p1 Pos, p2 Pos) Offset {
	return Offset{r: p1.r - p2.r, c: p1.c - p2.c}
}

func (p Pos) Add(o Offset) Pos {
	return Pos{r: p.r + o.r, c: p.c + o.c}
}

func (p Pos) Subtract(o Offset) Pos {
	return Pos{r: p.r - o.r, c: p.c - o.c}
}

func main() {
	var tiles = [][]rune{}
	var antennas = []Antenna{}

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		tiles = append(tiles, []rune(scanner.Text()))
	}

	for r := 0; r < len(tiles); r++ {
		for c := 0; c < len(tiles[r]); c++ {
			if a := tiles[r][c]; a != '.' {
				antennas = append(antennas, Antenna{frequency: a, pos: Pos{r: r, c: c}})
			}
		}
	}

	var withinBounds = func(p Pos) bool {
		return p.r >= 0 && p.r < len(tiles) && p.c >= 0 && p.c < len(tiles[0])
	}

	var pairOfAntennas = func(consumer func(a1 Antenna, a2 Antenna)) {
		for i, a1 := range antennas {
			for j, a2 := range antennas {
				if i > j && a1.frequency == a2.frequency {
					consumer(a1, a2)
				}
			}
		}
	}

	// Part-1
	var antinodes = map[Pos]bool{}

	pairOfAntennas(func(a1, a2 Antenna) {
		dt := delta(a1.pos, a2.pos)
		if a := a1.pos.Add(dt); withinBounds(a) {
			antinodes[a] = true
		}
		if a := a2.pos.Subtract(dt); withinBounds(a) {
			antinodes[a] = true
		}
	})

	fmt.Println(len(antinodes))

	// Part-2
	antinodes = map[Pos]bool{}

	pairOfAntennas(func(a1, a2 Antenna) {

		dt := delta(a1.pos, a2.pos)
		p := a1.pos

		for {
			antinodes[p] = true

			if a := p.Add(dt); withinBounds(a) {
				p = a
			} else {
				break
			}
		}

		p = a2.pos
		for {
			antinodes[p] = true

			if a := p.Subtract(dt); withinBounds(a) {
				p = a
			} else {
				break
			}
		}

	})

	fmt.Println(len(antinodes))
}
