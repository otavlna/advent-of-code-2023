package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	HIGH_CARD       = 1
	ONE_PAIR        = 2
	TWO_PAIRS       = 3
	THREE_OF_A_KIND = 4
	FULL_HOUSE      = 5
	FOUR_OF_A_KIND  = 6
	FIVE_OF_A_KIND  = 7
)

type HandBid struct {
	hand string
	bid  int
}

func main() {
	file, err := os.Open("../input.txt")
	checkErr(err)

	handsWithBids := make([]HandBid, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")

		if len(values) != 2 {
			continue
		}

		hand := values[0]

		bid, err := strconv.Atoi(values[1])
		checkErr(err)

		handsWithBids = append(handsWithBids, HandBid{hand, bid})
	}

	// fmt.Println(handsWithBids)
	slices.SortFunc(handsWithBids, compareHands)
	// fmt.Println(handsWithBids)

	sum := 0
	for i, hb := range handsWithBids {
		sum += (i + 1) * hb.bid
	}

	fmt.Println(sum)
}

func compareHands(hb1 HandBid, hb2 HandBid) int {
	type1, type2 := getHandType(hb1.hand), getHandType(hb2.hand)
	if type1 != type2 {
		return type1 - type2
	}
	for i := range hb1.hand {
		if hb1.hand[i] != hb2.hand[i] {
			return getCardValue(rune(hb1.hand[i])) - getCardValue(rune(hb2.hand[i]))
		}
	}
	return 0
}

func getHandType(h string) int {
	cardCounts := map[rune]int{}
	for _, r := range h {
		cardCounts[r]++
	}
	switch {
	case isFiveOfAKind(cardCounts):
		return FIVE_OF_A_KIND
	case isFourOfAKind(cardCounts):
		return FOUR_OF_A_KIND
	case isFullHouse(cardCounts):
		return FULL_HOUSE
	case isThreeOfAKind(cardCounts):
		return THREE_OF_A_KIND
	case isTwoPairs(cardCounts):
		return TWO_PAIRS
	case isOnePair(cardCounts):
		return ONE_PAIR
	default:
		return HIGH_CARD
	}
}

func isFiveOfAKind(cardCounts map[rune]int) bool {
	for _, count := range cardCounts {
		if count == 5 {
			return true
		}
	}
	return false

}

func isFourOfAKind(cardCounts map[rune]int) bool {
	for _, count := range cardCounts {
		if count == 4 {
			return true
		}
	}
	return false

}

func isFullHouse(cardCounts map[rune]int) bool {
	var hasThree, hasTwo bool = false, false
	for _, count := range cardCounts {
		if count == 3 {
			hasThree = true
		}
		if count == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}

func isThreeOfAKind(cardCounts map[rune]int) bool {
	for _, count := range cardCounts {
		if count == 3 {
			return true
		}
	}
	return false
}

func isTwoPairs(cardCounts map[rune]int) bool {
	pairCount := 0
	for _, count := range cardCounts {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount == 2
}

func isOnePair(cardCounts map[rune]int) bool {
	for _, count := range cardCounts {
		if count == 2 {
			return true
		}
	}
	return false
}

func getCardValue(card rune) int {
	cardValues := map[rune]int{
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'J': 10,
		'Q': 11,
		'K': 12,
		'A': 13,
	}
	return cardValues[card]
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
