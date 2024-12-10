package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readBlocks(diskmap []int) []int {
	blocks := []int{}

	for i := 0; i < len(diskmap); i++ {
		for b := 0; b < diskmap[i]; b++ {
			id := -1
			if (i % 2) == 0 {
				id = i / 2
			}
			blocks = append(blocks, id)
		}
	}
	return blocks
}

func compact(blocks []int) {
	last := len(blocks) - 1
	seek := 0

	for {
		if blocks[last] == -1 {
			last--
		} else if blocks[seek] == -1 {
			blocks[seek] = blocks[last]
			blocks[last] = -1
			seek++
			last--
		} else {
			seek++
		}
		if last < seek {
			break
		}
	}
}

func checksum(blocks []int) int {
	var sum int
	for i, id := range blocks {
		if id != -1 {
			sum += i * id
		}

	}
	return sum
}

func main() {
	var diskmap = []int{}

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	for scanner.Scan() {
		for _, s := range strings.Split(scanner.Text(), "") {
			i, _ := strconv.Atoi(s)
			diskmap = append(diskmap, i)
		}
	}

	// Part-1
	blocks := readBlocks(diskmap)
	compact(blocks)
	fmt.Println(checksum(blocks))
}
