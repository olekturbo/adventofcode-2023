package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := utils.ReadInput("input2.txt")
	t := s[0]
	d := s[1]
	t = strings.Replace(t, "Time:", "", 1)
	d = strings.Replace(d, "Distance:", "", 1)
	tt := strings.Split(t, " ")
	dd := strings.Split(d, " ")

	var ways []int
	for i := 0; i < len(tt); i++ {
		var currWays int
		distance, _ := strconv.Atoi(dd[i])
		time, _ := strconv.Atoi(tt[i])

		for j := 0; j <= time; j++ {
			v := (time - j) * j
			if v > distance {
				currWays++
			}
		}
		ways = append(ways, currWays)
	}

	ret := 1
	for _, v := range ways {
		ret *= v
	}

	fmt.Println(ret)

}
