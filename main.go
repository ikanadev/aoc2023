package main

import (
	"fmt"
	"time"

	"github.com/jqvk/aoc2023/day01"
	"github.com/jqvk/aoc2023/day02"
	"github.com/jqvk/aoc2023/day03"
	"github.com/jqvk/aoc2023/day04"
	"github.com/jqvk/aoc2023/day05"
	"github.com/jqvk/aoc2023/day06"
	"github.com/jqvk/aoc2023/day07"
	"github.com/jqvk/aoc2023/day08"
	"github.com/jqvk/aoc2023/day09"
	"github.com/jqvk/aoc2023/day10"
)

func measureAndPrint[T any](label string, fn func() T) {
	startTime := time.Now()
	res := fn()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("%s: %-15v\t(%v)\n", label, res, elapsed)
}

func main() {
	measureAndPrint("D01P1", day01.Part1)
	measureAndPrint("D01P2", day01.Part2)
	measureAndPrint("D02P1", day02.Part1)
	measureAndPrint("D02P2", day02.Part2)
	measureAndPrint("D03P1", day03.Part1)
	measureAndPrint("D03P2", day03.Part2)
	measureAndPrint("D04P1", day04.Part1)
	measureAndPrint("D04P2", day04.Part2)
	measureAndPrint("D05P1", day05.Part1)
	measureAndPrint("D05P2", day05.Part2)
	measureAndPrint("D06P1", day06.Part1)
	measureAndPrint("D06P2", day06.Part2)
	measureAndPrint("D07P1", day07.Part1)
	measureAndPrint("D07P2", day07.Part2)
	measureAndPrint("D08P1", day08.Part1)
	measureAndPrint("D08P2", day08.Part2)
	measureAndPrint("D09P1", day09.Part1)
	measureAndPrint("D09P2", day09.Part2)
	measureAndPrint("D10P1", day10.Part1)
	measureAndPrint("D10P2", day10.Part2)
}
