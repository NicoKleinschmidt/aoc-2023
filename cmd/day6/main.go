package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type RaceInfo struct {
	MaxTime     int
	MinDistance int
}

type TimeRange struct {
	Start int
	End   int
}

func main() {
	input, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(string(input), "\n")
	racesWithSpace := parseRaces(lines, false)
	racesWithoutSpace := parseRaces(lines, true)

	fmt.Printf("Ways to win (multi race): %d\n", winningMovesProduct(racesWithSpace))
	fmt.Printf("Ways to win (one long race): %d\n", winningMovesProduct(racesWithoutSpace))
}

func winningMovesProduct(races []RaceInfo) int {
	product := 1
	for _, race := range races {
		winningTimes := winningButtonPressTimes(race)
		winningMoveCount := winningTimes.End - winningTimes.Start
		product *= winningMoveCount
	}
	return product
}

func parseRaces(lines []string, ignoreSpace bool) []RaceInfo {
	if len(lines) != 2 {
		return nil
	}

	times := extractNumbers(strings.TrimPrefix(lines[0], "Time:"), ignoreSpace)
	distances := extractNumbers(strings.TrimPrefix(lines[1], "Distance:"), ignoreSpace)
	races := make([]RaceInfo, min(len(times), len(distances)))

	for i := range races {
		races[i].MaxTime = times[i]
		races[i].MinDistance = distances[i]
	}

	return races
}

func extractNumbers(line string, ignoreSpace bool) (numbers []int) {
	if ignoreSpace {
		line = strings.ReplaceAll(line, " ", "")
	}
	line = strings.TrimSpace(line)

	var currentNum string
	addNumber := func() {
		num, _ := strconv.Atoi(currentNum)
		numbers = append(numbers, num)
		currentNum = ""
	}

	for _, ch := range line {
		if unicode.IsDigit(ch) {
			currentNum += string(ch)
			continue
		}

		if currentNum != "" {
			addNumber()
		}
	}

	if currentNum != "" {
		addNumber()
	}

	return
}

func winningButtonPressTimes(race RaceInfo) TimeRange {
	r := TimeRange{
		Start: -1,
		End:   -1,
	}

	for buttonTime := 0; buttonTime < race.MaxTime; buttonTime++ {
		distance := calcDistance(buttonTime, race.MaxTime)

		if r.Start == -1 && distance > race.MinDistance {
			r.Start = buttonTime
			r.End = buttonTime + 1
		}

		if r.Start != -1 && distance <= race.MinDistance {
			break
		}

		r.End = buttonTime + 1
	}

	return r
}

func calcDistance(buttonTime, maxTime int) int {
	if buttonTime > maxTime {
		return 0
	}
	return buttonTime * (maxTime - buttonTime)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
