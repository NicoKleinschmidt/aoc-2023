package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sequences [][]int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		sequences = append(sequences, parseSequence(scanner.Text()))
	}

	var sumNext, sumPrev int

	for _, sequence := range sequences {
		sumNext += nextValue(sequence)
		sumPrev += prevValue(sequence)
	}

	fmt.Printf("Sum next: %d\n", sumNext)
	fmt.Printf("Sum previous: %d\n", sumPrev)
}

func parseSequence(line string) []int {
	split := strings.Split(line, " ")
	ints := make([]int, len(split))
	for i, str := range split {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

func difference(sequence []int) (diff []int, allZero bool) {
	diff = make([]int, len(sequence)-1)
	allZero = true

	for i := range diff {
		numA := sequence[i]
		numB := sequence[i+1]

		if numA != 0 || numB != 0 {
			allZero = false
		}

		diff[i] = numB - numA
	}

	return
}

func nextValue(sequence []int) int {
	diff, allZero := difference(sequence)
	if allZero {
		return 0
	}

	return sequence[len(sequence)-1] + nextValue(diff)
}

func prevValue(sequence []int) int {
	diff, allZero := difference(sequence)

	if allZero {
		return 0
	}
	return sequence[0] - prevValue(diff)
}
