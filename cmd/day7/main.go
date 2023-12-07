package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand string

type HandBid struct {
	Cards Hand
	Bid   int
}

func main() {
	var hands []HandBid

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		hands = append(hands, parseHandBid(scanner.Text()))
	}

	sort.Slice(hands, func(i, j int) bool {
		return handLess(hands[i].Cards, hands[j].Cards)
	})

	var totalWinnings int
	for i, handBid := range hands {
		totalWinnings += (i + 1) * handBid.Bid
	}

	fmt.Printf("Total winnings: %d\n", totalWinnings)
}

func parseHandBid(line string) HandBid {
	sep := strings.Index(line, " ")
	bid, _ := strconv.Atoi(strings.TrimSpace(line[sep+1:]))

	return HandBid{
		Cards: Hand(line[:sep]),
		Bid:   bid,
	}
}

func cardCounts(cards Hand) map[rune]int {
	m := make(map[rune]int)
	for _, card := range cards {
		m[card]++
	}
	jokerCount := m['J']
	if jokerCount != 0 {
		delete(m, 'J')
	}
	var maxCard rune
	var maxCount int
	for card, count := range m {
		if count > maxCount {
			maxCard = card
			maxCount = count
		}
	}
	m[maxCard] = maxCount + jokerCount

	return m
}

func nOfAKind(cards Hand, n int) bool {
	counts := cardCounts(cards)
	for _, count := range counts {
		if count >= n {
			return true
		}
	}
	return false
}

func nPairs(cards Hand, n int) bool {
	counts := cardCounts(cards)
	pairs := 0
	for _, count := range counts {
		if count >= 2 {
			pairs++
			if pairs == n {
				return true
			}
		}
	}
	return false
}

// const cardStrength = "23456789TJQKA" // Without joker
const cardStrength = "J23456789TQKA" // With joker

type HandType func(cards Hand) bool

var handTypes = []HandType{
	func(cards Hand) bool {
		return nOfAKind(cards, 5)
	},
	func(cards Hand) bool {
		return nOfAKind(cards, 4)
	},
	func(cards Hand) bool {
		return nOfAKind(cards, 3) && nPairs(cards, 2)
	},
	func(cards Hand) bool {
		return nOfAKind(cards, 3)
	},
	func(cards Hand) bool {
		return nPairs(cards, 2)
	},
	func(cards Hand) bool {
		return nPairs(cards, 1)
	},
	func(cards Hand) bool {
		counts := cardCounts(cards)
		return len(cards) == len(counts)
	},
}

var typeNames = []string{"5 of a kind", "4 of a kind", "full house", "3 of a kind", "2 pair", "1 pair", "high card"}

func handLess(i, j Hand) bool {
	for handIdx, handType := range handTypes {
		iOk := handType(i)
		jOk := handType(j)

		if iOk {
			fmt.Printf("%s = %s\n", i, typeNames[handIdx])
		}

		if iOk && !jOk {
			// i is stronger than j, so i is not less
			return false
		}

		if !iOk && jOk {
			// j is stronger than i, so i is less
			return true
		}

		if iOk && jOk {
			break
		}
	}

	for cardIdx := 0; cardIdx < len(i); cardIdx++ {
		iStrength := strings.Index(cardStrength, string(i[cardIdx]))
		jStrength := strings.Index(cardStrength, string(j[cardIdx]))

		if iStrength == jStrength {
			continue
		}

		return iStrength < jStrength
	}

	return false
}
