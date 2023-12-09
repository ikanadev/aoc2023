package day7

import (
	"sort"
	"strings"

	"github.com/jqvk/aoc2023/common"
)

var cardStrengthWithJoker = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': -1,
	'Q': 10,
	'K': 11,
	'A': 12,
}

/*
1 J
23456 2345J 23455 0 -> 1
A23A4 A23AJ A23AA 1 -> 3
23432 23J32 23332 2 -> 4
TTT98 TTT9J TTT9T 3 -> 5
23332
AA8AA AAJAA AAAAA 5 -> 6
2 J
23456
A23A4 J23J4 42344 1 -> 3
23432 2J4J2 22322 2 -> 5
TTT98
23332 J333J 33333 4 -> 6
AA8AA
3 J
23456
A23A4
23432
TTT98 JJJ98 99998 3 -> 5
23332 2JJJ2 22222 4 -> 6
AA8AA
4 J
23456
A23A4
23432
TTT98
23332
AA8AA JJ8JJ 88888 5 -> 6
*/

func (h *Hand) calculateKindWithJokers() {
	jokers := strings.Count(h.cards, "J")
	if jokers == 0 {
		return
	}
	if jokers == 1 {
		if h.kind == 0 || h.kind == 5 {
			h.kind += 1
			return
		}
		if h.kind >= 1 && h.kind <= 3 {
			h.kind += 2
			return
		}
	}
	if jokers == 2 {
		if h.kind == 1 || h.kind == 4 {
			h.kind += 2
			return
		}
		if h.kind == 2 {
			h.kind = 5
			return
		}
	}
	if jokers == 3 {
		if h.kind == 3 || h.kind == 4 {
			h.kind += 2
			return
		}
	}
	if jokers == 4 && h.kind == 5 {
		h.kind = 6
	}
}

func Part2() int {
	// lines := common.ReadFile("day7/input_small")
	lines := common.ReadFile("day7/input")
	// lines := common.ReadFile("day7/input_extra")
	hands := parseInput(lines)
	for i := range hands {
		hands[i].calculateKindWithJokers()
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].kind == hands[j].kind {
			for k, c := range hands[i].cards {
				iCardVal := cardStrengthWithJoker[c]
				jCardVal := cardStrengthWithJoker[rune(hands[j].cards[k])]
				if iCardVal != jCardVal {
					return iCardVal < jCardVal
				}
			}
		}
		return hands[i].kind < hands[j].kind
	})
	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}
	return total
}
