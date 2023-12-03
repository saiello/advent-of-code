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

var replaces = map[string]string{
	"one":   "on1ne",
	"two":   "tw2wo",
	"three": "thr3ee",
	"four":  "fo4ur",
	"five":  "fi5ve",
	"six":   "si6ix",
	"seven": "sev7ven",
	"eight": "eig8ght",
	"nine":  "ni9ne",
}

func sanitize(arg string) string {
	for k, v := range replaces {
		arg = strings.ReplaceAll(arg, k, v)
	}
	return arg
}

func main() {

	input, _ := os.ReadFile("./2023/day01/input.txt")
	scanner := bufio.NewScanner(bytes.NewReader(input))

	r, _ := regexp.Compile(`[0-9]`)

	sum1 := 0
	sum2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		// part 1
		m1 := r.FindAllString(line, -1)
		num, _ := strconv.Atoi(m1[0] + m1[len(m1)-1])
		sum1 += num

		// part 2
		m2 := r.FindAllString(sanitize(line), -1)
		num, _ = strconv.Atoi(m2[0] + m2[len(m2)-1])

		sum2 += num
	}

	fmt.Println("1: ", sum1)
	fmt.Println("2: ", sum2)
}
