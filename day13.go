package main

import (
	"fmt"
	"strings"
)

func main() {
	input := `#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

	// Split the input into lines and convert it into a 2D array
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			if char == '#' {
				grid[i][j] = 1
			}
		}
	}

	// Check for the center of horizontal and vertical reflection
	if center, found := findHorizontalReflectionCenter(grid); found {
		fmt.Printf("Horizontal reflection center found at line: %d\n", center+1) // Lines are 1-indexed
	} else {
		fmt.Println("No horizontal reflection center found.")
	}

	if center, found := findVerticalReflectionCenter(grid); found {
		fmt.Printf("Vertical reflection center found at column: %d\n", center+1) // Columns are 1-indexed
	} else {
		fmt.Println("No vertical reflection center found.")
	}
}

func findHorizontalReflectionCenter(grid [][]int) (int, bool) {
	numRows := len(grid)
	mid := numRows / 2
	for offset := 0; offset <= mid; offset++ {
		if mid-offset >= 0 && mid+offset < numRows && equals(grid[mid-offset], grid[mid+offset]) {
			return mid, true
		}
	}
	return -1, false
}

func findVerticalReflectionCenter(grid [][]int) (int, bool) {
	numCols := len(grid[0])
	mid := numCols / 2
	for offset := 0; offset <= mid; offset++ {
		if mid-offset >= 0 && mid+offset < numCols && equals(getColumn(grid, mid-offset), getColumn(grid, mid+offset)) {
			return mid, true
		}
	}
	return -1, false
}

func getColumn(grid [][]int, col int) []int {
	column := make([]int, len(grid))
	for i := range grid {
		column[i] = grid[i][col]
	}
	return column
}

func equals(row1, row2 []int) bool {
	if len(row1) != len(row2) {
		return false
	}
	for i := range row1 {
		if row1[i] != row2[i] {
			return false
		}
	}
	return true
}
