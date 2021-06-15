package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strings"
)

const (
	empty = "."
	mine  = "*"
)

func fill(matrix [][]string) {
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = empty
		}
	}
}

func print(matrix [][]string) {
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	var n, m, countLines, line int
	var matrix = make([][]string, 0)

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		if line == 0 {
			fmt.Sscanf(in.Text(), "%d %d %d", &m, &n, &countLines)
			line++
			matrix = make([][]string, m)
			for i := range matrix {
				matrix[i] = make([]string, n)
			}
			fill(matrix)
			continue
		}

		fmt.Sscanf(in.Text(), "%d %d", &m, &n)

		row, col := m-1, n-1
		matrix[row][col] = "*"

		if line == countLines {
			break
		}
		line++
	}
	traverseAndSolveField(matrix)
	print(matrix)
}

func traverseAndSolveField(matrix [][]string) {
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == empty {
				left := point{row, col - 1}
				right := point{row, col + 1}
				up := point{row - 1, col}
				down := point{row + 1, col}
				upperRight := point{row - 1, col + 1}
				upperLeft := point{row - 1, col - 1}
				downRight := point{row + 1, col + 1}
				downLeft := point{row + 1, col - 1}

				d := []point{
					left, right, up, down,
					upperRight, upperLeft,
					downLeft, downRight,
				}
				mines := 0
				for _, p := range d {
					if isValidPoint(len(matrix), len(matrix[0]), p) {
						if matrix[p.row][p.col] == mine {
							mines++
						}
					}
				}
				if mines > 0 {
					matrix[row][col] = fmt.Sprintf("%d", mines)
				}
			}
		}
	}

}

func isValidPoint(m, n int, p point) bool {
	if p.col < 0 || p.col >= n || p.row < 0 || p.row >= m {
		return false
	}

	return true
}

type point struct {
	row int
	col int
}
