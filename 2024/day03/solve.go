package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	r  = regexp.MustCompile(`(mul\(([0-9]+),([0-9]+)\))+?`)
	di = regexp.MustCompile(`(don't\(\).+?do\(\))`)
	de = regexp.MustCompile(`(don't\(\).+?$)`)
	ds = regexp.MustCompile(`(^.+?do\(\))`)
)

func executeMul(line string) int {
	var sum int
	m := r.FindAllStringSubmatch(line, -1)
	for i := 0; i < len(m); i++ {
		a, _ := strconv.Atoi(m[i][2])
		b, _ := strconv.Atoi(m[i][3])
		sum += a * b
	}
	return sum
}

func removeDisabled(line string, enabled bool) (string, bool) {

	// If not enabled remove up to the first do()
	if !enabled && ds.MatchString(line) {
		line = ds.ReplaceAllString(line, "")
	}

	// Remove inner dont()-do() pairs
	line = di.ReplaceAllString(line, "")

	// Keep track of last dont()
	if enabled = !de.MatchString(line); !enabled {
		line = de.ReplaceAllString(line, "")
	}

	return line, enabled
}

func main() {

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	var sum, sum2 int

	enabled := true

	for scanner.Scan() {
		line := scanner.Text()

		// Part-1
		sum += executeMul(line)

		// Part-2
		line, enabled = removeDisabled(line, enabled)
		sum2 += executeMul(line)
	}
	fmt.Println(sum)  // 187194524
	fmt.Println(sum2) // 127092535
}
