// Day 1
//
// This program:
//  1. Reads a string line by line (from STDIN)
//  2. Finds the first and last digits of each line
//  3. Joins the digits in that order
//  4. Returns the sum of the result of all the lines (to STDOUT)
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var digits = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	input, _ := io.ReadAll(os.Stdin)
	fmt.Print(parseCalibrationDoc(string(input)))
}

func parseCalibrationDoc(input string) int {
	return sum(calibrationValues(input, digits))
}

func calibrationValues(input string, digits map[string]int) (values []int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		values = append(values, parseLine(line, digits))
	}
	return
}

func parseLine(input string, digits map[string]int) int {
	var first, last int

forward_outer:
	for i := 0; i < len(input); i++ {
		for text, value := range digits {
			substr := input[i:]
			if strings.HasPrefix(substr, text) {
				first = value
				break forward_outer
			}
		}
	}

reverse_outer:
	for i := len(input) - 1; i >= 0; i-- {
		for text, value := range digits {
			substr := input[i:]
			if strings.HasPrefix(substr, text) {
				last = value
				break reverse_outer
			}
		}
	}

	return first*10 + last
}

func sum(ints []int) (sum int) {
	for _, num := range ints {
		sum += num
	}
	return
}
