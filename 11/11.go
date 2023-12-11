package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"math"
	"strings"
)

func parse(s []string) [][]int {
	y0 := make(map[int]bool)
	for i := range s {
		var ys string
		for x := range s[i] {
			ys += string(s[x][i])
		}
		if !strings.Contains(ys, "#") {
			y0[i] = true
		}
	}

	ss := make([]string, 0)
	for _, v := range s {
		if !strings.Contains(v, "#") {
			for i := 0; i < 1000; i++ {
				ss = append(ss, v)
			}
		} else {
			ss = append(ss, v)
		}
	}

	for i, v := range ss {
		var aa string
		for j, vv := range v {
			if _, ok := y0[j]; ok {
				for i := 0; i < 1000; i++ {
					aa += string(vv)
				}
			} else {
				aa += string(vv)
			}
		}
		ss[i] = aa
	}

	ret := make([][]int, len(ss))
	d := 0
	for k, v := range ss {
		ret[k] = make([]int, len(v))
		for kk, vv := range v {
			if string(vv) == "#" {
				d++
				ret[k][kk] = d
			} else {
				ret[k][kk] = 0
			}
		}
	}

	return ret
}

type point struct {
	x, y, v int
}

func points(graph [][]int) []point {
	ret := make([]point, 0)
	for y, v := range graph {
		for x, vv := range v {
			if vv != 0 {
				ret = append(ret, point{
					y: y,
					x: x,
					v: vv,
				})
			}
		}
	}
	return ret
}

func distances(points []point) []float64 {
	ret := make([]float64, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			ret = append(ret, math.Abs(float64(points[i].x-points[j].x))+math.Abs(float64(points[i].y-points[j].y)))
		}
	}

	return ret
}

func main() {
	in := utils.ReadInput("test.txt")
	p := parse(in)
	d := distances(points(p))
	var s int
	for _, dd := range d {
		idd := int(dd)
		r := idd % 1000
		z := (idd-r)*1000 + r
		s += z
	}
	fmt.Println(s)
}
