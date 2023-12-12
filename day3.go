package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Load data from file
	schematic, err := readSchematicFromFile("day3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	sum := 0
	for y, line := range schematic {
		for x := 0; x < len(line); x++ {
			if isDigit(rune(line[x])) {
				// Find the end of the number sequence
				end := x
				for end < len(line) && isDigit(rune(line[end])) {
					end++
				}

				// Check if the number sequence is adjacent to a symbol
				if isNumberAdjacentToSymbol(x, end-1, y, schematic) {
					value, _ := strconv.Atoi(line[x:end])
					sum += value
					x = end - 1 // Skip the rest of the number sequence
				}
			}
		}
	}
	fmt.Println("The sum of all part numbers is:", sum)
}
func readSchematicFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var schematic []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return schematic, nil
}
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isNumberAdjacentToSymbol(startX, endX, y int, schematic []string) bool {
	for x := startX; x <= endX; x++ {
		if isAdjacentToSymbol(x, y, schematic) {
			return true
		}
	}
	return false
}

func isAdjacentToSymbol(x, y int, schematic []string) bool {
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	for i := 0; i < 8; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx >= 0 && ny >= 0 && ny < len(schematic) && nx < len(schematic[ny]) {
			adjCh := schematic[ny][nx]
			if !isDigit(rune(adjCh)) && adjCh != '.' {
				return true
			}
		}
	}
	return false
}
