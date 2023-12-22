package day09

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func RopeBridge() int {

	file, err := os.Open("./day09/input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	head, tail := []int{0, 0}, []int{0, 0}
	visited := [][]int{}
	for sc.Scan() {
		move := strings.Split(sc.Text(), " ")
		steps, _ := strconv.Atoi(move[1])
		for i := 0; i < steps; i++ {
			switch move[0] {
			case "L":
				head[1] = head[1] - 1
			case "R":
				head[1] = head[1] + 1
			case "U":
				head[0] = head[0] + 1
			case "D":
				head[0] = head[0] - 1
			}
			if isTooFar, dir := tooFar(head, tail); isTooFar {
				tail = moveToward(tail, dir)
			}
			visited = addIfNotExists(visited, tail)
		}
	}

	return len(visited)
}

func RopeBridgeSnap() int {

	file, err := os.Open("./day09/input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	nodes := [10][]int{}
	for i := range nodes {
		nodes[i] = []int{0, 0}
	}
	visited := [][]int{}
	for sc.Scan() {
		move := strings.Split(sc.Text(), " ")
		steps, _ := strconv.Atoi(move[1])
		for i := 0; i < steps; i++ {
			switch move[0] {
			case "L":
				nodes[0][1] = nodes[0][1] - 1
			case "R":
				nodes[0][1] = nodes[0][1] + 1
			case "U":
				nodes[0][0] = nodes[0][0] + 1
			case "D":
				nodes[0][0] = nodes[0][0] - 1
			}
			for j := 1; j < 10; j++ {
				if isTooFar, dir := tooFar(nodes[j-1], nodes[j]); isTooFar {
					nodes[j] = moveToward(nodes[j], dir)
				}
			}
			visited = addIfNotExists(visited, nodes[9])
		}
	}
	return len(visited)
}

func addIfNotExists(visited [][]int, tail []int) [][]int {
	for _, pos := range visited {
		if pos[0] == tail[0] && pos[1] == tail[1] {
			return visited
		}
	}
	return append(visited, tail)
}

func tooFar(head, tail []int) (bool, []int) {
	if math.Abs(float64(head[0]-tail[0])) > 1 || math.Abs(float64(head[1]-tail[1])) > 1 {
		return true, []int{clamp(head[0]-tail[0], 1, -1), clamp(head[1]-tail[1], 1, -1)}
	}
	return false, nil
}

func moveToward(tail, direction []int) []int {
	return []int{tail[0] + direction[0], tail[1] + direction[1]}
}

func clamp(n, u, l int) int {
	if n < l {
		return l
	}
	if n > u {
		return u
	}
	return n
}
