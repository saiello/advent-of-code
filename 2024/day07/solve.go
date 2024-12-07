package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operator int

const (
	Add Operator = iota
	Multiply
	Concatenation
)

type OperatorFunc func(a int, b int) int

var Operators map[Operator]OperatorFunc = map[Operator]OperatorFunc{
	Add:           func(a int, b int) int { return a + b },
	Multiply:      func(a int, b int) int { return a * b },
	Concatenation: func(a int, b int) int { return atoi(strconv.Itoa(a) + strconv.Itoa(b)) },
}

func AllOperatorCombinations(size int, operators ...Operator) [][]Operator {

	var cartesian = func(sets [][]Operator) [][]Operator {
		ops := [][]Operator{}
		for _, set := range sets {
			for _, op := range operators {
				ops = append(ops, append([]Operator{op}, set...))
			}
		}
		return ops
	}

	out := [][]Operator{}
	for _, op := range operators {
		out = append(out, []Operator{op})
	}

	for i := 1; i < size; i++ {
		out = cartesian(out)
	}
	return out
}

type Equation struct {
	left  int
	right []int
}

func (e Equation) isValid(operators ...Operator) bool {
	for _, ops := range AllOperatorCombinations(len(e.right)-1, operators...) {
		if e.Evaluate(ops) == e.left {
			return true
		}
	}

	return false
}

func (e Equation) Evaluate(ops []Operator) int {
	out := e.right[0]
	for i, o := range ops {
		out = Operators[o](out, e.right[i+1])
	}
	return out
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func atois(s []string) []int {
	is := []int{}
	for i := 0; i < len(s); i++ {
		is = append(is, atoi(s[i]))
	}
	return is
}

func main() {

	equations := []Equation{}

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		substrings := strings.Split(scanner.Text(), ":")
		equations = append(equations, Equation{
			left:  atoi(substrings[0]),
			right: atois(strings.Fields(substrings[1])),
		})
	}

	var sum, sum2 int
	for _, eq := range equations {
		if eq.isValid(Add, Multiply) {
			sum += eq.left
		}

		if eq.isValid(Add, Multiply, Concatenation) {
			sum2 += eq.left
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}
