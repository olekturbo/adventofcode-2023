package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := utils.ReadInput("input.txt")
	fmt.Printf("A: %v\n", calculateCalibration(s, false))
	fmt.Printf("B: %v\n", calculateCalibration(s, true))
}

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func calibrate(slice []string, extended bool) [][]int {
	r := make([][]int, len(slice))

	for i, s := range slice {
		r[i] = make([]int, 2)
		m := make(map[int]int, 0)

		for k, v := range digits {
			if extended {
				if idx := strings.Index(s, k); idx >= 0 {
					m[idx] = v
				}

				if idx := strings.LastIndex(s, k); idx >= 0 {
					m[idx] = v
				}
			}

			if idx := strings.Index(s, strconv.Itoa(v)); idx >= 0 {
				m[idx] = v
			}

			if idx := strings.LastIndex(s, strconv.Itoa(v)); idx >= 0 {
				m[idx] = v
			}
		}

		idxs := make([]int, 0, len(m))
		for k := range m {
			idxs = append(idxs, k)
		}
		sort.Ints(idxs)

		r[i][0] = m[idxs[0]]
		r[i][1] = m[idxs[len(idxs)-1]]
	}

	return r
}

func calculateCalibration(slice []string, extended bool) int {
	var sum int

	for _, v := range calibrate(slice, extended) {
		sum += v[0]*10 + v[1]
	}

	return sum
}
