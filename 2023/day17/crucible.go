package day17

import (
	"container/heap"
	"fmt"
	"slices"
	"strconv"

	"github.com/BaoCaiH/aocode/2023/utils"
)

type coor struct {
	l, r, c int
	d       string
}

type lossless struct {
	r, c int
	d    string
}

type Heap []coor

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].l < h[j].l
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(coor))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Crucible() {
	sc, file := utils.Read("./day17/input.txt")
	// sc, file := utils.Read("./day17/input_1_example.txt")
	if sc == nil {
		return
	}
	defer file.Close()

	grid := [][]int{}
	line := ""
	for sc.Scan() {
		line = sc.Text()
		tmp := []int{}
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			tmp = append(tmp, n)
		}
		grid = append(grid, tmp)
	}

	width := len(grid[0])
	height := len(grid)
	offsets := map[string][]int{">": {0, 1}, "<": {0, -1}, "^": {-1, 0}, "v": {1, 0}}
	opposites := map[string]string{">": "<", "<": ">", "^": "v", "v": "^"}

	traverse := Heap{}
	heap.Push(&traverse, coor{0, 0, 0, "."})
	passed := []lossless{}
	for len(traverse) > 0 {
		curr := heap.Pop(&traverse).(coor)
		l, r, c, d := curr.l, curr.r, curr.c, curr.d

		if r == height-1 && c == width-1 {
			fmt.Printf("\tPart 1: %d\n", l)
			break
		}

		marker := lossless{r, c, d}
		if slices.Contains(passed, marker) {
			continue
		}
		passed = append(passed, marker)

		for dir, offset := range offsets {
			if dir == opposites[string(d[0])] {
				continue
			}
			if dir == string(d[0]) && len(d) >= 3 {
				continue
			}
			newR := r + offset[0]
			newC := c + offset[1]
			if newR < 0 || newR >= height || newC < 0 || newC >= width {
				continue
			}
			newD := ""
			if dir != string(d[0]) {
				newD = dir
			} else {
				newD = d + dir
			}
			heap.Push(&traverse, coor{l + grid[newR][newC], newR, newC, newD})
		}
	}

	traverse = Heap{}
	heap.Push(&traverse, coor{0, 0, 0, "."})
	passed = []lossless{}
	for len(traverse) > 0 {
		curr := heap.Pop(&traverse).(coor)
		l, r, c, d := curr.l, curr.r, curr.c, curr.d

		if r == height-1 && c == width-1 {
			fmt.Printf("\tPart 2: %d\n", l)
			break
		}

		marker := lossless{r, c, d}
		if slices.Contains(passed, marker) {
			continue
		}
		passed = append(passed, marker)

		for dir, offset := range offsets {
			if dir == opposites[string(d[0])] {
				continue
			}
			if dir == string(d[0]) && len(d) >= 10 {
				continue
			}
			if dir != string(d[0]) && len(d) < 4 && d != "." {
				continue
			}
			newR := r + offset[0]
			newC := c + offset[1]
			if newR < 0 || newR >= height || newC < 0 || newC >= width {
				continue
			}
			newD := ""
			if dir != string(d[0]) {
				newD = dir
			} else {
				newD = d + dir
			}
			heap.Push(&traverse, coor{l + grid[newR][newC], newR, newC, newD})
		}
	}
}
