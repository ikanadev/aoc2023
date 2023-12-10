package day8

import "github.com/jqvk/aoc2023/common"

func (m *Map) calculateStepsFrom(start string) int {
	m.position = start
	m.restarInstructions()
	steps := 0
	for m.position[2] != 'Z' {
		steps++
		ins := m.nextInstruction()
		m.position = m.mapDirs[m.position][ins]
	}
	return steps
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func listLcm(numbers []int) int {
	result := 1
	for _, n := range numbers {
		result = lcm(result, n)
	}
	return result
}

func Part2() int {
	lines := common.ReadFile("day8/input")
	m := parseInput(lines)
	list := make([]string, 0)
	for s := range m.mapDirs {
		if s[2] == 'A' {
			list = append(list, s)
		}
	}
	results := make([]int, len(list))
	for i, s := range list {
		steps := m.calculateStepsFrom(s)
		results[i] = steps
	}
	return listLcm(results)
}
