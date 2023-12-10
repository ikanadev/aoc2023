package day9

import (
	"strconv"
	"strings"

	"github.com/jqvk/aoc2023/common"
)

type Serie []int

func (s *Serie) isAllZero() bool {
	for _, n := range *s {
		if n != 0 {
			return false
		}
	}
	return true
}
func (s *Serie) calculateNextValue() int {
	matrix := make([]Serie, 0)
	matrix = append(matrix, *s)
	for !matrix[len(matrix)-1].isAllZero() {
		newRow := make(Serie, 0)
		row := matrix[len(matrix)-1]
		for j := 0; j < (len(row) - 1); j++ {
			newRow = append(newRow, row[j+1]-row[j])
		}
		matrix = append(matrix, newRow)
	}
	nextValue := 0
	for i := len(matrix) - 2; i >= 0; i-- {
		nextValue += matrix[i][len(matrix[i])-1]
	}
	return nextValue
}

func parseInput(lines []string) []Serie {
	series := make([]Serie, len(lines))
	for i, line := range lines {
		numbersStr := strings.Fields(line)
		serie := make(Serie, len(numbersStr))
		for j, n := range numbersStr {
			n, err := strconv.Atoi(n)
			common.ErrPanic(err)
			serie[j] = n
		}
		series[i] = serie
	}
	return series
}

func Part1() int {
	lines := common.ReadFile("day9/input")
	series := parseInput(lines)
	total := 0
	for _, serie := range series {
		total += serie.calculateNextValue()
	}
  return total
}
