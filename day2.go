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

	var sumOfIDs int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isGamePossible(line) {
			gameID := getGameID(line)
			sumOfIDs += gameID
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Sum of IDs of possible games:", sumOfIDs)
}

func isGamePossible(line string) bool {
	parts := strings.Split(line, ": ")
	cubeSets := strings.Split(parts[1], "; ")
	for _, set := range cubeSets {
		if !isSetPossible(set) {
			return false
		}
	}
	return true
}

func getGameID(line string) int {
	parts := strings.Split(line, ":")
	idStr := strings.TrimPrefix(parts[0], "Game ")
	id, _ := strconv.Atoi(idStr)
	return id
}

func isSetPossible(set string) bool {
	maxCubes := map[string]int{"red": 12, "green": 13, "blue": 14}
	cubes := strings.Split(set, ", ")
	for _, cube := range cubes {
		parts := strings.Split(cube, " ")
		count, _ := strconv.Atoi(parts[0])
		color := parts[1]
		if count > maxCubes[color] {
			return false
		}
	}
	return true
}
