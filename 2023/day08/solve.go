package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	label string
	left  string
	right string
}

func solve(file string) {
	input, _ := os.ReadFile(file)
	scanner := bufio.NewScanner(bytes.NewReader(input))

	sol1 := 0
	sol2 := 0

	reg := regexp.MustCompile(`(.+) = \((.+), (.+)\)`)

	scanner.Scan()
	directions := []bool{}

	for _, d := range scanner.Text() {
		directions = append(directions, fmt.Sprintf("%c", d) == "L")
	}

	fmt.Println("directions: ", directions)
	scanner.Scan()
	nodes := map[string]Node{}
	for scanner.Scan() {
		// AAA = (BBB, CCC)
		m := reg.FindAllStringSubmatch(scanner.Text(), -1)[0]

		nodes[m[1]] = Node{
			label: m[1],
			left:  m[2],
			right: m[3],
		}
	}

	// part 1
	curr := nodes["AAA"]
	i := 0
	for {
		if directions[i%len(directions)] {
			curr = nodes[curr.left]
		} else {
			curr = nodes[curr.right]
		}
		i++
		sol1 += 1
		if curr.label == "ZZZ" {
			break
		}
	}

	// Part 2
	currs := []Node{}
	for _, n := range nodes {
		if strings.HasSuffix(n.label, "A") {
			currs = append(currs, n)
		}
	}
	fmt.Println("Start", currs)

	ends := make([]int, len(currs))

	i = 0
	for {
		//fmt.Println("Currents", currs)
		for c := 0; c < len(currs); c++ {

			if directions[i%len(directions)] {
				currs[c] = nodes[currs[c].left]
			} else {
				currs[c] = nodes[currs[c].right]
			}

			if strings.HasSuffix(currs[c].label, "Z") {
				// c, i
				ends[c] = i + 1
			}
			fmt.Println(ends)
		}

		i++

		exit := true
		for _, x := range ends {
			if x == 0 {
				exit = false
			}
		}

		if exit {
			break
		}

	}

	sol2 = lcm(ends)

	fmt.Println("1: ", sol1)
	fmt.Println("2: ", sol2)
}

func gcd(a int, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func lcm(args []int) int {
	fmt.Println(args)
	ans := 1
	for _, n := range args {
		ans = ans * n / gcd(ans, n)
	}

	return ans
}

func main() {
	// solve("./2023/day08/cexample3.txt")
	solve("./2023/day08/input.txt")
}
