package day09

import "github.com/jqvk/aoc2023/common"

func (s *Serie) calculatePrevValue() int {
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
	prevValue := 0
	for i := len(matrix) - 2; i >= 0; i-- {
		prevValue = matrix[i][0] - prevValue
	}
	return prevValue
}

func Part2() int {
	// lines := common.ReadFile("day09/input_small")
	lines := common.ReadFile("day09/input")
	series := parseInput(lines)
	total := 0
	for _, serie := range series {
		total += serie.calculatePrevValue()
	}
	return total
}
