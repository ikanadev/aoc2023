package day7

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jqvk/aoc2023/common"
)

var cardStrength = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

/*

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
[{32T3K 1 765} {T55J5 3 684} {KK677 2 28} {KTJJT 2 220} {QQQJA 3 483}]
   Five of a kind, where all five cards have the same label: AAAAA
   Four of a kind, where four cards have the same label and one card has a different label: AA8AA
   Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
   Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
   Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
   One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
   High card, where all cards' labels are distinct: 23456
*/

type Hand struct {
	cards string
	// 0:high card,1:one pair,2:two pair,3:threeof kind,4:full house,5:four of a kind,6:five of a kind
	kind int
	bid  int
}

func (h *Hand) calculateKind() {
	runeMap := make(map[rune]int)
	for _, c := range h.cards {
		runeMap[c]++
	}
	// five of a kind
	if len(runeMap) == 1 {
		h.kind = 6
		return
	}
	if len(runeMap) == 2 {
		for _, l := range runeMap {
			// four of a kind
			if l == 4 {
				h.kind = 5
				return
			}
			// full house
			if l == 3 || l == 2 {
				h.kind = 4
				return
			}
		}
	}
	if len(runeMap) == 3 {
		for _, l := range runeMap {
			// three of a kind
			if l == 3 {
				h.kind = 3
				return
			}
			// two pair
			if l == 2 {
				h.kind = 2
				return
			}
		}
	}
	if len(runeMap) == 4 {
		h.kind = 1
		return
	}
	h.kind = 0
}

func parseInput(lines []string) []Hand {
	var hands []Hand
	for _, line := range lines {
		var hand Hand
		lineSplit := strings.Fields(line)
		bid, err := strconv.Atoi(lineSplit[1])
		common.ErrPanic(err)
		hand.bid = bid
		hand.cards = lineSplit[0]
		hand.calculateKind()
		hands = append(hands, hand)
	}
	return hands
}

func Part1() int {
	// lines := common.ReadFile("day7/input_small")
	lines := common.ReadFile("day7/input")
	hands := parseInput(lines)
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].kind == hands[j].kind {
			for k, c := range hands[i].cards {
				iCardVal := cardStrength[c]
				jCardVal := cardStrength[rune(hands[j].cards[k])]
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
