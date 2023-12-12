package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// Open the file (replace 'calibration.txt' with your file's name)
	file, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		calibrationValue := getCalibrationValue(line)
		total += calibrationValue
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fmt.Println("Total calibration value:", total)
}

func getCalibrationValue(s string) int {
	firstDigit, lastDigit := -1, -1

	for _, r := range s {
		if unicode.IsDigit(r) {
			digit, _ := strconv.Atoi(string(r))
			if firstDigit == -1 {
				firstDigit = digit
			}
			lastDigit = digit
		}
	}

	if firstDigit != -1 && lastDigit != -1 {
		value, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
		return value
	}

	return 0
}
