package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to read the grid from the file
func readGrid(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	return grid, scanner.Err()
}

// Function to find the starting position 'S' in the grid
func findStart(grid [][]rune) (int, int) {
	for y, row := range grid {
		for x, tile := range row {
			if tile == 'S' {
				return x, y
			}
		}
	}
	return -1, -1 // Starting position not found
}

type Point struct {
	x, y int
}

func isValid(grid [][]rune, x, y int) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) && grid[y][x] != '.'
}

func canMove(from, to rune, dir int) bool {
	switch from {
	case '|':
		return (dir == 0 || dir == 2) && (to == '|' || to == 'S' || to == '7' || to == 'F')
	case '-':
		return (dir == 1 || dir == 3) && (to == '-' || to == 'S' || to == 'L' || to == 'J')
	case 'L':
		return ((dir == 0 && (to == '|' || to == 'S' || to == '7' || to == 'F')) || (dir == 3 && (to == '-' || to == 'S' || to == 'L' || to == 'J')))
	case 'J':
		return ((dir == 0 && (to == '|' || to == 'S' || to == '7' || to == 'F')) || (dir == 1 && (to == '-' || to == 'S' || to == 'L' || to == 'J')))
	case '7':
		return ((dir == 2 && (to == '|' || to == 'S' || to == '7' || to == 'F')) || (dir == 1 && (to == '-' || to == 'S' || to == 'L' || to == 'J')))
	case 'F':
		return ((dir == 2 && (to == '|' || to == 'S' || to == '7' || to == 'F')) || (dir == 3 && (to == '-' || to == 'S' || to == 'L' || to == 'J')))
	case 'S':
		return to == '|' || to == '-' || to == 'L' || to == 'J' || to == '7' || to == 'F'
	}
	return false
}

func followLoop(grid [][]rune, startX, startY int) [][]int {
	distances := make([][]int, len(grid))
	for i := range distances {
		distances[i] = make([]int, len(grid[i]))
		for j := range distances[i] {
			distances[i][j] = -1
		}
	}

	// 方向：N, E, S, W
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{-1, 0, 1, 0}

	queue := []Point{{startX, startY}}
	distances[startY][startX] = 0
	visited := make(map[Point]bool)
	visited[Point{startX, startY}] = true

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		currentPipe := grid[point.y][point.x]

		for dir := 0; dir < 4; dir++ {
			nx, ny := point.x+dx[dir], point.y+dy[dir]

			if !isValid(grid, nx, ny) {
				continue
			}

			nextPipe := grid[ny][nx]
			nextPoint := Point{nx, ny}
			if canMove(currentPipe, nextPipe, dir) && !visited[nextPoint] {
				distances[ny][nx] = distances[point.y][point.x] + 1
				queue = append(queue, nextPoint)
				visited[nextPoint] = true
			}
		}
	}

	return distances
}

func main() {
	grid, err := readGrid("day10.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	startX, startY := findStart(grid)
	if startX == -1 || startY == -1 {
		fmt.Println("Starting position 'S' not found in the grid")
		return
	}

	distances := followLoop(grid, startX, startY)

	// Find the maximum distance in the loop
	maxDistance := 0
	for _, row := range distances {
		for _, dist := range row {
			if dist > maxDistance {
				maxDistance = dist
			}
		}
	}

	fmt.Printf("The farthest point from the start is %d steps away.\n", maxDistance)
}
