package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	schematic, err := readSchematicFromFile("day3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	gearSum := 0
	for y, line := range schematic {
		for x, ch := range line {
			if ch == '*' {
				partNumbers := findAdjacentPartNumbers(x, y, schematic)
				fmt.Println(partNumbers)
				if len(partNumbers) == 2 {
					gearRatio := partNumbers[0] * partNumbers[1]
					gearSum += gearRatio
				}
			}
		}
	}
	fmt.Println("The sum of all gear ratios is:", gearSum)
}

func findAdjacentPartNumbers(x, y int, schematic []string) []int {
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	seen := make(map[string]struct{}) // 用来记录已经处理过的数字序列
	var partNumbers []int

	for i := 0; i < 8; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx >= 0 && ny >= 0 && ny < len(schematic) && nx < len(schematic[ny]) {
			start, end := findNumberSequence(nx, ny, schematic)
			if start != -1 {
				seqKey := fmt.Sprintf("%d-%d-%d", start, end, ny) // 创建一个唯一的键来标识这个数字序列
				if _, exists := seen[seqKey]; !exists {
					seen[seqKey] = struct{}{}
					value, _ := strconv.Atoi(schematic[ny][start:end])
					partNumbers = append(partNumbers, value)
				}
			}
		}
	}
	return partNumbers
}

func findNumberSequence(x, y int, schematic []string) (int, int) {
	line := schematic[y]
	if !isDigit(rune(line[x])) {
		return -1, -1
	}
	start, end := x, x
	for start > 0 && isDigit(rune(line[start-1])) {
		start--
	}
	for end < len(line)-1 && isDigit(rune(line[end+1])) {
		end++
	}
	return start, end + 1
}
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
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

// Other helper functions (isDigit, readSchematicFromFile) remain unchanged
