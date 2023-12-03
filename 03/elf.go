package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"regexp"
	"strconv"
)

type point struct {
	x int
	y int
}

var numberExp = regexp.MustCompile("([0-9]+)")

type number struct {
	value  int
	points []point
}

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Printf("A: %v\n", sumContactPoints(findNumbers(input), findSymbols(input, isSymbol)))
}

func isSymbol(v string) bool {
	_, err := strconv.Atoi(v)
	return err != nil && v != "."
}

func isStar(v string) bool {
	return v == "*"
}

func sumContactPoints(numbers []number, symbols []point) int {
	var sum int
	for _, num := range numbers {
		if neighbour(num.points, symbols) {
			sum += num.value
		}
	}
	return sum
}

func neighbour(points []point, symbols []point) bool {
	for _, num := range points {
		for _, symbol := range symbols {
			if (num.x+1 == symbol.x && num.y == symbol.y) || (num.x-1 == symbol.x && num.y == symbol.y) ||
				(num.x == symbol.x && num.y+1 == symbol.y) || (num.x == symbol.x && num.y-1 == symbol.y) ||
				(num.x-1 == symbol.x && num.y-1 == symbol.y) || (num.x+1 == symbol.x && num.y+1 == symbol.y) ||
				(num.x+1 == symbol.x && num.y-1 == symbol.y) || (num.x-1 == symbol.x && num.y+1 == symbol.y) {
				return true
			}
		}
	}

	return false
}

func findNumbers(s []string) []number {
	numbers := make([]number, 0)

	for y, ss := range s {
		idxs := numberExp.FindAllStringIndex(ss, -1)
		for _, idx := range idxs {
			var value string
			points := make([]point, 0)
			for x := idx[0]; x < idx[1]; x++ {
				value += string(ss[x])
				points = append(points, point{x: x, y: y})
			}
			numbers = append(numbers, number{
				value: func() int {
					v, err := strconv.Atoi(value)
					if err != nil {
						return 0
					}
					return v
				}(),
				points: points,
			})
		}
	}

	return numbers
}

func findSymbols(s []string, f func(v string) bool) []point {
	pts := make([]point, 0)

	for y, ss := range s {
		for x, sss := range ss {
			v := string(sss)
			if f(v) {
				pts = append(pts, point{
					x: x,
					y: y,
				})
			}
		}
	}

	return pts
}
