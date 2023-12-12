package day10

import "github.com/jqvk/aoc2023/common"

func calculateIntermediate(prev, current Coord) Coord {
	newCoord := Coord{current.i, current.j}
	if prev.i == current.i {
		newCoord.j = prev.j + 1
		if current.j < prev.j {
			newCoord.j = prev.j - 1
		}
	}
	if prev.j == current.j {
		newCoord.i = prev.i + 1
		if current.i < prev.i {
			newCoord.i = prev.i - 1
		}
	}
	return newCoord
}

func Part2() int {
	lines := common.ReadFile("day10/input")
	// lines := common.ReadFile("day10/input_part2")
	field := parseInput(lines)
	startingCoords := field.getStartingCoords()
	var path []Coord
	coord := startingCoords[0]
	path = append(path, field.origin)
	for !coord.isEqual(field.origin) {
		next := field.nextPipe(path[len(path)-1], coord)
		path = append(path, coord)
		coord = next
	}
	// double matrix
	matrixSize := len(field.pipes)*2 - 1
	matrix := make([][]rune, matrixSize)
	lastIdx := len(matrix) - 1
	for i := 0; i < matrixSize; i++ {
		row := make([]rune, matrixSize)
		for j := 0; j < matrixSize; j++ {
			row[j] = '-'
		}
		matrix[i] = row
	}
	doubledPath := make([]Coord, len(path)*2)
	for i, coord := range path {
		doubledPath[i*2] = Coord{coord.i * 2, coord.j * 2}
		if i > 0 {
			idx := i*2 - 1
			current := doubledPath[i*2]
			prev := doubledPath[(i-1)*2]
			doubledPath[idx] = calculateIntermediate(prev, current)
		}
	}
	doubledPath[len(doubledPath)-1] = calculateIntermediate(doubledPath[0], doubledPath[len(doubledPath)-2])
	for _, c := range doubledPath {
		matrix[c.i][c.j] = '#'
	}
	outerCoords := make([]Coord, 0)
	// get external coords
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == '-' {
			outerCoords = append(outerCoords, Coord{0, i})
		}
		if matrix[lastIdx][i] == '-' {
			outerCoords = append(outerCoords, Coord{lastIdx, i})
		}
		// vartical
		if i > 0 && i < len(matrix)-2 {
			if matrix[i][0] == '-' {
				outerCoords = append(outerCoords, Coord{i, 0})
			}
			if matrix[i][lastIdx] == '-' {
				outerCoords = append(outerCoords, Coord{i, lastIdx})
			}
		}
	}

	for len(outerCoords) > 0 {
		for _, c := range outerCoords {
			matrix[c.i][c.j] = '*'
		}
		newOuterCoords := make([]Coord, 0)
		for _, c := range outerCoords {
			top := Coord{c.i - 1, c.j}
			if top.i >= 0 && matrix[top.i][top.j] == '-' {
				newOuterCoords = append(newOuterCoords, top)
			}
			bot := Coord{c.i + 1, c.j}
			if bot.i <= lastIdx && matrix[bot.i][bot.j] == '-' {
				newOuterCoords = append(newOuterCoords, bot)
			}
			left := Coord{c.i, c.j - 1}
			if left.j >= 0 && matrix[left.i][left.j] == '-' {
				newOuterCoords = append(newOuterCoords, left)
			}
			right := Coord{c.i, c.j + 1}
			if right.j <= lastIdx && matrix[right.i][right.j] == '-' {
				newOuterCoords = append(newOuterCoords, right)
			}
		}
		outerCoords = newOuterCoords
	}

	counter := 0
	for i := 0; i < lastIdx; i += 2 {
		for j := 0; j < lastIdx; j += 2 {
			if matrix[i][j] == '-' {
				counter++
			}
		}
	}
	return counter
}
