package day11

import (
	"math"
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

func isEmptyCol(i int, g []string) bool {
	for _, line := range g {
		if line[i] == '#' {
			return false
		}
	}
	return true
}

func isEmptyRow(r string) bool {
	for _, c := range r {
		if c == '#' {
			return false
		}
	}
	return true
}

type Coor struct {
	r, c int
}

func Cosmic() int {
	sc, file := utils.Read("./day11/input.txt")
	// sc, file := utils.Read("./day11/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	galaxies := []string{}
	line := ""
	for sc.Scan() {
		line = sc.Text()
		galaxies = append(galaxies, line)
	}

	// colExpands := []int{}
	// for i := range galaxies[0] {
	// 	if isEmpty(i, galaxies) {
	// 		colExpands = append(colExpands, i)
	// 	}
	// }
	i := 0
	for i < len(galaxies[0]) {
		if !isEmptyCol(i, galaxies) {
			i += 1
			continue
		}
		for k, line := range galaxies {
			galaxies[k] = line[:i] + "." + line[i:]
		}
		i += 2
	}
	i = 0
	for i < len(galaxies) {
		if !isEmptyRow(galaxies[i]) {
			i += 1
			continue
		}
		galaxies = append(append(galaxies[:i], galaxies[i]), galaxies[i:]...)
		i += 2
	}

	tinyDots := []Coor{}
	for r, line := range galaxies {
		for c, cha := range line {
			if cha == '#' {
				tinyDots = append(tinyDots, Coor{r, c})
			}
		}
		// fmt.Println(line)
	}

	total := 0
	for i, dot := range tinyDots {
		for _, anotherDot := range tinyDots[i:] {
			total += int(math.Abs(float64(dot.r)-float64(anotherDot.r))) + int(math.Abs(float64(dot.c)-float64(anotherDot.c)))
		}
	}

	return total
}

func nEmpty(left, right int, expands []int) int {
	from, to := 0, 0
	if left > right {
		from = right
		to = left
	} else {
		from = left
		to = right
	}
	cnt := 0
	for _, i := range expands {
		if i >= to {
			break
		}
		if i <= from {
			continue
		}
		cnt += 1
	}
	return cnt
}

func CosmicEnergy() int {
	sc, file := utils.Read("./day11/input.txt")
	// sc, file := utils.Read("./day11/input_1_example.txt")
	rate := 1000000
	if sc == nil {
		return 0
	}
	defer file.Close()

	galaxies := []string{}
	line := ""
	for sc.Scan() {
		line = sc.Text()
		galaxies = append(galaxies, line)
	}

	// colExpands := []int{}
	// for i := range galaxies[0] {
	// 	if isEmpty(i, galaxies) {
	// 		colExpands = append(colExpands, i)
	// 	}
	// }
	colExpansion := []int{}
	i := 0
	for i < len(galaxies[0]) {
		if !isEmptyCol(i, galaxies) {
			i += 1
			continue
		}
		for k, line := range galaxies {
			galaxies[k] = line[:i] + "0" + line[i+1:]
		}
		colExpansion = append(colExpansion, i)
		i += 1
	}
	rowExpansion := []int{}
	i = 0
	for i < len(galaxies) {
		if !isEmptyRow(galaxies[i]) {
			i += 1
			continue
		}
		galaxies[i] = strings.ReplaceAll(galaxies[i], ".", "0")
		rowExpansion = append(rowExpansion, i)
		i += 1
	}

	// fmt.Println(rowExpansion)
	// fmt.Println(colExpansion)

	tinyDots := []Coor{}
	for r, line := range galaxies {
		for c, cha := range line {
			if cha == '#' {
				tinyDots = append(tinyDots, Coor{r, c})
			}
		}
		// fmt.Println(line)
	}

	total, emptyRow, emptyCol := 0, 0, 0
	for i, dot := range tinyDots {
		for _, anotherDot := range tinyDots[i:] {
			emptyRow = nEmpty(dot.r, anotherDot.r, rowExpansion)
			emptyCol = nEmpty(dot.c, anotherDot.c, colExpansion)
			total += int(math.Abs(float64(dot.r)-float64(anotherDot.r))) - emptyRow + emptyRow*rate + int(math.Abs(float64(dot.c)-float64(anotherDot.c))) - emptyCol + emptyCol*rate
		}
	}

	return total
}
