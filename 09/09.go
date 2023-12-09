package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func extrapolate(s []int) [][]int {
	ret := make([][]int, len(s))

	ret[0] = s
	for i := 1; i < len(s); i++ {
		ret[i] = make([]int, len(s)-i)
		for j := 1; j <= len(s)-i; j++ {
			ret[i][j-1] = ret[i-1][j] - ret[i-1][j-1]
		}
	}

	return ret
}

func takeLast(s [][]int) []int {
	ret := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		ret[i] = s[i][len(s[i])-1]
	}

	return ret
}

func history(s []int) int {
	var ret int

	for i := 0; i < len(s); i++ {
		ret += s[i]
	}

	return ret
}

func parseInput(s []string) [][]int {
	ret := make([][]int, len(s))

	for i := 0; i < len(s); i++ {
		sp := strings.Split(s[i], " ")
		ret[i] = make([]int, len(sp))
		for j := 0; j < len(sp); j++ {
			ss := strings.TrimSpace(sp[j])
			ret[i][j] = func() int {
				v, _ := strconv.Atoi(ss)
				return v
			}()
		}
	}

	return ret
}

func main() {
	p := parseInput(utils.ReadInput("input.txt"))

	fmt.Printf("A: %v\n", func() int {
		var ret int
		for _, pp := range p {
			ret += history(takeLast(extrapolate(pp)))
		}
		return ret
	}())
}
