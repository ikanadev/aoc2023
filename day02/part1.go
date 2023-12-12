package day02

import (
	"strconv"
	"strings"

	"github.com/jqvk/aoc2023/common"
)

// red green blue
type Set struct {
	Red   int
	Green int
	Blue  int
}

func (s *Set) isSubSet(pool Set) bool {
	return s.Red <= pool.Red && s.Green <= pool.Green && s.Blue <= pool.Blue
}
func (s *Set) getPower() int {
	return s.Red * s.Green * s.Blue
}

type Game struct {
	ID   int
	Sets []Set
}

func (g *Game) canPlay(pool Set) bool {
	for _, set := range g.Sets {
		if !set.isSubSet(pool) {
			return false
		}
	}
	return true
}

func (g *Game) getMinimumSet() Set {
	var result Set
	for _, set := range g.Sets {
		if set.Red > result.Red {
			result.Red = set.Red
		}
		if set.Green > result.Green {
			result.Green = set.Green
		}
		if set.Blue > result.Blue {
			result.Blue = set.Blue
		}
	}
	return result
}

func parseGame(line string) Game {
	var game Game
	gameLine := strings.Split(line, ": ")
	// parse ID
	gameID, err := strconv.Atoi(strings.Split(gameLine[0], " ")[1])
	common.ErrPanic(err)
	game.ID = gameID
	// parse sets
	sets := strings.Split(gameLine[1], ";")
	for _, setStr := range sets {
		var set Set
		colors := strings.Split(setStr, ",")
		for _, color := range colors {
			// color is something like "4 red"
			colorComponents := strings.Split(strings.TrimSpace(color), " ")
			colorName := colorComponents[1]
			colorQtty, err := strconv.Atoi(colorComponents[0])
			common.ErrPanic(err)
			switch colorName {
			case "red":
				set.Red = colorQtty
			case "green":
				set.Green = colorQtty
			case "blue":
				set.Blue = colorQtty
			}
		}
		game.Sets = append(game.Sets, set)
	}
	return game
}

func parseData() []Game {
	// lines := common.ReadFile("day02/input_small")
	lines := common.ReadFile("day02/input")
	games := make([]Game, len(lines))
	for i, line := range lines {
		games[i] = parseGame(line)
	}
	return games
}

func Part1() int {
	pool := Set{Red: 12, Green: 13, Blue: 14}
	games := parseData()
	idSum := 0
	for _, game := range games {
		if game.canPlay(pool) {
			idSum += game.ID
		}
	}
	return idSum
}
