package day18

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

type tuple struct {
	r, c int
}

func Dig() {
	sc, file := utils.Read("./day18/input.txt")
	// sc, file := utils.Read("./day18/input_1_example.txt")
	if sc == nil {
		return
	}
	defer file.Close()

	trenches := []tuple{{0, 0}}
	differentTrenches := []tuple{{0, 0}}
	differentMap := map[byte]string{'0': "R", '1': "D", '2': "L", '3': "U"}
	border := 0
	differentBorder := 0
	input := []string{}
	d := ""
	dd := ""
	n := 0
	dn := int64(0)
	c := ""

	for sc.Scan() {
		input = strings.Split(sc.Text(), " ")
		d = input[0]
		n, _ = strconv.Atoi(input[1])
		c = strings.Trim(input[2], "()")
		dd = differentMap[c[len(c)-1]]
		dn, _ = strconv.ParseInt(c[1:len(c)-1], 16, 64)
		p := trenches[len(trenches)-1]
		dp := differentTrenches[len(differentTrenches)-1]
		if d == "R" {
			trenches = append(trenches, tuple{p.r, p.c + n})
		} else if d == "L" {
			trenches = append(trenches, tuple{p.r, p.c - n})
		} else if d == "U" {
			trenches = append(trenches, tuple{p.r - n, p.c})
		} else if d == "D" {
			trenches = append(trenches, tuple{p.r + n, p.c})
		}

		if dd == "R" {
			differentTrenches = append(differentTrenches, tuple{dp.r, dp.c + int(dn)})
		} else if dd == "L" {
			differentTrenches = append(differentTrenches, tuple{dp.r, dp.c - int(dn)})
		} else if dd == "U" {
			differentTrenches = append(differentTrenches, tuple{dp.r - int(dn), dp.c})
		} else if dd == "D" {
			differentTrenches = append(differentTrenches, tuple{dp.r + int(dn), dp.c})
		}

		border += n
		differentBorder += int(dn)
	}

	total := 0
	for i := range trenches {
		if i == 0 {
			total += trenches[i].r * (trenches[i+1].c - trenches[0].c)
		} else if i+1 == len(trenches) {
			total += trenches[i].r * (trenches[0].c - trenches[i-1].c)
		} else {
			total += trenches[i].r * (trenches[i+1].c - trenches[i-1].c)
		}
	}
	areaInside := (int(math.Abs(float64(total))) / 2) - (border / 2) + 1

	fmt.Printf("\tPart 1: %d\n", areaInside+border)

	total = 0
	for i := range differentTrenches {
		if i == 0 {
			total += differentTrenches[i].r * (differentTrenches[i+1].c - differentTrenches[0].c)
		} else if i+1 == len(differentTrenches) {
			total += differentTrenches[i].r * (differentTrenches[0].c - differentTrenches[i-1].c)
		} else {
			total += differentTrenches[i].r * (differentTrenches[i+1].c - differentTrenches[i-1].c)
		}
	}
	areaInside = (int(math.Abs(float64(total))) / 2) - (differentBorder / 2) + 1

	fmt.Printf("\tPart 2: %d\n", areaInside+differentBorder)
}
