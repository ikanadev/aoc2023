package main

import (
	"fmt"
	"time"

	"github.com/jqvk/aoc2023/day1"
	"github.com/jqvk/aoc2023/day2"
	"github.com/jqvk/aoc2023/day3"
	"github.com/jqvk/aoc2023/day4"
	"github.com/jqvk/aoc2023/day5"
	"github.com/jqvk/aoc2023/day6"
	"github.com/jqvk/aoc2023/day7"
	"github.com/jqvk/aoc2023/day8"
)

func measureAndPrint[T any](label string, fn func() T) {
	startTime := time.Now()
	res := fn()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("%s: %-15v\t(%v)\n", label, res, elapsed)
}

func main() {
	measureAndPrint("D1P1", day1.Part1)
	measureAndPrint("D1P2", day1.Part2)
	measureAndPrint("D2P1", day2.Part1)
	measureAndPrint("D2P2", day2.Part2)
	measureAndPrint("D3P1", day3.Part1)
	measureAndPrint("D3P2", day3.Part2)
	measureAndPrint("D4P1", day4.Part1)
	measureAndPrint("D4P2", day4.Part2)
	measureAndPrint("D5P1", day5.Part1)
	measureAndPrint("D5P2", day5.Part2)
	measureAndPrint("D6P1", day6.Part1)
	measureAndPrint("D6P2", day6.Part2)
	measureAndPrint("D7P1", day7.Part1)
	measureAndPrint("D7P2", day7.Part2)
	measureAndPrint("D8P1", day8.Part1)
	measureAndPrint("D8P2", day8.Part2)
}
