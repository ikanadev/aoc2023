package day6

import (
	"strconv"
	"strings"

	"github.com/jqvk/aoc2023/common"
)

func parseRace(lines []string) Race {
	var race Race
  timeStr := (strings.Split(strings.ReplaceAll(lines[0], " ", ""), ":"))[1]
	recordStr := (strings.Split(strings.ReplaceAll(lines[1], " ", ""), ":"))[1]
  time, err := strconv.Atoi(timeStr)
  common.ErrPanic(err)
  record, err := strconv.Atoi(recordStr)
  common.ErrPanic(err)
  race.milliSecs = time
  race.record = record
	return race
}

func Part2() int {
	lines := common.ReadFile("day6/input")
  race := parseRace(lines)
  ways := race.getWaysToBreakRecord()
  return ways
}
