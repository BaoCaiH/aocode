package day08

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type grid struct {
	rows [][]int
	cols [][]int
}

func (g *grid) makeCols(r, c int) {
	// Make cols
	g.cols = make([][]int, c)
	for i := 0; i < c; i++ {
		g.cols[i] = make([]int, r)
	}
}

func (g *grid) addElement(e, r, c int) {
	g.cols[c][r] = e
}

func (g grid) isVisible(e, r, c int) bool {
	val := e
	// Check left
	visibility := true
	for _, t := range g.rows[r][:c] {
		visibility = visibility && t < val
	}
	if visibility {
		return visibility
	}
	// Check right
	visibility = true
	for _, t := range g.rows[r][c+1:] {
		visibility = visibility && t < val
	}
	if visibility {
		return visibility
	}
	// Check top
	visibility = true
	for _, t := range g.cols[c][:r] {
		visibility = visibility && t < val
	}
	if visibility {
		return visibility
	}
	// Check bottom
	visibility = true
	for _, t := range g.cols[c][r+1:] {
		visibility = visibility && t < val
	}
	return visibility
}

func (g *grid) scenicScore(e, r, c int) int {
	left, right, top, bottom := 0, 0, 0, 0
	for _, t := range g.rows[r][:c] {
		if t < e {
			left++
		} else {
			left = 1
		}
	}
	for _, t := range g.rows[r][c+1:] {
		right++
		if t >= e {
			break
		}
	}
	for _, t := range g.cols[c][:r] {
		if t < e {
			top++
		} else {
			top = 1
		}
	}
	for _, t := range g.cols[c][r+1:] {
		bottom++
		if t >= e {
			break
		}
	}
	return left * right * top * bottom
}

func loadTrees() grid {
	file, err := os.Open("./day08/input.txt")
	if err != nil {
		fmt.Println(err)
		return grid{}
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	g := grid{}
	for sc.Scan() {
		line := sc.Text()
		row := []int{}
		for _, e := range line {
			val, _ := strconv.Atoi(string(e))
			row = append(row, val)
		}
		g.rows = append(g.rows, row)
	}
	g.makeCols(len(g.rows), len(g.rows[0]))
	for r, row := range g.rows {
		for c, val := range row {
			g.addElement(val, r, c)
		}
	}

	return g
}

func TreetopVisible() int {
	g := loadTrees()
	counts := 0
	for r, row := range g.rows {
		for c, val := range row {
			if g.isVisible(val, r, c) {
				counts++
			}
		}
	}

	return counts
}

func TreetopScenic() int {
	g := loadTrees()
	maxScore := 0
	for r, row := range g.rows {
		for c, val := range row {
			newScore := g.scenicScore(val, r, c)
			if newScore > maxScore {
				maxScore = newScore
			}
		}
	}

	return maxScore
}
