package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	bid      int
	cards    string
	cmpCards string
	point    int
}

func HandCompare(a Hand, b Hand) int {
	if a.point > b.point {
		return 1
	} else if a.point < b.point {
		return -1
	} else if a.cmpCards > b.cmpCards {
		return 1
	} else if a.cmpCards < b.cmpCards {
		return -1
	}
	return 0
}

var replaces = map[string]string{
	"A": "F",
	"K": "E",
	"Q": "D",
	"J": "C",
	"T": "B",
}

var replacesWithJoker = map[string]string{
	"A": "F",
	"K": "E",
	"Q": "D",
	"J": "1",
	"T": "B",
}

func replaceAll(arg string, replaces map[string]string) string {
	for k, v := range replaces {
		arg = strings.ReplaceAll(arg, k, v)
	}
	return arg
}

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
// A, K, Q, T, 9, 8, 7, 6, 5, 4, 3, 2, J

func Point(hand string) int {
	cards := map[string]int{}

	for _, r := range hand {
		cards[fmt.Sprintf("%c", r)] += 1
	}

	p0, p1 := 0, 0

	for _, v := range cards {
		if v > p0 {
			p1 = p0
			p0 = v
		} else if v > p1 {
			p1 = v
		}
	}
	point, _ := strconv.Atoi(fmt.Sprintf("%d%d", p0, p1))
	return point
}

func PointWithJoker(hand string) int {
	cards := map[string]int{}

	for _, r := range hand {
		cards[fmt.Sprintf("%c", r)] += 1
	}

	p0, p1, j := 0, 0, 0

	for t, v := range cards {
		if t == "J" {
			j = v
		} else if v > p0 {
			p1 = p0
			p0 = v
		} else if v > p1 {
			p1 = v
		}
	}

	p0 = p0 + j

	point, _ := strconv.Atoi(fmt.Sprintf("%d%d", p0, p1))
	return point
}

func solve(file string) {
	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	sol1 := 0
	sol2 := 0

	hands, handsWithJoker := []Hand{}, []Hand{}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		bid, _ := strconv.Atoi(fields[1])

		hands = append(hands, Hand{cards: fields[0], point: Point(fields[0]), bid: bid, cmpCards: replaceAll(fields[0], replaces)})

		handsWithJoker = append(handsWithJoker, Hand{cards: fields[0], point: PointWithJoker(fields[0]), bid: bid, cmpCards: replaceAll(fields[0], replacesWithJoker)})
	}

	slices.SortFunc(hands, HandCompare)
	slices.SortFunc(handsWithJoker, HandCompare)

	for rank, hand := range hands {
		sol1 += hand.bid * (rank + 1)
	}

	for rank, hand := range handsWithJoker {
		sol2 += hand.bid * (rank + 1)
	}

	fmt.Println("1: ", sol1)
	fmt.Println("2: ", sol2)
}

func main() {
	solve("./2023/day07/input.txt")
}
