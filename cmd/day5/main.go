// Uses a brute force approach, takes about 3.5 min on an M1 Mac mimi (16GB RAM)
package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type MapRange struct {
	DestintaionStart int
	SourceStart      int
	Length           int
}

type Map struct {
	SourceCategory      string
	DestintaionCategory string
	Ranges              []MapRange
}

type SeedRange struct {
	Start int
	End   int
}

func main() {
	input, _ := io.ReadAll(os.Stdin)
	inputSegments := strings.Split(string(input), "\n\n")
	seeds := parseSeeds(inputSegments[0])
	maps := parseMaps(inputSegments[1:])
	locations := mapSeeds(asRange(seeds), maps)

	fmt.Printf("Closest location: %d\n", min(locations))
}

func parseSeeds(seedLine string) []int {
	split := strings.Split(strings.TrimPrefix(seedLine, "seeds: "), " ")
	seeds := make([]int, len(split))
	for i, v := range split {
		seeds[i], _ = strconv.Atoi(v)
	}

	return seeds
}

func asRange(seedRanges []int) (seeds []SeedRange) {
	for i := 0; i < len(seedRanges); i += 2 {
		seeds = append(seeds, SeedRange{
			Start: seedRanges[i],
			End:   seedRanges[i] + seedRanges[i+1],
		})
	}

	return
}

func parseMaps(mapSegments []string) []*Map {
	maps := make([]*Map, len(mapSegments))
	for i, segment := range mapSegments {
		maps[i] = parseMap(segment)
	}

	return maps
}

func parseMap(segment string) *Map {
	lines := strings.Split(segment, "\n")
	label := strings.TrimSuffix(lines[0], " map:")
	labelSegements := strings.Split(label, "-")

	m := &Map{
		SourceCategory:      labelSegements[0],
		DestintaionCategory: labelSegements[2],
	}

	if len(lines) < 2 {
		return m
	}

	ranges := lines[1:]
	m.Ranges = make([]MapRange, len(ranges))

	for i, rangeStr := range ranges {
		values := strings.Split(rangeStr, " ")
		r := MapRange{}
		r.DestintaionStart, _ = strconv.Atoi(values[0])
		r.SourceStart, _ = strconv.Atoi(values[1])
		r.Length, _ = strconv.Atoi(values[2])
		m.Ranges[i] = r
	}

	return m
}

func mapSeeds(seeds []SeedRange, stages []*Map) (mappedSeeds []int) {
	for i, r := range seeds {
		fmt.Printf("%d/%d\n", i, len(seeds))

		for seed := r.Start; seed < r.End; seed++ {
			mapped := seed
			for _, stage := range stages {
				mapped = mapValue(mapped, stage)
			}

			mappedSeeds = append(mappedSeeds, mapped)
		}

		runtime.GC()
	}

	return
}

func mapValue(value int, m *Map) int {
	for _, r := range m.Ranges {
		if value >= r.SourceStart && value < r.SourceStart+r.Length {
			return r.DestintaionStart + (value - r.SourceStart)
		}
	}
	return value
}

func min(s []int) int {
	if len(s) == 0 {
		return 0
	}

	min := s[0]
	for _, i := range s {
		if i < min {
			min = i
		}
	}

	return min
}
