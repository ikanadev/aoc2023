package day10

import "github.com/jqvk/aoc2023/common"

type Coord struct {
	i int
	j int
}

func (c *Coord) isEqual(coord Coord) bool {
	return c.i == coord.i && c.j == coord.j
}

type Field struct {
	pipes  [][]rune
	origin Coord
}

func (f *Field) isValidCoord(c Coord) bool {
	if c.i < 0 || c.j < 0 {
		return false
	}
	if c.i >= len(f.pipes) || c.j >= len(f.pipes[0]) {
		return false
	}
	return true
}

func (f *Field) getStartingCoords() [2]Coord {
	var coords [2]Coord
	activeIdx := 0
	// checking top -> right -> bottom -> left
	toCheck := Coord{f.origin.i - 1, f.origin.j}
	if f.isValidCoord(toCheck) {
		char := f.pipes[toCheck.i][toCheck.j]
		if char == '|' || char == '7' || char == 'F' {
			coords[activeIdx] = Coord{toCheck.i, toCheck.j}
			activeIdx++
		}
	}
	toCheck.i = f.origin.i
	toCheck.j = f.origin.j + 1
	if f.isValidCoord(toCheck) {
		char := f.pipes[toCheck.i][toCheck.j]
		if char == '-' || char == 'J' || char == '7' {
			coords[activeIdx] = Coord{toCheck.i, toCheck.j}
			activeIdx++
		}
	}
	toCheck.i = f.origin.i + 1
	toCheck.j = f.origin.j
	if f.isValidCoord(toCheck) {
		char := f.pipes[toCheck.i][toCheck.j]
		if char == '|' || char == 'J' || char == 'L' {
			coords[activeIdx] = Coord{toCheck.i, toCheck.j}
			activeIdx++
		}
	}
	toCheck.i = f.origin.i
	toCheck.j = f.origin.j - 1
	if f.isValidCoord(toCheck) {
		char := f.pipes[toCheck.i][toCheck.j]
		if char == '-' || char == 'F' || char == 'L' {
			coords[activeIdx] = Coord{toCheck.i, toCheck.j}
			activeIdx++
		}
	}
	return coords
}

func parseInput(lines []string) Field {
	var field Field
	var origin Coord
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		row := make([]rune, len(line))
		for j, char := range line {
			if char == 'S' {
				origin.i = i
				origin.j = j
			}
			row[j] = char
		}
		matrix[i] = row
	}
	field.pipes = matrix
	field.origin = origin
	return field
}

/*
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
*/
func (f *Field) nextPipe(prev, current Coord) Coord {
	next := Coord{current.i, current.j}
	char := f.pipes[current.i][current.j]
	switch char {
	case '|':
		{
			next.i = current.i + 1
			if prev.i == next.i {
				next.i = current.i - 1
			}
		}
	case '-':
		{
			next.j = current.j + 1
			if prev.j == next.j {
				next.j = current.j - 1
			}
		}
	case 'L':
		{
			next.j = current.j + 1
			if prev.j == next.j {
				next.j = current.j
				next.i = current.i - 1
			}
		}
	case 'J':
		{
			next.j = current.j - 1
			if prev.j == next.j {
				next.j = current.j
				next.i = current.i - 1
			}
		}
	case '7':
		{
			next.j = current.j - 1
			if prev.j == next.j {
				next.j = current.j
				next.i = current.i + 1
			}
		}
	case 'F':
		{
			next.j = current.j + 1
			if prev.j == next.j {
				next.j = current.j
				next.i = current.i + 1
			}
		}
	}
	return next
}

func Part1() int {
	lines := common.ReadFile("day10/input")
	// lines := common.ReadFile("day10/input_extra")
	field := parseInput(lines)
	startingCoords := field.getStartingCoords()
	var path1 []Coord
	var path2 []Coord
	coord1 := startingCoords[0]
	coord2 := startingCoords[1]
	path1 = append(path1, field.origin)
	path2 = append(path2, field.origin)
	for !coord1.isEqual(coord2) {
		next1 := field.nextPipe(path1[len(path1)-1], coord1)
		path1 = append(path1, coord1)
		coord1 = next1
		next2 := field.nextPipe(path2[len(path2)-1], coord2)
		path2 = append(path2, coord2)
		coord2 = next2
	}
	path1 = append(path1, coord1)
	path2 = append(path2, coord2)
	return len(path1) - 1
}
