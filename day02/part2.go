package day02


func Part2() int {
	games := parseData()
  powerSum := 0
	for _, game := range games {
    set := game.getMinimumSet()
    power := set.getPower()
    powerSum += power
	}
  return powerSum
}
