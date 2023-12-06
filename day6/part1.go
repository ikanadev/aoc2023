package day6

import (
	"strconv"
	"strings"

	"github.com/jqvk/aoc2023/common"
)

type Race struct {
	milliSecs int
	record    int
}

// Returns the first Milliseconds of pressing the button required to break the record
func (r *Race) getFirstRecord() int {
	for i := 1; i < r.milliSecs/2; i++ {
		if i*(r.milliSecs-i) > r.record {
			return i
		}
	}
	panic("there should be always a record")
}

func (r *Race) getWaysToBreakRecord() int {
	firstRecord := r.getFirstRecord()
	ways := 2 * ((r.milliSecs+1)/2 - firstRecord)
	if r.milliSecs%2 == 0 {
		ways++
	}
	return ways
}

func parseInput(lines []string) []Race {
	var races []Race
	timesLine := strings.Fields(lines[0])
	recordLine := strings.Fields(lines[1])
	for i := 1; i < len(timesLine); i++ {
		time, err := strconv.Atoi(timesLine[i])
		common.ErrPanic(err)
		record, err := strconv.Atoi(recordLine[i])
		common.ErrPanic(err)
		races = append(races, Race{milliSecs: time, record: record})
	}
	return races
}

func Part1() int {
	// lines := common.ReadFile("day6/input_small")
	lines := common.ReadFile("day6/input")
	races := parseInput(lines)
  result := 1
	for _, race := range races {
		ways := race.getWaysToBreakRecord()
    result *= ways
	}
  return result
}
