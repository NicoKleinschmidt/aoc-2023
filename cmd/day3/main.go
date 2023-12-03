package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Vec2 struct {
	X int
	Y int
}

func main() {
	input, _ := io.ReadAll(os.Stdin)
	inputLines := strings.Split(string(input), "\n")
	numbers := findPartNumbers(inputLines)
	ratios := findGearRatios(inputLines)

	fmt.Printf("Sum of parts: %d\n", sum(numbers))
	fmt.Printf("Sum gear ratios: %d", sum(ratios))
}

func findGearRatios(lines []string) (ratios []int) {
	addedNumbers := make(map[int]int)

	for y, line := range lines {
		for x, r := range line {
			if r != '*' {
				continue
			}

			surroundingDigits := findSurroundingDigits(lines, x, y)
			var gears []int

			for _, index := range surroundingDigits {
				line := lines[index.Y]
				number, start := extractNumber(line, index.X)
				charIndex := index.Y*len(line) + start

				if _, ok := addedNumbers[charIndex]; ok {
					continue
				}

				addedNumbers[charIndex] = number
				gears = append(gears, number)
			}

			if len(gears) == 2 {
				ratios = append(ratios, gears[0]*gears[1])
			}
		}
	}
	return
}

func findPartNumbers(lines []string) (numbers []int) {
	addedNumbers := make(map[int]int)

	for y, line := range lines {
		for x, r := range line {
			if r == '.' || unicode.IsDigit(r) {
				continue
			}

			surroundingDigits := findSurroundingDigits(lines, x, y)
			for _, index := range surroundingDigits {
				line := lines[index.Y]
				number, start := extractNumber(line, index.X)
				charIndex := index.Y*len(line) + start

				if _, ok := addedNumbers[charIndex]; ok {
					continue
				}

				addedNumbers[charIndex] = number
				numbers = append(numbers, number)
			}
		}
	}
	return
}

func findSurroundingDigits(lines []string, x, y int) (indecies []Vec2) {
	for offsetX := -1; offsetX <= 1; offsetX++ {
		for offsetY := -1; offsetY <= 1; offsetY++ {
			newX := x + offsetX
			newY := y + offsetY

			if newY < 0 || newX < 0 || newY >= len(lines) {
				continue
			}

			line := lines[newY]

			if newX >= len(line) {
				continue
			}

			if unicode.IsDigit(rune(line[newX])) {
				indecies = append(indecies, Vec2{X: newX, Y: newY})
			}
		}
	}
	return
}

func extractNumber(line string, index int) (number int, startIndex int) {
	numberStart := 0
	numberEnd := len(line)

	for i := index; i < numberEnd; i++ {
		r := rune(line[i])
		if !unicode.IsDigit(r) {
			numberEnd = i
			break
		}
	}

	for i := index; i >= numberStart; i-- {
		r := rune(line[i])
		if !unicode.IsDigit(r) {
			numberStart = i + 1
			break
		}
	}

	number, _ = strconv.Atoi(line[numberStart:numberEnd])
	return number, numberStart
}

func sum(ints []int) (sum int) {
	for _, num := range ints {
		sum += num
	}
	return
}
