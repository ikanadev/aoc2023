package day5

import (
	"strconv"
	"strings"

	"github.com/jqvk/aoc2023/common"
)

type RangeMap struct {
	inputStart  int
	outputStart int
	lenght      int
}
type Mapper struct {
	name string
	maps []RangeMap
}

func (m *Mapper) mapSeeds(seeds []int) []int {
	newSeeds := make([]int, len(seeds))
	for i, seed := range seeds {
		seedAdded := false
		for _, rm := range m.maps {
			if seed >= rm.inputStart && seed <= (rm.inputStart+rm.lenght-1) {
				newSeeds[i] = rm.outputStart + (seed - rm.inputStart)
				seedAdded = true
			}
		}
		if !seedAdded {
			newSeeds[i] = seed
		}
	}
	return newSeeds
}

func parseMapper(lines []string) Mapper {
	var mapper Mapper
	mapper.name = (strings.Split(lines[0], " "))[0]
	rest := lines[1:]
	for _, line := range rest {
		numbersStr := strings.Fields(line)
		numbers := make([]int, len(numbersStr))
		for i, numberStr := range numbersStr {
			n, err := strconv.Atoi(numberStr)
			common.ErrPanic(err)
			numbers[i] = n
		}
		var rangeMap RangeMap
		rangeMap.outputStart = numbers[0]
		rangeMap.inputStart = numbers[1]
		rangeMap.lenght = numbers[2]
		mapper.maps = append(mapper.maps, rangeMap)
	}
	return mapper
}

func parseSeeds(line string) []int {
	seedsStr := strings.Fields(strings.TrimSpace((strings.Split(line, ":"))[1]))
	seeds := make([]int, len(seedsStr))
	for i, seedStr := range seedsStr {
		n, err := strconv.Atoi(seedStr)
		common.ErrPanic(err)
		seeds[i] = n
	}
	return seeds
}

func parseMappers(lines []string) []Mapper {
	var mappers []Mapper
	var mapperLines []string
	for _, line := range lines {
		if line == "" {
			mapper := parseMapper(mapperLines)
			mappers = append(mappers, mapper)
			mapperLines = []string{}
		} else {
			mapperLines = append(mapperLines, line)
		}
	}
	mapper := parseMapper(mapperLines)
	mappers = append(mappers, mapper)
	return mappers
}

func Part1() int {
	// lines := common.ReadFile("day5/input_small")
	lines := common.ReadFile("day5/input")
	seeds := parseSeeds(lines[0])
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
