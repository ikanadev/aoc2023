package day1

import "strings"

var numbersMap map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// -1 means not found
func startsWithNumber(str string) int {
	for name, number := range numbersMap {
		if strings.HasPrefix(str, name) {
			return number
		}
	}
	return -1
}
func endsWithNumber(str string) int {
	for name, number := range numbersMap {
		if strings.HasSuffix(str, name) {
			return number
		}
	}
	return -1
}

func getLineSumWithStrings(line string) int {
	first := -1
	last := -1
	index := 0

	for true {
		if first >= 0 && last >= 0 {
			break
		}
		if first == -1 && isDigit(line[index]) {
			first = int(line[index] - '0')
		}
		if last == -1 && isDigit(line[len(line)-1-index]) {
			last = int(line[len(line)-1-index] - '0')
		}
		if possibleFirst := startsWithNumber(line[index:]); first == -1 && possibleFirst != -1 {
			first = possibleFirst
		}
		if possibleLast := endsWithNumber(line[:len(line)-index]); last == -1 && possibleLast != -1 {
			last = possibleLast
		}
		index++
	}

	return first*10 + last
}

func Part2() int {
	// content := readInput("day1/input_small2")
	content := readInput("day1/input")
	total := 0
	for _, line := range content {
		sum := getLineSumWithStrings(line)
		total += sum
	}
	return total
}
