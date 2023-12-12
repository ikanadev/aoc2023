package day05

import (
	"strconv"
	"strings"

	"github.com/jqvk/aoc2023/common"
)

// seeds: 79 14 55 13
func parseRangeSeeds(line string) []int {
	seedsStr := strings.Fields(strings.TrimSpace((strings.Split(line, ":"))[1]))
	seedsRanges := make([]int, len(seedsStr))
	for i, seedStr := range seedsStr {
		n, err := strconv.Atoi(seedStr)
		common.ErrPanic(err)
		seedsRanges[i] = n
	}
	var seeds []int
	for i := 0; i < len(seedsRanges); i = i + 2 {
		start := seedsRanges[i]
		length := seedsRanges[i+1]
		for j := start; j < (start + length); j++ {
      seeds = append(seeds, j)
		}
	}
	return seeds
}

func Part2() int {
	lines := common.ReadFile("day05/input_small")
	// lines := common.ReadFile("day05/input")
	seeds := parseRangeSeeds(lines[0])
	mappers := parseMappers(lines[2:])
	for _, mapper := range mappers {
		seeds = mapper.mapSeeds(seeds)
	}
	lowestLocation := seeds[0]
	for _, seed := range seeds {
		if seed < lowestLocation {
			lowestLocation = seed
		}
	}
	return lowestLocation
}
