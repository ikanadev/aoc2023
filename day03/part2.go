package day03

import (
	"strconv"
	"unicode"

	"github.com/jqvk/aoc2023/common"
)

type Part struct {
	number  int
	engines []Coord
}

func getEnginesAround(rowIndex, from, to int, matrix [][]rune) []Coord {
	var engines []Coord
	coordToCheck := getCoordsToCheck(rowIndex, from, to, matrix)
	for _, coord := range coordToCheck {
		char := matrix[coord.row][coord.col]
		if char == '*' {
			engines = append(engines, coord)
		}
	}
	return engines
}

func getEngineParts(parts []Part, engine Coord) []Part {
	var engineParts []Part
	for _, part := range parts {
		for _, e := range part.engines {
			if e.row == engine.row && e.col == engine.col {
				engineParts = append(engineParts, part)
			}
		}
	}
	return engineParts
}

func calculateGearsRatio(parts [][]Part, engines [][]Coord, matrix [][]rune) int {
	var sum int
	for i, row := range engines {
		for _, engine := range row {
			var engineParts []Part
			if i > 0 {
				engineParts = append(engineParts, getEngineParts(parts[i-1], engine)...)
			}
			engineParts = append(engineParts, getEngineParts(parts[i], engine)...)
			if i < (len(engines) - 1) {
				engineParts = append(engineParts, getEngineParts(parts[i+1], engine)...)
			}
			if len(engineParts) == 2 {
				sum += (engineParts[0].number * engineParts[1].number)
			}
		}
	}

	return sum
}
func getPartsAndEngines(matrix [][]rune) ([][]Part, [][]Coord) {
	parts := make([][]Part, len(matrix))
	engines := make([][]Coord, len(matrix))
	var start, end int
	var number string
	var parsingNumber bool
	for i, row := range matrix {
		start, end = 0, 0
		number, parsingNumber = "", false
		for j, char := range row {
			if char == '*' {
				engines[i] = append(engines[i], Coord{row: i, col: j})
			}
			if unicode.IsDigit(char) {
				number += string(char)
			}
			if unicode.IsDigit(char) && !parsingNumber {
				parsingNumber = true
				start = j
			}
			if !unicode.IsDigit(char) && parsingNumber {
				end = j - 1
				// check
				enginesAround := getEnginesAround(i, start, end, matrix)
				n, err := strconv.Atoi(number)
				common.ErrPanic(err)
				part := Part{number: n, engines: enginesAround}
				parts[i] = append(parts[i], part)
				number, start, end, parsingNumber = "", 0, 0, false
			}
		}
		if parsingNumber {
			enginesAround := getEnginesAround(i, start, len(row)-1, matrix)
			n, err := strconv.Atoi(number)
			common.ErrPanic(err)
			part := Part{number: n, engines: enginesAround}
			parts[i] = append(parts[i], part)
		}
	}
	return parts, engines
}

func Part2() int {
	matrix := parseInput()
	parts, engines := getPartsAndEngines(matrix)
	return calculateGearsRatio(parts, engines, matrix)
}
