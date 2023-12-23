package day12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

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

var cache map[int]map[int][]string = make(map[int]map[int][]string)
var genDuration time.Duration = 0
var countDuration time.Duration = 0

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
	missingMap, ok := cache[missing]
	if !ok {
		cache[missing] = make(map[int][]string)
	}
	if damagedCache, ok := missingMap[damaged]; ok {
		return damagedCache
	}
	generateCombinations("", missing, damaged, &result)
	cache[missing][damaged] = result
	return result
}

func (c *Condition) countCombinations() int {
	start := time.Now()
	combs := c.possibleCombinations()
	genDuration += time.Since(start)

	start = time.Now()
	re := regexp.MustCompile(`\?`)
	combinations := make([]string, len(combs))
	for i, comb := range combs {
		combIdx := -1
		newStr := re.ReplaceAllStringFunc(c.record, func(match string) string {
			combIdx++
			return string(comb[combIdx])
		})
		combinations[i] = newStr
	}
	countDuration += time.Since(start)
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
	fmt.Println("Generating:", genDuration)
	fmt.Println("Counting:", countDuration)
	return total
}
