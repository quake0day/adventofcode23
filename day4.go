package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	WinningNumbers []int
	YourNumbers    []int
}

func readAndParseFile(filename string) ([]Card, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cards []Card
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		card := parseCard(line)
		cards = append(cards, card)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return cards, nil
}

func parseCard(input string) Card {
	parts := strings.Split(input, "|")
	if len(parts) != 2 {
		log.Fatalf("Invalid card format: %s", input)
	}

	// Skipping the "Card X:" part
	winningNumbers := parseNumbers(parts[0][strings.Index(parts[0], ":")+1:])
	yourNumbers := parseNumbers(parts[1])

	return Card{
		WinningNumbers: winningNumbers,
		YourNumbers:    yourNumbers,
	}
}

func parseNumbers(input string) []int {
	fields := strings.Fields(strings.TrimSpace(input))
	var numbers []int
	for _, f := range fields {
		num, err := strconv.Atoi(f)
		if err != nil {
			log.Fatalf("Invalid number format: %s", f)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func calculatePoints(card Card) int {
	points := 0
	matchFound := false
	for _, num := range card.YourNumbers {
		if contains(card.WinningNumbers, num) {
			if matchFound {
				points *= 2
			} else {
				points = 1
				matchFound = true
			}
		}
	}
	return points
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
func totalPoints(cards []Card) int {
	total := 0
	for _, card := range cards {
		total += calculatePoints(card)
	}
	return total
}

func main() {
	cards, err := readAndParseFile("day4.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	fmt.Println("Total points:", totalPoints(cards))
}
