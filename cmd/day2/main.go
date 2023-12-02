// This program reads a list of games from STDIN.
// It then determines which games are possible using a built in list of maximum values.
// It also calculates the minimum numbers of cubes required for each game to be possible.
//
// The sum of possible game ids is output to STDOUT line 1
// The result of the minimum requires cubes in output TO STDOUT line 2 (Sum of the cubes of each set multiplied together)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeSet map[string]int

type GameRecord struct {
	Id   int
	Sets []CubeSet
}

func main() {
	maxCubes := CubeSet{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var games []GameRecord

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		games = append(games, parseGameRecord(scanner.Text()))
	}

	possible := filterPossibleGames(games, maxCubes)
	possibleGamesSum := sumGamesIds(possible)
	minCubesPowerSum := sumMinimumSetPower(games)

	fmt.Println(possibleGamesSum)
	fmt.Print(minCubesPowerSum)
}

func parseGameRecord(line string) GameRecord {
	gameIdEnd := strings.Index(line, ":")
	gameIdStr := strings.TrimPrefix(line[:gameIdEnd], "Game ")
	gameId, _ := strconv.Atoi(gameIdStr)
	sets := strings.Split(line[gameIdEnd+1:], ";")

	game := GameRecord{
		Id:   gameId,
		Sets: make([]CubeSet, len(sets)),
	}

	for i, set := range sets {
		game.Sets[i] = parseSet(set)
	}

	return game
}

func parseSet(setStr string) (set CubeSet) {
	cubes := strings.Split(setStr, ",")
	set = make(CubeSet, len(cubes))

	for _, cubeStr := range cubes {
		split := strings.SplitN(strings.TrimSpace(cubeStr), " ", 2)
		countStr := split[0]
		colorStr := split[1]
		count, _ := strconv.Atoi(countStr)

		set[colorStr] += count
	}
	return
}

func filterPossibleGames(games []GameRecord, maxCubes CubeSet) (possible []GameRecord) {
	for _, game := range games {
		if gamePossible(game, maxCubes) {
			possible = append(possible, game)
		}
	}
	return
}

func sumMinimumSetPower(games []GameRecord) (sum int) {
	for _, game := range games {
		minimumCubes := minimumRequired(game)
		sum += powerOfSet(minimumCubes)
	}
	return
}

func gamePossible(game GameRecord, maxCubes CubeSet) bool {
	for _, set := range game.Sets {
		for color, count := range set {
			if count > maxCubes[color] {
				return false
			}
		}
	}

	return true
}

func minimumRequired(game GameRecord) CubeSet {
	minCubes := make(CubeSet)

	for _, set := range game.Sets {
		for color, count := range set {
			if count > minCubes[color] {
				minCubes[color] = count
			}
		}
	}

	return minCubes
}

func sumGamesIds(games []GameRecord) (sum int) {
	for _, game := range games {
		sum += game.Id
	}
	return
}

func powerOfSet(set CubeSet) (power int) {
	power = 1
	for _, count := range set {
		power *= count
	}
	return
}
