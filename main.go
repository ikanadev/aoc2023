package main

import (
	"fmt"
	"time"

	"github.com/jqvk/aoc2023/day1"
	"github.com/jqvk/aoc2023/day2"
	"github.com/jqvk/aoc2023/day3"
	"github.com/jqvk/aoc2023/day4"
)

func measureAndPrint[T any](label string, fn func() T) {
	startTime := time.Now()
	res := fn()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("%s: %v\t\t(%v)\n", label, res, elapsed)
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
}
