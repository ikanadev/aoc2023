package day04

func Part2() int {
	sets := parseInput()
	for i := range sets {
		sets[i].calculateWinnerCards()
	}
	limit := len(sets) - 1
	for i := range sets {
		wins := sets[i].winnerCards
		for times := 1; times <= sets[i].instances; times++ {
			for j := i + 1; j <= i+wins; j++ {
				if j <= limit {
					sets[j].instances++
				}
			}
		}
	}
	totalInstances := 0
	for i := range sets {
		totalInstances += sets[i].instances
	}
	return totalInstances
}
