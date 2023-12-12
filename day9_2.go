package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to calculate the previous value in the sequence by extrapolating backwards
func extrapolatePreviousValue(history []int) int {
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

	// Extrapolate the previous value
	for i := len(sequences) - 2; i >= 0; i-- {
		firstValue := sequences[i][0]
		diffValue := sequences[i+1][0]
		newFirstValue := firstValue - diffValue
		sequences[i] = append([]int{newFirstValue}, sequences[i]...)
	}

	// Return the new first value of the original sequence
	return sequences[0][0]
}

// Function to read the file and process the histories
func processFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		history := make([]int, len(numbers))
		for i, n := range numbers {
			history[i], err = strconv.Atoi(n)
			if err != nil {
				return 0, err
			}
		}
		sum += extrapolatePreviousValue(history)
	}
	return sum, scanner.Err()
}

func main() {
	sum, err := processFile("day9.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Sum of extrapolated previous values: %d\n", sum)
}
