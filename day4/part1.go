package day4

import (
	"math"
	"strconv"
	"strings"

	"github.com/jqvk/aoc2023/common"
)

type Set struct {
	id          int
	winners     []int
	cards       []int
	winnerCards int
	instances   int
}

func (s *Set) calculateWinnerCards() {
	winerCards := 0
	for _, card := range s.cards {
		for _, winner := range s.winners {
			if card == winner {
				winerCards++
				break
			}
		}
	}
	s.winnerCards = winerCards
}

func parseInput() []Set {
	// lines := common.ReadFile("day4/input_small")
	lines := common.ReadFile("day4/input")
	var sets []Set
	for _, line := range lines {
		var set Set
		set.instances = 1
		split := strings.Split(line, ":")
		setSplit := strings.Fields(split[0])
		cardsSplit := strings.Split(split[1], "|")
		id, err := strconv.Atoi(setSplit[1])
		common.ErrPanic(err)
		set.id = id
		winners := strings.Fields(strings.TrimSpace(cardsSplit[0]))
		cards := strings.Fields(strings.TrimSpace(cardsSplit[1]))
		for _, str := range winners {
			n, err := strconv.Atoi(str)
			common.ErrPanic(err)
			set.winners = append(set.winners, n)
		}
		for _, str := range cards {
			n, err := strconv.Atoi(str)
			common.ErrPanic(err)
			set.cards = append(set.cards, n)
		}
		sets = append(sets, set)
	}
	return sets
}

func Part1() int {
	sets := parseInput()
	totalPoints := 0
	for i := range sets {
		sets[i].calculateWinnerCards()
		if sets[i].winnerCards > 0 {
			totalPoints += int(math.Pow(2, float64(sets[i].winnerCards-1)))
		}
	}
	return totalPoints
}
