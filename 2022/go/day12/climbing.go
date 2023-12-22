package day12

import (
	"fmt"
	"sort"

	"github.com/BaoCaiH/aocode/utils"
)

type coord struct {
	x, y int
}

func parseMap(path string) ([][]rune, coord, coord) {
	sc, file := utils.Read(path)
	if sc == nil {
		return [][]rune{}, coord{}, coord{}
	}
	defer file.Close()

	heightMap := make([][]rune, 0)
	var start, end coord
	for sc.Scan() {
		line := sc.Text()
		lat := []rune{}
		for i, s := range line {
			if s == 'S' {
				start = coord{len(heightMap), i}
				s = 'a'
			}
			if s == 'E' {
				end = coord{len(heightMap), i}
				s = 'z'
			}

			lat = append(lat, s)
		}
		heightMap = append(heightMap, lat)
	}
	return heightMap, start, end
}

func neighbours(c coord, w, h int) []coord {
	n := []coord{}
	row, col := c.x, c.y
	if row-1 >= 0 {
		n = append(n, coord{row - 1, col})
	}
	if col-1 >= 0 {
		n = append(n, coord{row, col - 1})
	}
	if row+1 < h {
		n = append(n, coord{row + 1, col})
	}
	if col+1 < w {
		n = append(n, coord{row, col + 1})
	}

	return n
}

func climbable(curr, next rune) bool {
	return (next - curr) <= 1
}

func climb(hm [][]rune, start, end coord) int {
	visited := make(map[coord]bool)
	path := []coord{start}
	dist := map[coord]int{start: 0}
	w, h := len(hm[0]), len(hm)

	for {
		curr := path[0]
		visited[curr] = true
		path = path[1:]

		if curr == end {
			return dist[end]
		}

		ns := neighbours(curr, w, h)
		for _, neighbour := range ns {
			if !visited[neighbour] && climbable(hm[curr.x][curr.y], hm[neighbour.x][neighbour.y]) {
				if _, ok := dist[neighbour]; !ok {
					path = append(path, neighbour)
				}
				dist[neighbour] = dist[curr] + 1
			}
		}

		sort.Slice(path, func(i, j int) bool { return dist[path[i]] < dist[path[j]] })
	}
}

func climbCheat(hm [][]rune, start, end coord) int {
	visited := make(map[coord]bool)
	path := []coord{start}
	dist := map[coord]int{start: 0}
	w, h := len(hm[0]), len(hm)

	for {
		curr := path[0]
		visited[curr] = true
		path = path[1:]

		if curr == end {
			return dist[end]
		}

		ns := neighbours(curr, w, h)
		for _, neighbour := range ns {
			if !visited[neighbour] && climbable(hm[curr.x][curr.y], hm[neighbour.x][neighbour.y]) {
				if _, ok := dist[neighbour]; !ok {
					path = append(path, neighbour)
				}
				if hm[neighbour.x][neighbour.y] == 'a' {
					dist[neighbour] = 0
				} else {
					dist[neighbour] = dist[curr] + 1
				}
				// dist[neighbour] = dist[curr] + 1
			}
		}

		sort.Slice(path, func(i, j int) bool { return dist[path[i]] < dist[path[j]] })
	}
}

func HillClimb() int {
	// hm, s, e := parseMap("./day12/example.txt")
	hm, s, e := parseMap("./day12/input.txt")

	if len(hm) == 0 {
		fmt.Println("Error occurred while parsing map.")
		return 0
	}

	return climb(hm, s, e)
}

func HillClimbCheat() int {
	// hm, s, e := parseMap("./day12/example.txt")
	hm, s, e := parseMap("./day12/input.txt")

	if len(hm) == 0 {
		fmt.Println("Error occurred while parsing map.")
		return 0
	}

	return climbCheat(hm, s, e) - 1
}
