package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type k struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(f), "\n\n")

	var seeds []int
	var soils []int
	var fertilizer []int
	var waters []int
	var lights []int
	var temperatures []int
	var humidities []int
	var locations []int

	for i, s := range input {

		// seeds
		if i == 0 {
			s = strings.Replace(s, "seeds: ", "", 1)
			sp := strings.Split(s, " ")
			for x, seed := range sp {
				if x%2 == 1 {
					continue
				}
				sseed, _ := strconv.Atoi(seed)
				dseed, _ := strconv.Atoi(sp[x+1])
				for b := 0; b < dseed; b++ {
					seeds = append(seeds, b+sseed)
				}
			}
		}

		// seed-to-soil
		if i == 1 {
			soils = make([]int, len(seeds))
			for j, ss := range strings.Split(s, "\n") {
				if j == 0 {
					continue
				}
				sp := strings.Split(ss, " ")
				dest, source, l := sp[0], sp[1], sp[2]
				ddest, _ := strconv.Atoi(dest)
				ssource, _ := strconv.Atoi(source)
				ll, _ := strconv.Atoi(l)

				for z := 0; z < ll; z++ {
					for x, kk := range seeds {
						if kk == ssource+z {
							soils[x] = ddest + z
						}
					}
				}
			}

			for x, kk := range soils {
				if kk == 0 {
					soils[x] = seeds[x]
				}
			}

		}

		// soil-to-fertilizer
		if i == 2 {
			fertilizer = make([]int, len(seeds))
			for j, ss := range strings.Split(s, "\n") {
				if j == 0 {
					continue
				}
				sp := strings.Split(ss, " ")
				dest, source, l := sp[0], sp[1], sp[2]
				ddest, _ := strconv.Atoi(dest)
				ssource, _ := strconv.Atoi(source)
				ll, _ := strconv.Atoi(l)

				for z := 0; z < ll; z++ {
					for x, kk := range soils {
						if kk == ssource+z {
							fertilizer[x] = ddest + z
						}
					}
				}
			}

			for x, kk := range fertilizer {
				if kk == 0 {
					fertilizer[x] = soils[x]
				}
			}

		}

		// fertilizer-to-water
		if i == 3 {
			waters = make([]int, len(seeds))
			for j, ss := range strings.Split(s, "\n") {
				if j == 0 {
					continue
				}
				sp := strings.Split(ss, " ")
				dest, source, l := sp[0], sp[1], sp[2]
				ddest, _ := strconv.Atoi(dest)
				ssource, _ := strconv.Atoi(source)
				ll, _ := strconv.Atoi(l)

				for z := 0; z < ll; z++ {
					for x, kk := range fertilizer {
						if kk == ssource+z {
							waters[x] = ddest + z
						}
					}
				}
			}

			for x, kk := range waters {
				if kk == 0 {
					waters[x] = fertilizer[x]
				}
			}

		}

		// water-to-light
		if i == 4 {
			lights = make([]int, len(seeds))
			for j, ss := range strings.Split(s, "\n") {
				if j == 0 {
					continue
				}
				sp := strings.Split(ss, " ")
				dest, source, l := sp[0], sp[1], sp[2]
				ddest, _ := strconv.Atoi(dest)
				ssource, _ := strconv.Atoi(source)
				ll, _ := strconv.Atoi(l)

				for z := 0; z < ll; z++ {
					for x, kk := range waters {
						if kk == ssource+z {
							lights[x] = ddest + z
						}
					}
				}
			}

			for x, kk := range lights {
				if kk == 0 {
					lights[x] = waters[x]
				}
			}

		}

		// light-to-temperature
		if i == 5 {
			temperatures = make([]int, len(seeds))
			for j, ss := range strings.Split(s, "\n") {
				if j == 0 {
					continue
				}
				sp := strings.Split(ss, " ")
				dest, source, l := sp[0], sp[1], sp[2]
				ddest, _ := strconv.Atoi(dest)
				ssource, _ := strconv.Atoi(source)
				ll, _ := strconv.Atoi(l)

				for z := 0; z < ll; z++ {
					for x, kk := range lights {
						if kk == ssource+z {
							temperatures[x] = ddest + z
						}
					}
				}
			}

			for x, kk := range temperatures {
				if kk == 0 {
					temperatures[x] = lights[x]
				}
			}

		}

		// temperature-to-humidity
		if i == 6 {
			humidities = make([]int, len(seeds))
			for j, ss := range strings.Split(s, "\n") {
				if j == 0 {
					continue
				}
				sp := strings.Split(ss, " ")
				dest, source, l := sp[0], sp[1], sp[2]
				ddest, _ := strconv.Atoi(dest)
				ssource, _ := strconv.Atoi(source)
				ll, _ := strconv.Atoi(l)

				for z := 0; z < ll; z++ {
					for x, kk := range temperatures {
						if kk == ssource+z {
							humidities[x] = ddest + z
						}
					}
				}
			}

			for x, kk := range humidities {
				if kk == 0 {
					humidities[x] = temperatures[x]
				}
			}

		}

		// humidity-to-location
		if i == 7 {
			locations = make([]int, len(seeds))
			for j, ss := range strings.Split(s, "\n") {
				if j == 0 {
					continue
				}
				sp := strings.Split(ss, " ")
				dest, source, l := sp[0], sp[1], sp[2]
				ddest, _ := strconv.Atoi(dest)
				ssource, _ := strconv.Atoi(source)
				ll, _ := strconv.Atoi(l)

				for z := 0; z < ll; z++ {
					for x, kk := range humidities {
						if kk == ssource+z {
							locations[x] = ddest + z
						}
					}
				}
			}

			for x, kk := range locations {
				if kk == 0 {
					locations[x] = humidities[x]
				}
			}

		}

	}

	var locs []int
	for _, kk := range locations {
		locs = append(locs, kk)
	}

	fmt.Println(locations)

	fmt.Println(slices.Min(locs))

}
