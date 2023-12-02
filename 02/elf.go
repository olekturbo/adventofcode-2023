package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	gms := extractGames(utils.ReadInput("input.txt"))
	fmt.Printf("A: %v\n", count(gms))
	fmt.Printf("B: %v\n", countB(gms))
}

var greenexp = regexp.MustCompile("([0-9]+) green")
var blueexp = regexp.MustCompile("([0-9]+) blue")
var redexp = regexp.MustCompile("([0-9]+) red")

type subset map[string]int

type games map[int][]subset

func countB(g games) int {
	var s int
	for _, game := range g {
		reds := make([]int, 0)
		greens := make([]int, 0)
		blues := make([]int, 0)
		for _, s := range game {
			reds = append(reds, s["red"])
			greens = append(greens, s["green"])
			blues = append(blues, s["blue"])
		}
		s += slices.Max(reds) * slices.Max(greens) * slices.Max(blues)
	}

	return s
}

func count(g games) int {
	var s int
	for id, game := range g {
		if possible(game) {
			s += id
		}
	}
	return s
}

func possible(g []subset) bool {
	for _, subsets := range g {
		if subsets["blue"] > 14 || subsets["green"] > 13 || subsets["red"] > 12 {
			return false
		}
	}
	return true
}

func extractGames(s []string) games {
	gameMap := make(games, 0)
	for _, l := range s {
		ll := strings.Split(l, ":")
		g := strings.Replace(ll[0], "Game ", "", 1)
		id, err := strconv.Atoi(g)
		if err != nil {
			panic(err)
		}

		subsets := make([]subset, 0)
		for _, lll := range strings.Split(ll[1], ";") {
			m := make(subset, 0)
			m["green"] = findNumber(lll, greenexp)
			m["red"] = findNumber(lll, redexp)
			m["blue"] = findNumber(lll, blueexp)
			subsets = append(subsets, m)
		}

		gameMap[id] = subsets
	}
	return gameMap
}

func findNumber(s string, reg *regexp.Regexp) int {
	if gr := reg.FindStringSubmatch(s); len(gr) > 1 {
		grid, err := strconv.Atoi(gr[1])
		if err != nil {
			panic(err)
		}
		return grid
	}
	return 0
}
