package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	//"strings"
)

var (
	re = regexp.MustCompile(`XMAS`)
	rr = regexp.MustCompile(`SAMX`)
)

func countMatch(letters [][]rune) int {
	var matches int
	for r := range letters {
		matches += len(re.FindAllString(string(letters[r]), -1))
		matches += len(rr.FindAllString(string(letters[r]), -1))
	}
	return matches
}

func transpose(matrix [][]rune) [][]rune {
	rows, cols, out := len(matrix), len(matrix[0]), [][]rune{}

	for c := 0; c < cols; c++ {
		out = append(out, make([]rune, rows))

		for r := 0; r < rows; r++ {
			out[c] = append(out[c], matrix[r][c])
		}
	}

	return out
}

// Perhaps, there is a better approach to traverse a matrix in a diagonal order
func diagonalLR(matrix [][]rune) [][]rune {
	rows, cols, out := len(matrix), len(matrix[0]), [][]rune{}

	var diagonal = func(sr int, sc int) []rune {
		d := []rune{}
		for r, c := sr, sc; r < rows && c < cols; r, c = r+1, c+1 {
			d = append(d, matrix[r][c])
		}
		return d
	}

	for sr := 0; sr < rows; sr++ {
		out = append(out, diagonal(sr, 0))
	}

	for sc := 1; sc < cols; sc++ {
		out = append(out, diagonal(0, sc))
	}

	return out
}

func diagonalRL(matrix [][]rune) [][]rune {
	rows, cols, out := len(matrix), len(matrix[0]), [][]rune{}

	var diagonal = func(sr int, sc int) []rune {
		d := []rune{}
		for r, c := sr, sc; r >= 0 && c < cols; r, c = r-1, c+1 {
			d = append(d, matrix[r][c])
		}
		return d
	}

	for sr := rows - 1; sr >= 0; sr-- {
		out = append(out, diagonal(sr, 0))
	}

	for sc := 1; sc < cols; sc++ {
		out = append(out, diagonal(rows-1, sc))
	}

	return out
}

func main() {

	var letters [][]rune

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		letters = append(letters, []rune(scanner.Text()))
	}

	var sum, sum2 int

	// Part-1
	// Using regex and matrix manipulation
	sum += countMatch(letters)
	sum += countMatch(transpose(letters))
	sum += countMatch(diagonalLR(letters))
	sum += countMatch(diagonalRL(letters))
	fmt.Println(sum)

	// Part-2
	for r := 1; r < len(letters)-1; r++ {
		for c := 1; c < len(letters[r])-1; c++ {
			if letters[r][c] == 'A' {
				tl, tr, bl, br := letters[r-1][c-1], letters[r-1][c+1], letters[r+1][c-1], letters[r+1][c+1]
				if ((tl == 'M' && br == 'S') || (tl == 'S' && br == 'M')) && ((tr == 'M' && bl == 'S') || (tr == 'S' && bl == 'M')) {
					sum2++
				}
			}
		}
	}

	fmt.Println(sum2)
}
