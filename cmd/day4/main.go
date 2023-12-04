package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	WinningNumbers map[int]struct{}
	HaveNumbers    []int
	Copies         int
}

func main() {
	var totalScore, totalCopies int
	var cards []Card

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cards = append(cards, parseCard(scanner.Text()))
	}

	for i, card := range cards {
		totalCopies += card.Copies
		ownWinningNumbers := getOwnWinningNumbers(card.WinningNumbers, card.HaveNumbers)
		totalScore += calculateScore(ownWinningNumbers)

		for c := 0; c < card.Copies; c++ {
			for j := range ownWinningNumbers {
				if i+j+1 >= len(cards) {
					break
				}

				card := cards[i+j+1]
				card.Copies++
				cards[i+j+1] = card
			}
		}
	}

	fmt.Printf("Single copy score: %d\n", totalScore)
	fmt.Printf("Number of copies: %d", totalCopies)
}

func parseCard(line string) Card {
	withoutPrefix := line[strings.Index(line, ":")+1:]
	columns := strings.Split(strings.ReplaceAll(withoutPrefix, "  ", " "), "|")
	winningList := strings.Split(strings.TrimSpace(columns[0]), " ")
	haveList := strings.Split(strings.TrimSpace(columns[1]), " ")

	card := Card{
		Copies:         1,
		WinningNumbers: make(map[int]struct{}, len(winningList)),
		HaveNumbers:    make([]int, len(haveList)),
	}

	for _, str := range winningList {
		num, _ := strconv.Atoi(strings.TrimSpace(str))
		card.WinningNumbers[num] = struct{}{}
	}

	for i, str := range haveList {
		num, _ := strconv.Atoi(strings.TrimSpace(str))
		card.HaveNumbers[i] = num
	}

	return card
}

func getOwnWinningNumbers(winning map[int]struct{}, have []int) (result []int) {
	for _, num := range have {
		if _, ok := winning[num]; ok {
			result = append(result, num)
		}
	}
	return
}

func calculateScore(winningNumbers []int) int {
	count := len(winningNumbers)
	if count == 0 {
		return 0
	}

	result := 1
	for i := 0; i < count-1; i++ {
		result *= 2
	}
	return result
}
