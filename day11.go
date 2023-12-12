package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Expand the universe based on the given rules
func expandUniverse(universe []string) []string {
	rows := len(universe)
	if rows == 0 {
		return []string{}
	}
	cols := len(universe[0])

	// Step 1: Identify empty rows and columns
	emptyRows := make([]bool, rows)
	emptyCols := make([]bool, cols)
	for r, row := range universe {
		for c, val := range row {
			if val == '#' {
				emptyRows[r] = true
				emptyCols[c] = true
			}
		}
	}

	// Calculate new size
	newRows := 0
	for _, empty := range emptyRows {
		if !empty {
			newRows += 2
		} else {
			newRows++
		}
	}
	newCols := 0
	for _, empty := range emptyCols {
		if !empty {
			newCols += 2
		} else {
			newCols++
		}
	}

	// Step 2: Create new expanded universe
	newUniverse := make([]string, newRows)
	for i := range newUniverse {
		newUniverse[i] = strings.Repeat(".", newCols)
	}

	// Step 3: Copy data to the new grid
	newRow := 0
	for r, row := range universe {
		newCol := 0
		for c, val := range row {
			if val != '.' {
				s := []rune(newUniverse[newRow])
				s[newCol] = val
				newUniverse[newRow] = string(s)
			}
			if !emptyCols[c] {
				newCol++
			}
			newCol++
		}
		if !emptyRows[r] {
			newUniverse[newRow+1] = newUniverse[newRow]
			newRow++
		}
		newRow++
	}

	return newUniverse
}

type Point struct {
	x, y int
}

// Checks if a point is within the grid
func isValid(universe []string, point Point) bool {
	return point.x >= 0 && point.x < len(universe) && point.y >= 0 && point.y < len(universe[0])
}

// BFS to find the shortest path
func shortestPathLength(universe []string, galaxyMap map[int][2]int, galaxy1, galaxy2 int) int {
	src := Point{galaxyMap[galaxy1][0], galaxyMap[galaxy1][1]}
	dest := Point{galaxyMap[galaxy2][0], galaxyMap[galaxy2][1]}

	// Directions: up, down, left, right
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, len(universe))
	for i := range visited {
		visited[i] = make([]bool, len(universe[0]))
	}

	queue := []Point{src}
	visited[src.x][src.y] = true
	distance := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			point := queue[0]
			queue = queue[1:]

			if point.x == dest.x && point.y == dest.y {
				return distance
			}

			for _, dir := range directions {
				nextPoint := Point{point.x + dir.x, point.y + dir.y}
				if isValid(universe, nextPoint) && !visited[nextPoint.x][nextPoint.y] {
					queue = append(queue, nextPoint)
					visited[nextPoint.x][nextPoint.y] = true
				}
			}
		}
		distance++
	}

	return -1 // Destination not reachable
}

// Calculate the sum of the shortest path lengths between all pairs of galaxies
func sumOfShortestPathLengths(universe []string) int {
	expandedUniverse := expandUniverse(universe)
	galaxyCount := 0
	galaxyMap := make(map[int][2]int)

	// Assign unique numbers to galaxies and build galaxy map
	for i, row := range expandedUniverse {
		for j, char := range row {
			if char == '#' {
				galaxyCount++
				galaxyMap[galaxyCount] = [2]int{i, j}
			}
		}
	}
	fmt.Println(galaxyMap)
	fmt.Println(shortestPathLength(expandedUniverse, galaxyMap, 1, 7))
	fmt.Println(shortestPathLength(expandedUniverse, galaxyMap, 3, 6))
	fmt.Println(shortestPathLength(expandedUniverse, galaxyMap, 8, 9))

	// Calculate the sum of the shortest path lengths
	sum := 0
	for i := 1; i <= galaxyCount; i++ {
		for j := i + 1; j <= galaxyCount; j++ {
			sum += shortestPathLength(expandedUniverse, galaxyMap, i, j)
		}
	}

	return sum
}

// ReadUniverse reads the universe configuration from a file
func ReadUniverse(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var universe []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		universe = append(universe, scanner.Text())
	}

	return universe, scanner.Err()
}

func main() {
	// Read universe from file
	universe, err := ReadUniverse("day11.txt")
	if err != nil {
		log.Fatalf("Failed to read universe from file: %v", err)
	}
	//fmt.Println(expandUniverse(universe))
	sum := sumOfShortestPathLengths(universe)

	fmt.Println("The sum of the shortest path lengths between all pairs of galaxies is:", sum)
}
