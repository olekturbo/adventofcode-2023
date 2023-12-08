package main

import (
	"adventofcode-2023/utils"
	"fmt"
	"strings"
)

type collection struct {
	left  string
	right string
}

func main() {
	collections, ways := parse(utils.ReadInput("input.txt"))

	kla := steps(collections, ways, "KLA", "SNZ")
	aaa := steps(collections, ways, "AAA", "ZZZ")
	nda := steps(collections, ways, "NDA", "DMZ")
	qva := steps(collections, ways, "QVA", "GTZ")
	lba := steps(collections, ways, "LBA", "CJZ")
	nna := steps(collections, ways, "NNA", "BDZ")

	fmt.Printf("A: %v\n", aaa)
	fmt.Printf("B: %v\n", lcm(kla, lcm(aaa, lcm(nda, lcm(qva, lcm(lba, nna))))))
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func steps(collections map[string]collection, ways []string, start, end string) int {
	curr := start
	var s int
	for {
		for _, way := range ways {
			col := collections[curr]
			s++
			if way == "L" {
				curr = col.left
			}
			if way == "R" {
				curr = col.right
			}
			if curr == end {
				return s
			}
		}
	}
}

func parse(s []string) (map[string]collection, []string) {
	ways := make([]string, 0)
	for _, ss := range s[0] {
		ways = append(ways, string(ss))
	}

	collections := make(map[string]collection)
	for i := 2; i < len(s); i++ {
		sp := strings.Split(s[i], " = ")

		cols := strings.Replace(sp[1], "(", "", 1)
		cols = strings.Replace(cols, ")", "", 1)

		colsSp := strings.Split(cols, ", ")

		collections[sp[0]] = collection{left: colsSp[0], right: colsSp[1]}
	}

	return collections, ways
}
