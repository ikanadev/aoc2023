package day11

import "github.com/jqvk/aoc2023/common"

func Part2() int {
	// lines := common.ReadFile("day11/input_small")
	lines := common.ReadFile("day11/input")
	universe := parseInput(lines)
	sum := 0
	emtpySpaceValue := 1000000 - 1
	for i := 0; i < len(universe.galaxies)-1; i++ {
		for j := i + 1; j < len(universe.galaxies); j++ {
			g1 := universe.galaxies[i]
			g2 := universe.galaxies[j]
			sum += g1.distance(g2) + universe.emptySpaces(g1, g2)*emtpySpaceValue
		}
	}
	println(sum)
	return sum
}
