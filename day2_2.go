package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2_data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var totalPower int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		red, green, blue := findMinimumCubes(line)
		power := red * green * blue
		totalPower += power
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Total power of all games:", totalPower)
}

func findMinimumCubes(line string) (int, int, int) {
	parts := strings.Split(line, ": ")
	cubeSets := strings.Split(parts[1], "; ")
	maxCubes := map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, set := range cubeSets {
		updateMaxCubes(maxCubes, set)
	}
	return maxCubes["red"], maxCubes["green"], maxCubes["blue"]
}

func updateMaxCubes(maxCubes map[string]int, set string) {
	cubes := strings.Split(set, ", ")
	for _, cube := range cubes {
		parts := strings.Split(cube, " ")
		count, _ := strconv.Atoi(parts[0])
		color := parts[1]
		if count > maxCubes[color] {
			maxCubes[color] = count
		}
	}
}
