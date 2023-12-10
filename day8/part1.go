package day8

import "github.com/jqvk/aoc2023/common"

type Dir string

const (
	Left  Dir = "L"
	Right Dir = "R"
)

type MapDirs map[string]map[Dir]string

type Map struct {
	instructions string
	insIndex     int
	position     string
	mapDirs      MapDirs
}

func (m *Map) nextInstruction() Dir {
	m.insIndex = (m.insIndex + 1) % len(m.instructions)
	if m.instructions[m.insIndex] == 'L' {
		return Left
	}
	return Right
}

func (m *Map) restarInstructions() {
	m.insIndex = -1
}

func parseInput(lines []string) Map {
	var m Map
	m.instructions = lines[0]
	m.insIndex = -1
	m.mapDirs = make(MapDirs)
	insLines := lines[2:]
	for _, line := range insLines {
		position := line[0:3]
		left := line[7:10]
		right := line[12:15]
		m.mapDirs[position] = map[Dir]string{
			Left:  left,
			Right: right,
		}
	}
	return m
}

func Part1() int {
	lines := common.ReadFile("day8/input")
	m := parseInput(lines)
	m.position = "AAA"

	steps := 0
	for m.position != "ZZZ" {
		steps++
		ins := m.nextInstruction()
		m.position = m.mapDirs[m.position][ins]
	}
	return steps
}
