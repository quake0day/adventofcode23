package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// ReadLines 从指定的文件中读取所有行并返回一个字符串切片
func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
func main() {
	lines, err := ReadLines("day1.txt")
	if err != nil {
		// 错误处理
		fmt.Println("Error reading file:", err)
		return
	}
	total := 0
	for _, line := range lines {
		calibrationValue := getCalibrationValue(line)
		total += calibrationValue
	}

	fmt.Println("Total calibration value:", total)
}
func getCalibrationValue(line string) int {
	wordToDigit := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	// 寻找第一个数字
	firstDigit := findFirstDigit(line, wordToDigit)
	// 寻找最后一个数字
	lastDigit := findLastDigit(line, wordToDigit)

	if firstDigit == -1 || lastDigit == -1 {
		return 0
	}

	return firstDigit*10 + lastDigit
}

// 寻找第一个数字
func findFirstDigit(line string, wordToDigit map[string]string) int {
	for i := 0; i < len(line); i++ {
		for word, digit := range wordToDigit {
			if strings.HasPrefix(line[i:], word) {
				d, _ := strconv.Atoi(digit)
				return d
			}
		}
		if unicode.IsDigit(rune(line[i])) {
			d, _ := strconv.Atoi(string(line[i]))
			return d
		}
	}
	return -1
}

// 寻找最后一个数字
func findLastDigit(line string, wordToDigit map[string]string) int {
	for i := len(line) - 1; i >= 0; i-- {
		for word, digit := range wordToDigit {
			if i >= len(word)-1 && strings.HasSuffix(line[:i+1], word) {
				d, _ := strconv.Atoi(digit)
				return d
			}
		}
		if unicode.IsDigit(rune(line[i])) {
			d, _ := strconv.Atoi(string(line[i]))
			return d
		}
	}
	return -1
}
