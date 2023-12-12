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
	for _, card := range hand.Cards {
		counts[card]++
	}

	var pairCount, tripleCount, quadCount, fiveCount int
	for _, count := range counts {
		switch count {
		case 2:
			pairCount++
		case 3:
			tripleCount++
		case 4:
			quadCount++
		case 5:
			fiveCount++
		}
	}

	// 记录手牌中每张牌的顺序和强度
	for _, card := range hand.Cards {
		strength = append(strength, cardStrength[card])
	}

	// 根据手牌类型设置handType
	switch {
	case fiveCount == 1:
		handType = 1 // five of a kind
	case quadCount == 1:
		handType = 2 // Four of a kind
	case tripleCount == 1 && pairCount == 1:
		handType = 3 // Full house
	case tripleCount == 1:
		handType = 4 // Three of a kind
	case pairCount == 2:
		handType = 5 // Two pair
	case pairCount == 1:
		handType = 6 // One pair
	default:
		handType = 7 // High card
	}

	return handType, strength
}

func evaluateHand2(hand Hand) (handType int, strength []int) {
	counts := make(map[rune]int)
	for _, card := range hand.Cards {
		counts[card]++
	}

	// Convert counts to slices for easier analysis
	var pairCount, tripleCount, quadCount int
	var pairs, triples, quads, singles []int
	for card, count := range counts {
		switch count {
		case 2:
			pairCount++
			pairs = append(pairs, cardStrength[card])
		case 3:
			tripleCount++
			triples = append(triples, cardStrength[card])
		case 4:
			quadCount++
			quads = append(quads, cardStrength[card])
		case 1:
			singles = append(singles, cardStrength[card])
		}
	}

	sort.Ints(pairs)
	sort.Ints(triples)
	sort.Ints(quads)
	sort.Ints(singles)

	// Determine the hand type
	switch {
	case quadCount == 1:
		handType = 2 // Four of a kind
		strength = append(strength, quads...)
		strength = append(strength, singles...)
	case tripleCount == 1 && pairCount == 1:
		handType = 3 // Full house
		strength = append(strength, triples...)
		strength = append(strength, pairs...)
	case tripleCount == 1:
		handType = 4 // Three of a kind
		strength = append(strength, triples...)
		// 排序剩余的单张牌，以确保强度比较是正确的
		sort.Sort(sort.Reverse(sort.IntSlice(singles)))
		strength = append(strength, singles...)
	case pairCount == 2:
		handType = 5 // Two pair
		strength = append(strength, pairs...)
		strength = append(strength, singles...)
	case pairCount == 1:
		handType = 6 // One pair
		strength = append(strength, pairs...)
		strength = append(strength, singles...)
	default:
		handType = 7 // High card
		strength = append(strength, singles...)
	}
	//switch {
	//case quadCount == 1:
	//	handType = 2 // Four of a kind
	//	strength = append(strength, quads...)
	//	strength = append(strength, singles...)
	//case tripleCount == 1 && pairCount == 1:
	//	handType = 3 // Full house
	//	strength = append(strength, triples...)
	//	strength = append(strength, pairs...)
	//case tripleCount == 1:
	//	handType = 4 // Three of a kind
	//	strength = append(strength, triples...)
	//	// 排序剩余的单张牌，以确保强度比较是正确的
	//	sort.Sort(sort.Reverse(sort.IntSlice(singles)))
	//	strength = append(strength, singles...)
	//case pairCount == 2:
	//	handType = 5 // Two pair
	//	strength = append(strength, pairs...)
	//	strength = append(strength, singles...)
	//case pairCount == 1:
	//	handType = 6 // One pair
	//	strength = append(strength, pairs...)
	//	strength = append(strength, singles...)
	//default:
	//	handType = 7 // High card
	//	strength = append(strength, singles...)
	//}

	//// Reverse the slices so that the highest values come first
	//for i, j := 0, len(strength)-1; i < j; i, j = i+1, j-1 {
	//	strength[i], strength[j] = strength[j], strength[i]
	//}
	fmt.Println(hand, handType, strength)
	return handType, strength
}

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
