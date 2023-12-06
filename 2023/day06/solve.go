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

type Race struct {
	duration int
	record   int
}

func mustBeInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic("Unexpected not integer string")
	}

	return i
}

func mustBeIntArray(arg []string) []int {

	out := []int{}
	for _, s := range arg {
		out = append(out, mustBeInt(s))
	}
	return out
}

func solve(file string) {
	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	sol1 := 1
	sol2 := 0

	numreg := regexp.MustCompile(`([0-9]+)`)

	var times []int
	var distances []int

	var bigRaceTime int
	var bigRaceDistance int

	// ---------
	// Parse
	// ---------
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if reg := regexp.MustCompile(`^Time:`); reg.MatchString(line) {
			matches := numreg.FindAllString(line, -1)
			times = mustBeIntArray(matches)
			bigRaceTime = mustBeInt(strings.Join(matches, ""))
		}

		if reg := regexp.MustCompile(`^Distance:`); reg.MatchString(line) {
			matches := numreg.FindAllString(line, -1)
			distances = mustBeIntArray(matches)
			bigRaceDistance = mustBeInt(strings.Join(matches, ""))
		}
	}

	fmt.Println("times: ", times)
	fmt.Println("distances: ", distances)
	fmt.Println("bigRaceTime: ", bigRaceTime)
	fmt.Println("bigRaceDistance: ", bigRaceDistance)

	// Part 1
	numberOfWays := []int{}
	for race := 0; race < len(times); race++ {
		raceTime, raceRecord := times[race], distances[race]
		waysToWin := findWayToWin(raceTime, raceRecord)
		numberOfWays = append(numberOfWays, len(waysToWin))
		sol1 *= len(waysToWin)
	}

	fmt.Println("numberOfWays: ", numberOfWays)
	fmt.Println("1: ", sol1)

	// Part 2
	sol2 = len(findWayToWin(bigRaceTime, bigRaceDistance))
	fmt.Println("2: ", sol2)
}

func findWayToWin(raceTime int, raceRecord int) []int {

	waysToWin := []int{}
	for pressTime := 0; pressTime < raceTime; pressTime++ {
		remainingTime := raceTime - pressTime
		speed := pressTime
		distance := speed * remainingTime
		if distance > raceRecord {
			waysToWin = append(waysToWin, pressTime)
		}
	}
	return waysToWin
}

func main() {
	solve("./2023/day06/input.txt")
}
