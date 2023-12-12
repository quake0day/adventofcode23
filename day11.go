package main

import (
	"bufio"
	"fmt"
	"os"
)

// Point represents a coordinate in the grid.
type Point struct {
	x, y int
}

// readInput reads the input file and returns a 2D slice of galaxies and empty spaces.
func readInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	return grid
}

// expandUniverse doubles the size of rows and columns with no galaxies.
func expandUniverse(grid [][]rune) [][]rune {
	rows, cols := len(grid), len(grid[0])

	// Initialize empty rows and columns as true
	emptyRows, emptyCols := make([]bool, rows), make([]bool, cols)
	for i := 0; i < rows; i++ {
		emptyRows[i] = true
	}
	for j := 0; j < cols; j++ {
		emptyCols[j] = true
	}

	// Mark non-empty rows and columns
	for i, row := range grid {
		for j, cell := range row {
			if cell == '#' {
				emptyRows[i] = false
				emptyCols[j] = false
			}
		}
	}

	// Double the size of empty rows and columns
	newGrid := make([][]rune, 0)
	for i, row := range grid {
		newRow := make([]rune, 0)
		for j, cell := range row {
			newRow = append(newRow, cell)
			if emptyCols[j] {
				newRow = append(newRow, '.')
			}
		}
		newGrid = append(newGrid, newRow)
		if emptyRows[i] {
			duplicateRow := make([]rune, len(newRow))
			copy(duplicateRow, newRow)
			newGrid = append(newGrid, duplicateRow)
		}
	}

	return newGrid
}

// labelGalaxies assigns unique numbers to galaxies and returns a map of galaxy number to its coordinates.
func labelGalaxies(grid [][]rune) map[int]Point {
	galaxyCount := 1
	galaxies := make(map[int]Point)
	for i, row := range grid {
		for j, cell := range row {
			if cell == '#' {
				grid[i][j] = rune(galaxyCount)
				galaxies[galaxyCount] = Point{i, j}
				galaxyCount++
			}
		}
	}
	return galaxies
}

// bfs computes the shortest path between two points using Breadth-First Search.
func bfs(grid [][]rune, start, end Point, galaxies map[int]Point) int {
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make(map[Point]bool)
	queue := []Point{start}

	visited[start] = true // Mark the starting point as visited

	var steps int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			point := queue[0]
			queue = queue[1:]

			if point == end {
				return steps
			}

			for _, dir := range directions {
				next := Point{point.x + dir.x, point.y + dir.y}
				if next.x >= 0 && next.x < len(grid) && next.y >= 0 && next.y < len(grid[0]) &&
					!visited[next] {
					if grid[next.x][next.y] == '.' || (grid[next.x][next.y] >= '1' && grid[next.x][next.y] <= rune('0'+len(galaxies))) {
						queue = append(queue, next)
						visited[next] = true
					}
				}
			}
		}
		steps++
	}
	return -1 // Return -1 if no path is found
}

// calculateTotalPathLength computes the sum of the shortest paths between all pairs of galaxies.
func calculateTotalPathLength(grid [][]rune, galaxies map[int]Point) int {
	totalPathLength := 0
	for i := 1; i < len(galaxies); i++ {
		for j := i + 1; j <= len(galaxies); j++ {
			pathLength := bfs(grid, galaxies[i], galaxies[j], galaxies)
			totalPathLength += pathLength
		}
	}
	return totalPathLength
}

func main() {
	grid := readInput("day11.txt")
	expandedGrid := expandUniverse(grid)
	galaxies := labelGalaxies(expandedGrid)
	totalPathLength := calculateTotalPathLength(expandedGrid, galaxies)
	fmt.Println("Total Path Length:", totalPathLength)
}
