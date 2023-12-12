package main

import (
	"fmt"
	"strconv"
)

func main() {
	schematic := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	sum := 0
	for y, line := range schematic {
		for x, ch := range line {
			if isDigit(ch) && isAdjacentToSymbol(x, y, schematic) {
				value, _ := strconv.Atoi(string(ch))
				sum += value
			}
		}
	}
	fmt.Println("The sum of all part numbers is:", sum)
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
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
