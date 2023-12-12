package day03

import (
	"strconv"
	"unicode"

	"github.com/jqvk/aoc2023/common"
)

type Coord struct {
	row int
	col int
}

func parseInput() [][]rune {
	// lines := common.ReadFile("day03/input_small")
	lines := common.ReadFile("day03/input")
	var matrix [][]rune
	for _, line := range lines {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func getCoordsToCheck(rowIndex, from, to int, matrix [][]rune) []Coord {
	maxRow, maxCol := len(matrix)-1, len(matrix[0])-1
	var initialList []Coord
	// above and botton row
	for col := from - 1; col <= to+1; col++ {
		initialList = append(
			initialList,
			Coord{row: rowIndex - 1, col: col},
			Coord{row: rowIndex + 1, col: col},
		)
	}
	initialList = append(initialList, Coord{row: rowIndex, col: from - 1}, Coord{row: rowIndex, col: to + 1})
	var cleanList []Coord
	for _, coord := range initialList {
		row, col := coord.row, coord.col
		if row >= 0 && row <= maxRow && col >= 0 && col <= maxCol {
			cleanList = append(cleanList, coord)
		}
	}
	return cleanList
}

func containsSpecialChars(coords []Coord, matrix [][]rune) bool {
	for _, coord := range coords {
		char := matrix[coord.row][coord.col]
		if char != '.' && !unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func calculatePartsSum(matrix [][]rune) int {
	var sum, start, end int
	var number string
	var parsingNumber bool
	for i, row := range matrix {
		start, end = 0, 0
		number, parsingNumber = "", false
		for j, char := range row {
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
				coordsToCheck := getCoordsToCheck(i, start, end, matrix)
				if containsSpecialChars(coordsToCheck, matrix) {
					n, err := strconv.Atoi(number)
					common.ErrPanic(err)
					sum += n
				}
				number, start, end, parsingNumber = "", 0, 0, false
			}
		}
		if parsingNumber {
			coordsToCheck := getCoordsToCheck(i, start, len(row)-1, matrix)
			if containsSpecialChars(coordsToCheck, matrix) {
				n, err := strconv.Atoi(number)
				common.ErrPanic(err)
				sum += n
			}
		}
	}
	return sum
}

func Part1() int {
	matrix := parseInput()
	return calculatePartsSum(matrix)
}
