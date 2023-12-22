package day05

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

func getNum(inp int, data [][]int) int {
	for _, set := range data {
		if inp >= set[1] && inp < set[1]+set[2] {
			return set[0] + (inp - set[1])
		}
	}
	return inp
}

func Soil() int {
	sc, file := utils.Read("./day05/input.txt")
	// sc, file := utils.Read("./day05/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	var seeds []int
	var soil, fert, wat, light, temp, hum, loc [][]int
	var line, marker string
	var tmpInt, dest, src, n int
	var tmpStrArr []string

	for sc.Scan() {
		line = sc.Text()
		if len(line) < 1 {
			continue
		}
		if line[:5] == "seeds" {
			marker = "seeds"
			tmpStrArr = strings.Split(strings.Split(line, ": ")[1], " ")
			for _, v := range tmpStrArr {
				tmpInt, _ = strconv.Atoi(v)
				seeds = append(seeds, tmpInt)
			}
			continue
		} else if line == "seed-to-soil map:" {
			marker = "soil"
			continue
		} else if line == "soil-to-fertilizer map:" {
			marker = "fert"
			continue
		} else if line == "fertilizer-to-water map:" {
			marker = "wat"
			continue
		} else if line == "water-to-light map:" {
			marker = "light"
			continue
		} else if line == "light-to-temperature map:" {
			marker = "temp"
			continue
		} else if line == "temperature-to-humidity map:" {
			marker = "hum"
			continue
		} else if line == "humidity-to-location map:" {
			marker = "loc"
			continue
		}

		// Number line
		tmpStrArr = strings.Split(line, " ")
		dest, _ = strconv.Atoi(tmpStrArr[0])
		src, _ = strconv.Atoi(tmpStrArr[1])
		n, _ = strconv.Atoi(tmpStrArr[2])

		if marker == "soil" {
			soil = append(soil, []int{dest, src, n})
		} else if marker == "fert" {
			fert = append(fert, []int{dest, src, n})
		} else if marker == "wat" {
			wat = append(wat, []int{dest, src, n})
		} else if marker == "light" {
			light = append(light, []int{dest, src, n})
		} else if marker == "temp" {
			temp = append(temp, []int{dest, src, n})
		} else if marker == "hum" {
			hum = append(hum, []int{dest, src, n})
		} else if marker == "loc" {
			loc = append(loc, []int{dest, src, n})
		}
	}

	minLoc := -1

	for _, seed := range seeds {
		tmpInt = seed
		tmpInt = getNum(tmpInt, soil)
		tmpInt = getNum(tmpInt, fert)
		tmpInt = getNum(tmpInt, wat)
		tmpInt = getNum(tmpInt, light)
		tmpInt = getNum(tmpInt, temp)
		tmpInt = getNum(tmpInt, hum)
		tmpInt = getNum(tmpInt, loc)

		if minLoc == -1 || tmpInt < minLoc {
			minLoc = tmpInt
		}
	}

	return minLoc
}

func getNumRange(seeds []Tuple, data [][]int) []Tuple {
	start, end := 0, 0
	newSeeds := []Tuple{}
	for len(seeds) > 0 {
		start, end = seeds[0].a, seeds[0].b
		seeds = seeds[1:]
		found := false
		for _, v := range data {
			dest, src, length := v[0], v[1], v[2]
			overlapStart := math.Max(float64(start), float64(src))
			overlapEnd := math.Min(float64(end), float64(src+length))
			if overlapStart < overlapEnd {
				newSeeds = append(newSeeds, Tuple{int(overlapStart) - src + dest, int(overlapEnd) - src + dest})
				if int(overlapStart) > start {
					seeds = append(seeds, Tuple{start, int(overlapStart)})
				}
				if end > int(overlapEnd) {
					seeds = append(seeds, Tuple{int(overlapEnd), end})
				}
				found = true
				break
			}
		}
		if !found {
			newSeeds = append(newSeeds, Tuple{start, end})
		}
	}
	return newSeeds
}

type Tuple struct {
	a int
	b int
}

func SoilAnnoy() int {
	sc, file := utils.Read("./day05/input.txt")
	// sc, file := utils.Read("./day05/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	var seeds []int
	var soil, fert, wat, light, temp, hum, loc [][]int
	var line, marker string
	var tmpInt, dest, src, n int
	var tmpStrArr []string

	for sc.Scan() {
		line = sc.Text()
		if len(line) < 1 {
			continue
		}
		if line[:5] == "seeds" {
			marker = "seeds"
			tmpStrArr = strings.Split(strings.Split(line, ": ")[1], " ")
			for _, v := range tmpStrArr {
				tmpInt, _ = strconv.Atoi(v)
				seeds = append(seeds, tmpInt)
			}
			continue
		} else if line == "seed-to-soil map:" {
			marker = "soil"
			continue
		} else if line == "soil-to-fertilizer map:" {
			marker = "fert"
			continue
		} else if line == "fertilizer-to-water map:" {
			marker = "wat"
			continue
		} else if line == "water-to-light map:" {
			marker = "light"
			continue
		} else if line == "light-to-temperature map:" {
			marker = "temp"
			continue
		} else if line == "temperature-to-humidity map:" {
			marker = "hum"
			continue
		} else if line == "humidity-to-location map:" {
			marker = "loc"
			continue
		}

		// Number line
		tmpStrArr = strings.Split(line, " ")
		dest, _ = strconv.Atoi(tmpStrArr[0])
		src, _ = strconv.Atoi(tmpStrArr[1])
		n, _ = strconv.Atoi(tmpStrArr[2])

		if marker == "soil" {
			soil = append(soil, []int{dest, src, n})
		} else if marker == "fert" {
			fert = append(fert, []int{dest, src, n})
		} else if marker == "wat" {
			wat = append(wat, []int{dest, src, n})
		} else if marker == "light" {
			light = append(light, []int{dest, src, n})
		} else if marker == "temp" {
			temp = append(temp, []int{dest, src, n})
		} else if marker == "hum" {
			hum = append(hum, []int{dest, src, n})
		} else if marker == "loc" {
			loc = append(loc, []int{dest, src, n})
		}
	}

	var seedRanges []Tuple

	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, Tuple{seeds[i], seeds[i] + seeds[i+1]})
	}
	seedRanges = getNumRange(seedRanges, soil)
	seedRanges = getNumRange(seedRanges, fert)
	seedRanges = getNumRange(seedRanges, wat)
	seedRanges = getNumRange(seedRanges, light)
	seedRanges = getNumRange(seedRanges, temp)
	seedRanges = getNumRange(seedRanges, hum)
	seedRanges = getNumRange(seedRanges, loc)

	// fmt.Println(seedRanges)

	sortFunc := func(a, b Tuple) int {
		return cmp.Compare(a.a, b.a)
	}
	slices.SortFunc(seedRanges, sortFunc)
	// fmt.Println(seedRanges)

	return seedRanges[0].a
}
