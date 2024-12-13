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

	var stones = []int{}

	scanner.Scan()
	for _, f := range strings.Fields(scanner.Text()) {
		i, _ := strconv.Atoi(f)
		stones = append(stones, i)
	}

	var blink = func() {

		newstones := []int{}

		for _, s := range stones {
			if s == 0 {
				newstones = append(newstones, 1)
			} else if ss := strconv.Itoa(s); len(ss)%2 == 0 {
				s1, _ := strconv.Atoi(ss[:len(ss)/2])
				s2, _ := strconv.Atoi(ss[len(ss)/2:])
				newstones = append(newstones, s1, s2)
			} else {
				newstones = append(newstones, s*2024)
			}
		}

		stones = newstones
	}

	for i := 0; i < 75; i++ {
		blink()
		fmt.Println(i, ">>", len(stones))
	}

	fmt.Println(len(stones))
}
