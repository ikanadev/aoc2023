package day13

import (
	"strings"

	"github.com/jqvk/aoc2023/common"
)

var maps [][]string

func getRowMirrorCount(matrix []string) int {
	rowsToCheck := 0
	len := len(matrix)
	for i := 0; i < len-1; i++ {
		rowsToCheck++
		if matrix[i] == matrix[i+1] {
			toCheck := rowsToCheck
			if rowsToCheck > len/2 {
				toCheck = len - rowsToCheck
			}
			equals := true
			for j := i; j > i-toCheck; j-- {
				if matrix[j] != matrix[i+i-j+1] {
					equals = false
				}
			}
			if equals {
				return i + 1
			}
		}
	}
	return -1
}

func rotate(matrix []string) []string {
	rows := len(matrix[0])
	res := make([]string, rows)
	for j := 0; j < rows; j++ {
		strBuilder := strings.Builder{}
		for i := 0; i < len(matrix); i++ {
			strBuilder.WriteByte(matrix[i][j])
		}
		res[j] = strBuilder.String()
	}
	return res
}

func Part1() int {
	// lines := common.ReadFile("day13/input_small")
	lines := common.ReadFile("day13/input")
	var m []string
	for i := range lines {
		if lines[i] == "" {
			maps = append(maps, m)
			m = []string{}
			continue
		}
		m = append(m, lines[i])
	}
	maps = append(maps, m)

	total := 0

	for i := range maps {
		matrix := maps[i]
		factor := 100
		linesUntilMirror := getRowMirrorCount(matrix)
		if linesUntilMirror == -1 {
			factor = 1
			linesUntilMirror = getRowMirrorCount(rotate(matrix))
		}
		total += linesUntilMirror * factor
	}
	return total
}
