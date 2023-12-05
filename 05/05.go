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

	var ks []k

	for i, s := range input {

		// seeds
		if i == 0 {
			s = strings.Replace(s, "seeds: ", "", 1)
			sp := strings.Split(s, " ")
			for _, seed := range sp {
				sseed, _ := strconv.Atoi(seed)
				ks = append(ks, k{seed: sseed})
			}
		}

		// seed-to-soil
		if i == 1 {
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
					for x, kk := range ks {
						if kk.seed == ssource+z {
							ks[x].soil = ddest + z
						}
					}
				}
			}

			for x, kk := range ks {
				if kk.soil == 0 {
					ks[x].soil = kk.seed
				}
			}

		}

		// soil-to-fertilizer
		if i == 2 {
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
					for x, kk := range ks {
						if kk.soil == ssource+z {
							ks[x].fertilizer = ddest + z
						}
					}
				}
			}

			for x, kk := range ks {
				if kk.fertilizer == 0 {
					ks[x].fertilizer = kk.soil
				}
			}

		}

		// fertilizer-to-water
		if i == 3 {
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
					for x, kk := range ks {
						if kk.fertilizer == ssource+z {
							ks[x].water = ddest + z
						}
					}
				}
			}

			for x, kk := range ks {
				if kk.water == 0 {
					ks[x].water = kk.fertilizer
				}
			}

		}

		// water-to-light
		if i == 4 {
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
					for x, kk := range ks {
						if kk.water == ssource+z {
							ks[x].light = ddest + z
						}
					}
				}
			}

			for x, kk := range ks {
				if kk.light == 0 {
					ks[x].light = kk.water
				}
			}

		}

		// light-to-temperature
		if i == 5 {
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
					for x, kk := range ks {
						if kk.light == ssource+z {
							ks[x].temperature = ddest + z
						}
					}
				}
			}

			for x, kk := range ks {
				if kk.temperature == 0 {
					ks[x].temperature = kk.light
				}
			}

		}

		// temperature-to-humidity
		if i == 6 {
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
					for x, kk := range ks {
						if kk.temperature == ssource+z {
							ks[x].humidity = ddest + z
						}
					}
				}
			}

			for x, kk := range ks {
				if kk.humidity == 0 {
					ks[x].humidity = kk.temperature
				}
			}

		}

		// humidity-to-location
		if i == 7 {
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
					for x, kk := range ks {
						if kk.humidity == ssource+z {
							ks[x].location = ddest + z
						}
					}
				}
			}

			for x, kk := range ks {
				if kk.location == 0 {
					ks[x].location = kk.humidity
				}
			}

		}

	}

	var locs []int
	for _, kk := range ks {
		locs = append(locs, kk.location)
	}

	fmt.Println(slices.Min(locs))

}
