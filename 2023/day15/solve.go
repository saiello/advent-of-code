package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func hash(text string) int {
	hash := 0
	for _, c := range text {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}

type Len struct {
	label string
	focal int
}

func matcher(needle Len) func(Len) bool {
	return func(l Len) bool { return l.label == needle.label }
}

var reg = regexp.MustCompile(`([a-z]+)([=-])([0-9]*)?`)

func parse(text string) (string, Len) {

	if reg.MatchString(text) {
		all := reg.FindAllStringSubmatch(text, -1)
		label, op := all[0][1], all[0][2]
		focal, _ := strconv.Atoi(all[0][len(all[0])-1])

		return op, Len{label: label, focal: focal}
	}

	return "", Len{}
}

func part1(file string) int {
	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	sol := 0

	for scanner.Scan() {
		line := scanner.Text()
		for _, text := range strings.Split(line, ",") {
			sol += hash(text)
		}
	}

	return sol
}

func part2(file string) int {
	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	sol := 0

	type Box []Len

	boxes := make([]Box, 256)

	for scanner.Scan() {
		line := scanner.Text()
		for _, text := range strings.Split(line, ",") {

			op, flen := parse(text)
			box := &boxes[hash(flen.label)]

			switch op {
			case "=":
				i := slices.IndexFunc(*box, matcher(flen))
				if i >= 0 {
					(*box)[i] = flen
				} else {
					*box = append(*box, flen)
				}
			case "-":
				*box = slices.DeleteFunc(*box, matcher(flen))
			}
		}
	}

	for bi, box := range boxes {
		for li, ln := range box {
			sol += (bi + 1) * (li + 1) * ln.focal
		}
	}
	return sol
}

func solve(file string) {
	fmt.Println("1: ", part1(file))
	fmt.Println("2: ", part2(file))
}

func main() {
	solve("./2023/day15/input.txt")
}
