package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to calculate the next value in the sequence
func extrapolateNextValue(history []int) int {
	sequences := make([][]int, 0)
	sequences = append(sequences, history)

	for {
		// Calculate differences
		lastSeq := sequences[len(sequences)-1]
		diff := make([]int, len(lastSeq)-1)
		for i := 0; i < len(lastSeq)-1; i++ {
			diff[i] = lastSeq[i+1] - lastSeq[i]
		}
		sequences = append(sequences, diff)

		// Check if all differences are zero
		allZero := true
		for _, d := range diff {
			if d != 0 {
				allZero = false
				break
			}
		}

		if allZero {
			break
		}
	}

	// Extrapolate the next value
	for i := len(sequences) - 2; i >= 0; i-- {
		lastValue := sequences[i][len(sequences[i])-1]
		diffValue := sequences[i+1][len(sequences[i+1])-1]
		sequences[i] = append(sequences[i], lastValue+diffValue)
	}

	// Return the next value of the original sequence
	return sequences[0][len(sequences[0])-1]
}

// Function to calculate the sum of extrapolated values
func calculateSumOfExtrapolatedValues(reports [][]int) int {
	sum := 0
	for _, report := range reports {
		sum += extrapolateNextValue(report)
	}
	return sum
}

func readInput(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var report []int
		for _, field := range fields {
			value, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			report = append(report, value)
		}
		reports = append(reports, report)
	}

	return reports, scanner.Err()
}

func main() {
	reports, err := readInput("day9.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sumOfExtrapolatedValues := calculateSumOfExtrapolatedValues(reports)
	fmt.Printf("Sum of extrapolated values: %d", sumOfExtrapolatedValues)
}
