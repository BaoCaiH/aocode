package day16

import (
	"fmt"
	"slices"

	"github.com/BaoCaiH/aocode/2023/utils"
)

type Coor struct {
	r, c int
	d    string
}

func next(c Coor, grid []string) []Coor {
	mirror := grid[c.r][c.c]
	nexts := []Coor{}
	if mirror == '\\' {
		if c.d == ">" && c.r < len(grid)-1 {
			nexts = append(nexts, Coor{c.r + 1, c.c, "v"})
		} else if c.d == "<" && c.r > 0 {
			nexts = append(nexts, Coor{c.r - 1, c.c, "^"})
		} else if c.d == "^" && c.c > 0 {
			nexts = append(nexts, Coor{c.r, c.c - 1, "<"})
		} else if c.d == "v" && c.c < len(grid[0])-1 {
			nexts = append(nexts, Coor{c.r, c.c + 1, ">"})
		}
	} else if mirror == '/' {
		if c.d == ">" && c.r > 0 {
			nexts = append(nexts, Coor{c.r - 1, c.c, "^"})
		} else if c.d == "<" && c.r < len(grid)-1 {
			nexts = append(nexts, Coor{c.r + 1, c.c, "v"})
		} else if c.d == "^" && c.c < len(grid[0])-1 {
			nexts = append(nexts, Coor{c.r, c.c + 1, ">"})
		} else if c.d == "v" && c.c > 0 {
			nexts = append(nexts, Coor{c.r, c.c - 1, "<"})
		}
	} else if mirror == '-' {
		if c.d == ">" && c.c < len(grid[0])-1 {
			nexts = append(nexts, Coor{c.r, c.c + 1, ">"})
		} else if c.d == "<" && c.c > 0 {
			nexts = append(nexts, Coor{c.r, c.c - 1, "<"})
		} else if c.d == "^" || c.d == "v" {
			if c.c < len(grid[0])-1 {
				nexts = append(nexts, Coor{c.r, c.c + 1, ">"})
			}
			if c.c > 0 {
				nexts = append(nexts, Coor{c.r, c.c - 1, "<"})
			}
		}
	} else if mirror == '|' {
		if c.d == ">" || c.d == "<" {
			if c.r < len(grid)-1 {
				nexts = append(nexts, Coor{c.r + 1, c.c, "v"})
			}
			if c.r > 0 {
				nexts = append(nexts, Coor{c.r - 1, c.c, "^"})
			}
		} else if c.d == "^" && c.r > 0 {
			nexts = append(nexts, Coor{c.r - 1, c.c, "^"})

		} else if c.d == "v" && c.r < len(grid)-1 {
			nexts = append(nexts, Coor{c.r + 1, c.c, "v"})
		}
	} else if mirror == '.' {
		if c.d == ">" && c.c < len(grid[0])-1 {
			nexts = append(nexts, Coor{c.r, c.c + 1, ">"})
		} else if c.d == "<" && c.c > 0 {
			nexts = append(nexts, Coor{c.r, c.c - 1, "<"})
		} else if c.d == "^" && c.r > 0 {
			nexts = append(nexts, Coor{c.r - 1, c.c, "^"})
		} else if c.d == "v" && c.r < len(grid)-1 {
			nexts = append(nexts, Coor{c.r + 1, c.c, "v"})
		}
	}
	return nexts
}

func spark(start Coor, grid []string) int {
	remaining := []Coor{start}
	tmp := []Coor{}
	energising := []Coor{}
	passed := []Coor{}
	for len(remaining) > 0 {
		for _, c := range remaining {
			if !slices.Contains(energising, Coor{c.r, c.c, ""}) {
				energising = append(energising, Coor{c.r, c.c, ""})
			}
			if slices.Contains(passed, c) {
				continue
			}
			passed = append(passed, c)
			tmp = append(tmp, next(c, grid)...)
		}
		remaining = tmp
		tmp = []Coor{}
	}
	return len(energising)
}

func Energise() {
	sc, file := utils.Read("./day16/input.txt")
	// sc, file := utils.Read("./day16/input_1_example.txt")
	if sc == nil {
		return
	}
	defer file.Close()

	grid := []string{}

	for sc.Scan() {
		grid = append(grid, sc.Text())
	}

	fmt.Printf("\tPart 1: %d\n", spark(Coor{0, 0, ">"}, grid))

	// // Commented out because too slow
	// max := 0
	// n := 0
	// width := len(grid[0])
	// height := len(grid)
	// for i := range grid[0] {
	// 	n = spark(Coor{0, i, "v"}, grid)
	// 	if n > max {
	// 		max = n
	// 	}
	// 	n = spark(Coor{height - 1, i, "^"}, grid)
	// 	if n > max {
	// 		max = n
	// 	}
	// }
	// for i := range grid {
	// 	n = spark(Coor{i, 0, ">"}, grid)
	// 	if n > max {
	// 		max = n
	// 	}
	// 	n = spark(Coor{i, width - 1, "<"}, grid)
	// 	if n > max {
	// 		max = n
	// 	}
	// }
	//
	// fmt.Printf("\tPart 2: %d\n", max)
}
