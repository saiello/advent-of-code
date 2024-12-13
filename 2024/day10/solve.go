package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func (p Pos) Move(o Offset) Pos {
	return Pos{r: p.r + o.r, c: p.c + o.c}
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

func forEachTile(tiles [][]int, consumer func(p Pos, v int)) {
	for r := 0; r < len(tiles); r++ {
		for c := 0; c < len(tiles[0]); c++ {
			consumer(Pos{r: r, c: c}, tiles[r][c])
		}
	}
}

func forEachPair(tiles [][]int, consumer func(p1 Pos, v1 int, p2 Pos, v2 int)) {

	forEachTile(tiles, func(p Pos, v int) {
		for _, o := range Moves {
			pp := p.Move(o)
			if pp.r >= 0 && pp.r < len(tiles) && pp.c >= 0 && pp.c < len(tiles[0]) {
				consumer(p, v, pp, tiles[pp.r][pp.c])
			}
		}
	})
}

func countReacheablePeaks(p Pos, edges map[Pos][]Pos, trailpeaks map[Pos]bool, visited map[Pos]bool) int {

	if visited[p] {
		return 0
	}

	sum := 0
	visited[p] = true

	if trailpeaks[p] {
		return 1
	}

	if tos, ok := edges[p]; ok {
		for i := 0; i < len(tos); i++ {
			sum += countReacheablePeaks(tos[i], edges, trailpeaks, visited)
		}
	}
	return sum
}

func countDistinctTrails(p Pos, edges map[Pos][]Pos, trailpeaks map[Pos]bool) int {

	var distinctPaths = map[string]bool{}

	var inner func(p Pos, pathPrefix []Pos)
	inner = func(p Pos, pathPrefix []Pos) {

		if trailpeaks[p] {
			text := ""
			for _, pp := range append([]Pos{p}, pathPrefix...) {
				text += fmt.Sprintf("%d,%d|", pp.r, pp.c)
			}
			distinctPaths[text] = true
		}

		if tos, ok := edges[p]; ok {
			for i := 0; i < len(tos); i++ {
				//fmt.Println(p, "--> ", tos)
				inner(tos[i], append([]Pos{p}, pathPrefix...))
			}
		}
	}

	inner(p, []Pos{})

	return len(distinctPaths)
}

func main() {

	var (
		tiles      = [][]int{}
		edges      = map[Pos][]Pos{}
		trailheads = map[Pos]bool{}
		trailpeaks = map[Pos]bool{}
	)

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	for scanner.Scan() {
		row := []int{}
		for _, s := range scanner.Text() {
			i, _ := strconv.Atoi(string(s))
			row = append(row, i)
		}
		tiles = append(tiles, row)
	}
	forEachTile(tiles, func(p Pos, v int) {
		if v == 0 {
			trailheads[p] = true
		} else if v == 9 {
			trailpeaks[p] = true
		}
	})

	forEachPair(tiles, func(p1 Pos, v1 int, p2 Pos, v2 int) {
		if v1+1 == v2 {

			if _, ok := edges[p1]; !ok {
				edges[p1] = []Pos{}
			}

			edges[p1] = append(edges[p1], p2)
		}
	})

	// Part-1
	sum := 0
	for head := range trailheads {
		peaks := countReacheablePeaks(head, edges, trailpeaks, map[Pos]bool{})
		sum += peaks
	}
	fmt.Println(sum)

	// Part-2
	rating := 0
	for head := range trailheads {
		rating += countDistinctTrails(head, edges, trailpeaks)
	}
	fmt.Println(rating)

}
