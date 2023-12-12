package day11

import (
	"math"

	"github.com/jqvk/aoc2023/common"
)

type Coord struct {
	i int
	j int
}

func (c *Coord) distance(coord Coord) int {
	iDistance := math.Abs(float64(c.i - coord.i))
	jDistance := math.Abs(float64(c.j - coord.j))
	return int(iDistance) + int(jDistance)
}

type Universe struct {
	galaxies []Coord
	// if a galaxy exists at some row index
	rows []bool
	// if a galaxy exists at some col index
	cols []bool
}

func (u *Universe) emptySpaces(c1, c2 Coord) int {
	spaces := 0
	start := c1.i + 1
	end := c2.i - 1
	if c1.i > c2.i {
		start = c2.i + 1
		end = c1.i - 1
	}
	for i := start; i <= end; i++ {
		if !u.rows[i] {
			spaces++
		}
	}
	start = c1.j + 1
	end = c2.j - 1
	if c1.j > c2.j {
		start = c2.j + 1
		end = c1.j - 1
	}
	for i := start; i <= end; i++ {
		if !u.cols[i] {
			spaces++
		}
	}
	return spaces
}

func parseInput(lines []string) Universe {
	universe := Universe{}
	universe.rows = make([]bool, len(lines))
	universe.cols = make([]bool, len(lines[0]))
	universe.galaxies = make([]Coord, 0)
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				coord := Coord{i, j}
				universe.galaxies = append(universe.galaxies, coord)
				universe.rows[i] = true
				universe.cols[j] = true
			}
		}
	}
	return universe
}

func Part1() int {
	// lines := common.ReadFile("day11/input_small")
	lines := common.ReadFile("day11/input")
	universe := parseInput(lines)
	sum := 0
	for i := 0; i < len(universe.galaxies)-1; i++ {
		for j := i + 1; j < len(universe.galaxies); j++ {
			g1 := universe.galaxies[i]
			g2 := universe.galaxies[j]
			sum += g1.distance(g2) + universe.emptySpaces(g1, g2)
		}
	}
	return sum
}
