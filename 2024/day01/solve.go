package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	lsa := []int{} // locations side A
	lsb := []int{} // locations side B

	// Part-1
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if loc, err := strconv.Atoi(fields[0]); err != nil {
			panic(err)
		} else {
			lsa = append(lsa, loc)
		}

		if loc, err := strconv.Atoi(fields[1]); err != nil {
			panic(err)
		} else {
			lsb = append(lsb, loc)
		}
	}

	sort.Ints(lsa)
	sort.Ints(lsb)

	totalDistance := 0
	for i := 0; i < len(lsa); i++ {
		if lsa[i] > lsb[i] {
			totalDistance += lsa[i] - lsb[i]
		} else {
			totalDistance += lsb[i] - lsa[i]
		}
	}

	fmt.Println(totalDistance)

	// Part-2
	lastloc, lastoccurrences, score := -1, 0, 0

	for i := 0; i < len(lsa); i++ {

		// Calculate occurrences only for locations different from previous
		if lastloc != lsa[i] {
			lastloc = lsa[i]
			lastoccurrences = 0

			for j := 0; j < len(lsa); j++ {
				if lsb[j] == lsa[i] {
					lastoccurrences += 1
				}
			}
		}

		// fmt.Printf("Loc %d occurres %d times\n", lsa[i], lastoccurrences)
		score += lsa[i] * lastoccurrences
	}
	fmt.Println(score)
}
