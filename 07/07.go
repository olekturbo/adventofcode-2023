package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"strconv"
	"strings"
)

var cards = map[string]int{
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

const (
	HighCard = iota
	Pair
	TwoPairs
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

type hand struct {
	values []int
	bet    int
	rank   int
}

func main() {
	parsed := parse(utils.ReadInput("input.txt"))
	fmt.Printf("A: %v\n", bid(parsed))
}

func bid(hands []hand) int {
	ranks := make(map[int][]hand, 0)

	for _, h := range hands {
		ranks[h.rank] = append(ranks[h.rank], h)
	}

	var sum int
	j := len(hands)
	for i := FiveOfKind; i >= HighCard; i-- {
		if v, ok := ranks[i]; ok {
			v = sort(v, len(v))
			for _, vv := range v {
				sum += vv.bet * j
				j--
			}
		}
	}

	return sum
}

func sort(hands []hand, l int) []hand {
	if l == 0 {
		return hands
	}
	for i := 1; i < len(hands); i++ {
		if higher(hands[i], hands[i-1]) {
			cp := hands[i-1]
			hands[i-1] = hands[i]
			hands[i] = cp
		}
	}
	return sort(hands, l-1)
}

func higher(h hand, hh hand) bool {
	for i := 0; i < 5; i++ {
		if h.values[i] == hh.values[i] {
			continue
		}
		if h.values[i] > hh.values[i] {
			return true
		} else {
			return false
		}
	}
	return false
}

func rank(values []int) int {
	m := make(map[int]int, 0)
	for _, v := range values {
		m[v]++
	}

	ret := make([]int, 0)

	for _, v := range m {
		if v == 5 {
			ret = append(ret, FiveOfKind)
		}
		if v == 4 {
			ret = append(ret, FourOfKind)
		}
		if v == 3 {
			ret = append(ret, ThreeOfKind)
		}
		if v == 2 {
			ret = append(ret, Pair)
		}
	}

	if inArray(ret, FiveOfKind) {
		return FiveOfKind
	}

	if inArray(ret, FourOfKind) {
		return FourOfKind
	}

	if inArray(ret, ThreeOfKind) && inArray(ret, Pair) {
		return FullHouse
	}

	if inArray(ret, ThreeOfKind) && len(ret) == 1 {
		return ThreeOfKind
	}

	if inArray(ret, Pair) && len(ret) == 2 {
		return TwoPairs
	}

	if inArray(ret, Pair) && len(ret) == 1 {
		return Pair
	}

	return HighCard
}

func inArray(i []int, v int) bool {
	for _, ii := range i {
		if ii == v {
			return true
		}
	}
	return false
}

func parse(input []string) []hand {
	hands := make([]hand, 0, len(input))
	for _, s := range input {
		sp := strings.Split(s, " ")
		bet, _ := strconv.Atoi(sp[1])

		values := make([]int, 0, 5)
		for _, card := range sp[0] {
			k := string(card)
			val, err := strconv.Atoi(k)
			if err == nil {
				values = append(values, val)
			} else {
				values = append(values, cards[k])
			}
		}
		hands = append(hands, hand{bet: bet, values: values, rank: rank(values)})
	}
	return hands
}
