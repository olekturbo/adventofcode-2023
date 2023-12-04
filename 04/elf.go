package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type scratch struct {
	id      int
	winning []int
	owned   []int
}

type copy struct {
	id    int
	count int
}

func main() {
	input := utils.ReadInput("input.txt")
	scratches := toScratches(input)
	fmt.Printf("A: %v\n", sum(scratches))
	fmt.Printf("B: %v\n", copies(scratches))
}

func copies(scratches []scratch) int {
	c := make([]int, len(scratches))

	for i := range scratches {
		c[i] = 1
	}

	for i := range scratches {
		for j := i + 1; j < i+1+count(scratches[i].winning, scratches[i].owned); j++ {
			c[j] += c[i]
		}
	}

	t := 0
	for _, cc := range c {
		t += cc
	}

	return t
}

func sum(scratches []scratch) int {
	var ret int
	for _, s := range scratches {
		ret += worth(count(s.winning, s.owned))
	}
	return ret
}

func worth(c int) int {
	return int(math.Pow(2, float64(c-1)))
}

func count(w []int, o []int) int {
	c := 0
	for _, oo := range o {
		for _, ww := range w {
			if oo == ww {
				c++
			}
		}
	}
	return c
}

func toScratches(s []string) []scratch {
	ret := make([]scratch, 0, len(s))

	for _, ss := range s {
		sp := strings.Split(ss, ":")
		id := strings.Replace(sp[0], "Card ", "", 1)
		nmbs := strings.Split(sp[1], "|")
		winning := strings.Split(nmbs[0], " ")
		owned := strings.Split(nmbs[1], " ")
		ret = append(ret, scratch{
			id:      stringToInt(id),
			winning: mustSliceStringToInt(winning),
			owned:   mustSliceStringToInt(owned),
		})
	}

	return ret
}

func mustSliceStringToInt(s []string) []int {
	ret := make([]int, 0, len(s))
	for _, ss := range s {
		if v := stringToInt(ss); v != 0 {
			ret = append(ret, v)
		}
	}
	return ret
}

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
