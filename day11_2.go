package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Point struct to represent coordinates
type Point struct {
	x, y int
}

// Identify which rows and columns are empty in the universe and map galaxies
func analyzeUniverse(universe []string) (map[int]bool, map[int]bool, map[int]Point, int) {
	emptyRows, emptyCols := identifyEmptyRowsAndCols(universe)
	galaxyMap := make(map[int]Point)
	galaxyCount := 0

	for i, row := range universe {
		for j, char := range row {
			if char == '#' {
				galaxyMap[galaxyCount] = Point{i, j}
				galaxyCount++
			}
		}
	}
	fmt.Println(emptyRows)
	fmt.Println(emptyCols)
	fmt.Println(galaxyMap)

	return emptyRows, emptyCols, galaxyMap, galaxyCount
}

// Checks if a point is within the grid
func isValid(universe []string, point Point) bool {
	return point.x >= 0 && point.x < len(universe) && point.y >= 0 && point.y < len(universe[0])
}

func shortestPathLength(universe []string, galaxyMap map[int][2]int, galaxy1, galaxy2 int, emptyRows, emptyCols map[int]bool, expansionFactor int) int {
	src := Point{galaxyMap[galaxy1][0], galaxyMap[galaxy1][1]}
	dest := Point{galaxyMap[galaxy2][0], galaxyMap[galaxy2][1]}

	// Directions: up, down, left, right
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, len(universe))
	for i := range visited {
		visited[i] = make([]bool, len(universe[0]))
	}

	// Queue to store points and their distances
	type QueueItem struct {
		point    Point
		distance int
	}
	queue := []QueueItem{{src, 0}}
	visited[src.x][src.y] = true

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		point := item.point
		distance := item.distance

		if point.x == dest.x && point.y == dest.y {
			return distance
		}

		for _, dir := range directions {
			nextPoint := Point{point.x + dir.x, point.y + dir.y}
			if isValid(universe, nextPoint) && !visited[nextPoint.x][nextPoint.y] {
				nextDistance := distance + 1
				if (dir.x != 0 && emptyRows[nextPoint.x]) || (dir.y != 0 && emptyCols[nextPoint.y]) {
					nextDistance += expansionFactor - 1
				}
				queue = append(queue, QueueItem{nextPoint, nextDistance})
				visited[nextPoint.x][nextPoint.y] = true
			}
		}
	}

	return -1 // Destination not reachable
}
func sumOfShortestPathLengths(universe []string, expansionFactor int) int {
	emptyRows, emptyCols, galaxyMap, _ := analyzeUniverse(universe)
	sum := 0

	for g1 := 1; g1 <= len(galaxyMap); g1++ {
		for g2 := g1 + 1; g2 <= len(galaxyMap); g2++ {
			sum += shortestPathLength(universe, galaxyMap, g1, g2, emptyRows, emptyCols, expansionFactor)
		}
	}

	return sum
}

// Helper functions to find min and max
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Identify which rows and columns are empty in the universe
func identifyEmptyRowsAndCols(universe []string) (map[int]bool, map[int]bool) {
	rows := len(universe)
	cols := len(universe[0])

	emptyRows := make(map[int]bool)
	emptyCols := make(map[int]bool)
	for i, row := range universe {
		if !strings.Contains(row, "#") {
			emptyRows[i] = true
		}
	}

	for j := 0; j < cols; j++ {
		columnIsEmpty := true
		for i := 0; i < rows; i++ {
			if universe[i][j] == '#' {
				columnIsEmpty = false
				break
			}
		}
		emptyCols[j] = columnIsEmpty
	}

	return emptyRows, emptyCols
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

func calculateNumberOfPairs(galaxyCount int) int {
	return galaxyCount * (galaxyCount - 1) / 2
}

func main() {
	// Read universe from file
	universe, err := ReadUniverse("day11.txt")
	if err != nil {
		log.Fatalf("Failed to read universe from file: %v", err)
	}
	_, _, _, galaxyCount := analyzeUniverse(universe)
	fmt.Println(galaxyCount)
	numberOfPairs := calculateNumberOfPairs(galaxyCount)
	fmt.Println(numberOfPairs)
	expansionFactor := 1 // Adjust as needed
	sum := sumOfShortestPathLengths(universe, expansionFactor)
	fmt.Println("The sum of the shortest path lengths between all pairs of galaxies is:", sum)
}
