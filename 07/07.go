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
	fmt.Printf("A: %v\n", bid(parse(utils.ReadInput("input.txt"), false)))
	fmt.Printf("B: %v\n", bid(parse(utils.ReadInput("input.txt"), true)))
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

func maxIdx(m map[int]int) int {
	idx := 0
	max := 0
	for k, v := range m {
		if k == 0 {
			if v == 5 {
				return v
			}
			continue
		}
		if v > max {
			max = v
			idx = k
		}
	}

	return idx
}

func rank(values []int) int {
	m := make(map[int]int, 0)
	for _, v := range values {
		m[v]++
	}

	mIdx := maxIdx(m)

	m[mIdx] += m[0]

	ret := make([]int, 0)

	for k, v := range m {
		if k == 0 {
			continue
		}
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

func parse(input []string, j bool) []hand {
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
				if j && k == "J" {
					values = append(values, 0)
					continue
				}
				values = append(values, cards[k])
			}
		}
		hands = append(hands, hand{bet: bet, values: values, rank: rank(values)})
	}
	return hands
}
