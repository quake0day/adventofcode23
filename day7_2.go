package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Hand represents a poker hand and its bid
type Hand struct {
	Cards string
	Bid   int
	Rank  int
}

// ByStrength implements sort.Interface based on the strength of the hands
type ByStrength []Hand

func (h ByStrength) Len() int {
	return len(h)
}

func (h ByStrength) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h ByStrength) Less(i, j int) bool {
	typeI, strengthI := evaluateHand(h[i])
	typeJ, strengthJ := evaluateHand(h[j])

	if typeI != typeJ {
		return typeI < typeJ // Assuming higher type numbers are stronger
	}

	// Compare strengths if types are the same
	for k := range strengthI {
		if strengthI[k] != strengthJ[k] {
			return strengthI[k] > strengthJ[k]
		}
	}

	return false
}

func parseInput(input []string) []Hand {
	var hands []Hand
	for _, line := range input {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input format")
			continue
		}
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Invalid bid format: %s\n", parts[1])
			continue
		}
		hands = append(hands, Hand{Cards: parts[0], Bid: bid})
	}
	return hands
}

// cardStrength maps a card to its strength
var cardStrength = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10,
	'9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
}

func evaluateHand(hand Hand) (handType int, strength []int) {
	counts := make(map[rune]int)
	jokerCount := 0

	for _, card := range hand.Cards {
		if card == 'J' {
			jokerCount++
			continue
		}
		counts[card]++
		strength = append(strength, cardStrength[card])
	}

	var pairCount, tripleCount, quadCount int
	for _, count := range counts {
		switch count {
		case 2:
			pairCount++
		case 3:
			tripleCount++
		case 4:
			quadCount++
		}
	}

	// 特殊情况处理
	if jokerCount > 0 {
		// 有 J 存在的情况下，处理特殊情况
		switch {
		case quadCount > 0:
			handType = 1 // "Five of a kind"
		case tripleCount > 0:
			handType = 2 // "Four of a kind"
		case pairCount > 0:
			if jokerCount >= 2 {
				handType = 1 // "Five of a kind"
			} else {
				handType = 3 // "One pair"
			}
		default:
			handType = 4 // "High card"
		}
	} else {
		// 无 J 的情况下，根据常规规则确定手牌类型
		switch {
		case quadCount > 0 || (tripleCount > 0 && pairCount > 0):
			handType = 1 // "Five of a kind" 或 "Four of a kind"
		case tripleCount > 0 || (pairCount > 0 && jokerCount > 0):
			handType = 2 // "Three of a kind"
		case pairCount > 0 || jokerCount > 0:
			handType = 3 // "One pair"
		default:
			handType = 4 // "High card"
		}
	}

	// 为了打破平局，J 总是被视为最弱的牌
	for i := 0; i < jokerCount; i++ {
		strength = append(strength, 1) // J 代表的弱牌
	}

	sort.Sort(sort.Reverse(sort.IntSlice(strength)))
	return handType, strength
}

//func evaluateHand(hand Hand) (handType int, strength []int) {
//	counts := make(map[rune]int)
//	jokerCount := 0
//
//	for _, card := range hand.Cards {
//		if card == 'J' {
//			jokerCount++
//			continue
//		}
//		counts[card]++
//		strength = append(strength, cardStrength[card])
//	}
//
//	var pairCount, tripleCount, quadCount int
//	for _, count := range counts {
//		switch count {
//		case 2:
//			pairCount++
//		case 3:
//			tripleCount++
//		case 4:
//			quadCount++
//		}
//	}
//
//	switch {
//	case quadCount > 0 || (tripleCount > 0 && pairCount > 0) || jokerCount >= 2:
//		handType = 1 // "Five of a kind" 或 "Four of a kind" 或 "Full house"
//	case tripleCount > 0 || (pairCount > 0 && jokerCount > 0):
//		handType = 2 // "Three of a kind" 或 "Two pair"
//	case pairCount > 0 || jokerCount > 0:
//		handType = 3 // "One pair"
//	default:
//		handType = 4 // "High card"
//	}
//
//	// 为了打破平局，J 总是被视为最弱的牌
//	for i := 0; i < jokerCount; i++ {
//		strength = append(strength, 1) // J 代表的弱牌
//	}
//
//	sort.Sort(sort.Reverse(sort.IntSlice(strength)))
//	return handType, strength
//}

func main() {
	file, err := os.Open("day7.txt")
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) != 2 {
			continue // skip invalid lines
		}
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			continue // skip invalid lines
		}
		hands = append(hands, Hand{Cards: parts[0], Bid: bid})
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	sort.Sort(ByStrength(hands))

	totalWinnings := 0
	for i, hand := range hands {
		rank := len(hands) - i // 从末尾开始计数，最强的手牌获得最高等级
		fmt.Println(hand.Bid)
		fmt.Println(rank)
		totalWinnings += hand.Bid * rank
	}
	fmt.Printf("Total Winnings: %d\n", totalWinnings)
}
