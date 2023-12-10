package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"math"
	"time"
)

type point struct {
	value   string
	visited bool
	x       int
	y       int
	d       int
}

func (p *point) direction(s string) []int {
	switch s {
	case "NORTH":
		return []int{p.x, p.y - 1}
	case "SOUTH":
		return []int{p.x, p.y + 1}
	case "EAST":
		return []int{p.x + 1, p.y}
	case "WEST":
		return []int{p.x - 1, p.y}
	default:
		panic("not possible")
	}
}

func (p *point) neighbours() [][]int {
	switch p.value {
	case "S":
		return [][]int{
			p.direction("NORTH"),
			p.direction("SOUTH"),
			p.direction("EAST"),
			p.direction("WEST"),
		}
	case "|":
		return [][]int{
			p.direction("NORTH"),
			p.direction("SOUTH"),
		}
	case "-":
		return [][]int{
			p.direction("EAST"),
			p.direction("WEST"),
		}
	case "L":
		return [][]int{
			p.direction("NORTH"),
			p.direction("EAST"),
		}
	case "J":
		return [][]int{
			p.direction("NORTH"),
			p.direction("WEST"),
		}
	case "7":
		return [][]int{
			p.direction("SOUTH"),
			p.direction("WEST"),
		}
	case "F":
		return [][]int{
			p.direction("SOUTH"),
			p.direction("EAST"),
		}
	case ".":
		return nil

	default:
		panic("not possible")
	}
}

func (p *point) markAsVisited(d int) {
	p.d = d
	p.visited = true
}

type points []*point

func (p *points) findByPos(x, y int) *point {
	for _, v := range *p {
		if v.x == x && v.y == y && v.value != "." {
			return v
		}
	}
	return nil
}

func (p *points) findByValue(value string) *point {
	for _, v := range *p {
		if v.value == value {
			return v
		}
	}
	return nil
}

func parse(s []string) points {
	ret := make(points, 0)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			ret = append(ret, &point{
				value:   string(s[i][j]),
				visited: false,
				x:       j,
				y:       i,
			})
		}
	}

	return ret
}

func main() {
	pts := parse(utils.ReadInput("test.txt"))

	aStart := time.Now()
	start := pts.findByValue("S")
	d := 0
	start.markAsVisited(d)

	f(d, pts, start)

	max := func() float64 {
		var m float64
		for _, v := range pts {
			if float64(v.d) > m {
				m = float64(v.d)
			}
		}
		return m
	}()

	fmt.Printf("%v: %v\n", math.Ceil(max/2), time.Since(aStart))
}

func f(d int, pts points, start *point) {
	d++
	for _, v := range start.neighbours() {
		found := pts.findByPos(v[0], v[1])
		if found != nil && !found.visited {
			found.markAsVisited(d)
			f(d, pts, found)
		}
	}
}
