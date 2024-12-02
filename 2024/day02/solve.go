package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	count, dampenerCount := 0, 0

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		report := []int{}

		for i := 0; i < len(fields); i++ {
			if level, err := strconv.Atoi(fields[i]); err != nil {
				panic(err)
			} else {
				report = append(report, level)
			}
		}

		if isSafe(report) {
			count += 1
		} else if isSafeWithDampener(report) {
			dampenerCount += 1
		}
	}

	fmt.Println(count)
	fmt.Println(count + dampenerCount)

}

func isSafe(report []int) bool {
	if report[0] == report[1] {
		return false
	} else if report[0] < report[1] {
		// increasing

		for i := 0; i < len(report)-1; i++ {
			levelDiff := report[i+1] - report[i]

			if levelDiff < 1 || levelDiff > 3 {
				return false
			}
		}

	} else {

		for i := 0; i < len(report)-1; i++ {
			levelDiff := report[i] - report[i+1]

			if levelDiff < 1 || levelDiff > 3 {
				return false
			}
		}
	}

	return true
}

func deleteElement(slice []int, index int) []int {
	out := []int{}
	out = append(out, slice[:index]...)
	out = append(out, slice[index+1:]...)
	return out
}

func isSafeWithDampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		if isSafe(deleteElement(report, i)) {
			return true
		}
	}
	return false
}
