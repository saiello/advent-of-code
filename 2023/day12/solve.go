package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Row struct {
	record  string
	damaged []int
}

func part1(file string) int {

	springs := []Row{}

	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	for scanner.Scan() {
		f := strings.Fields(scanner.Text())

		damaged := []int{}
		ds := strings.Split(f[1], ",")
		for _, d := range ds {
			i, _ := strconv.Atoi(d)
			damaged = append(damaged, i)
		}

		springs = append(springs, Row{
			record:  f[0],
			damaged: damaged,
		})
	}

	return countValidArrangements(springs)
}

func part2(file string) int {

	springs := []Row{}

	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	for scanner.Scan() {
		f := strings.Fields(scanner.Text())

		rs := make([]string, 5)
		ds := make([]string, 5)
		for i := 0; i < 5; i++ {
			rs[i] = f[0]
			ds[i] = f[1]
		}

		damaged := []int{}
		for _, d := range strings.Split(strings.Join(ds, ","), ",") {
			i, _ := strconv.Atoi(d)
			damaged = append(damaged, i)
		}

		springs = append(springs, Row{
			record:  strings.Join(rs, ","),
			damaged: damaged,
		})
	}

	return countValidArrangements(springs)
}

func countValidArrangements(springs []Row) int {

	sol := 0

	symbols := []string{".", "#"}

	for _, r := range springs {
		valid := 0

		reg := fmt.Sprintf("^\\.*#{%d}", r.damaged[0])
		for i := 1; i < len(r.damaged); i++ {
			reg += fmt.Sprintf("\\.+#{%d}", r.damaged[i])
		}
		reg += "\\.*$"

		count := strings.Count(r.record, "?")
		fmt.Println("Make all possibilities", r.record, r.damaged, reg)
		replaces := possibleStrings(symbols, count)

		for _, replace := range replaces {
			record := r.record
			for _, repSym := range replace {
				record = strings.Replace(record, "?", fmt.Sprintf("%c", repSym), 1)
			}

			if regexp.MustCompile(reg).MatchString(record) {
				//fmt.Println(record)
				valid++
			}
		}

		fmt.Println(valid)
		sol += valid
	}
	return sol
}

func possibleStrings(symbols []string, lenght int) []string {
	out := []string{}

	if lenght == 1 {
		return symbols
	}

	for _, sy := range symbols {
		for _, o := range possibleStrings(symbols, lenght-1) {
			out = append(out, sy+o)
		}
	}

	return out
}

func solve(file string) {
	//fmt.Println("1: ", part1(file))
	fmt.Println("2: ", part2(file))
}

func main() {
	solve("./2023/day12/example.txt")
}
