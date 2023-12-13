package day12

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jqvk/aoc2023/common"
)

const (
	DAMAGED     = "#"
	OPERATIONAL = "."
)

func parseInput(lines []string) []Condition {
	var conds []Condition
	for _, line := range lines {
		var cond Condition
		split := strings.Fields(line)
		cond.record = split[0]
		damageStr := strings.Split(split[1], ",")
		damageList := make([]int, len(damageStr))
		for i, str := range damageStr {
			n, err := strconv.Atoi(str)
			common.ErrPanic(err)
			damageList[i] = n
		}
		cond.damageList = damageList
		conds = append(conds, cond)
	}
	return conds
}

func generateCombinations(prefix string, length, remainingDamaged int, result *[]string) {
	if length == 0 {
		if remainingDamaged == 0 {
			*result = append(*result, prefix)
		}
		return
	}
	generateCombinations(prefix+DAMAGED, length-1, remainingDamaged-1, result)
	generateCombinations(prefix+OPERATIONAL, length-1, remainingDamaged, result)
}

type Condition struct {
	record     string
	damageList []int
}

func (c *Condition) missingDamaged() int {
	total := 0
	for _, n := range c.damageList {
		total += n
	}
	return total - strings.Count(c.record, DAMAGED)
}

func (c *Condition) possibleCombinations() []string {
	damaged := c.missingDamaged()
	missing := strings.Count(c.record, "?")
	var result []string
	generateCombinations("", missing, damaged, &result)
	return result
}

func (c *Condition) countCombinations() int {
	combs := c.possibleCombinations()
	re := regexp.MustCompile(`\?`)
	var combinations []string
	for i := range combs {
		comb := []rune(combs[i])
		newStr := re.ReplaceAllStringFunc(c.record, func(match string) string {
			if len(comb) > 0 {
				char := comb[0]
				comb = comb[1:]
				return string(char)
			}
			return match
		})
		combinations = append(combinations, newStr)
	}
	var valid []string
	for _, comb := range combinations {
		split := strings.Fields(strings.ReplaceAll(comb, OPERATIONAL, " "))
		if len(split) == len(c.damageList) {
			equal := true
			for i, n := range c.damageList {
				if n != len(split[i]) {
					equal = false
				}
			}
			if equal {
				valid = append(valid, comb)
			}
		}
	}
	return len(valid)
}

func Part1() int {
	// lines := common.ReadFile("day12/input_small")
	lines := common.ReadFile("day12/input")
	conds := parseInput(lines)
	total := 0
	for _, c := range conds {
		total += c.countCombinations()
	}
	return total
}
