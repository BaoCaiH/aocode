package day10

import (
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

type Coor struct {
	r, c int
}

type Node struct {
	links []Coor
}

func notIn(n Coor, past []Coor) bool {
	for _, node := range past {
		if node.c == n.c && node.r == n.r {
			return false
		}
	}
	return true
}

func AMaze() int {
	sc, file := utils.Read("./day10/input.txt")
	// sc, file := utils.Read("./day10/input_2_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	line := ""
	inputs := []string{}
	maze := [][]Node{}
	var tmpNode Node
	var start Coor
	for sc.Scan() {
		line = sc.Text()
		inputs = append(inputs, line)
	}

	height := len(inputs)
	length := len(inputs[0])
	for x := 0; x < height; x++ {
		tmpRow := []Node{}
		for y := 0; y < length; y++ {
			tmpNode = Node{}
			if string(inputs[x][y]) == "|" {
				if x != 0 {
					tmpNode.links = append(tmpNode.links, Coor{x - 1, y})
				}
				if x < height-1 {
					tmpNode.links = append(tmpNode.links, Coor{x + 1, y})
				}
			}
			if string(inputs[x][y]) == "-" {
				if y != 0 {
					tmpNode.links = append(tmpNode.links, Coor{x, y - 1})
				}
				if y < length-1 {
					tmpNode.links = append(tmpNode.links, Coor{x, y + 1})
				}
			}
			if string(inputs[x][y]) == "F" {
				if x < height-1 {
					tmpNode.links = append(tmpNode.links, Coor{x + 1, y})
				}
				if y < length-1 {
					tmpNode.links = append(tmpNode.links, Coor{x, y + 1})
				}
			}
			if string(inputs[x][y]) == "L" {
				if x != 0 {
					tmpNode.links = append(tmpNode.links, Coor{x - 1, y})
				}
				if y < length-1 {
					tmpNode.links = append(tmpNode.links, Coor{x, y + 1})
				}
			}
			if string(inputs[x][y]) == "J" {
				if x != 0 {
					tmpNode.links = append(tmpNode.links, Coor{x - 1, y})
				}
				if y != 0 {
					tmpNode.links = append(tmpNode.links, Coor{x, y - 1})
				}
			}
			if string(inputs[x][y]) == "7" {
				if x < height-1 {
					tmpNode.links = append(tmpNode.links, Coor{x + 1, y})
				}
				if y != 0 {
					tmpNode.links = append(tmpNode.links, Coor{x, y - 1})
				}
			}
			if string(inputs[x][y]) == "S" {
				if x != 0 && (string(inputs[x-1][y]) == "|" || string(inputs[x-1][y]) == "7" || string(inputs[x-1][y]) == "F") {
					tmpNode.links = append(tmpNode.links, Coor{x - 1, y})
				}
				if x < height-1 && (string(inputs[x+1][y]) == "|" || string(inputs[x+1][y]) == "L" || string(inputs[x+1][y]) == "J") {
					tmpNode.links = append(tmpNode.links, Coor{x + 1, y})
				}
				if y != 0 && (string(inputs[x][y-1]) == "-" || string(inputs[x][y-1]) == "L" || string(inputs[x][y-1]) == "F") {
					tmpNode.links = append(tmpNode.links, Coor{x, y - 1})
				}
				if y < length-1 && (string(inputs[x][y+1]) == "-" || string(inputs[x][y+1]) == "J" || string(inputs[x][y+1]) == "7") {
					tmpNode.links = append(tmpNode.links, Coor{x, y + 1})
				}
				start = Coor{x, y}
			}
			tmpRow = append(tmpRow, tmpNode)
		}
		maze = append(maze, tmpRow)
	}

	past := []Coor{start}
	currNode := maze[start.r][start.c].links[0]

	for currNode.r != start.r || currNode.c != start.c {
		tmpNode = maze[currNode.r][currNode.c]
		past = append(past, currNode)
		if notIn(tmpNode.links[0], past) {
			currNode = tmpNode.links[0]
		} else if notIn(tmpNode.links[1], past) {
			currNode = tmpNode.links[1]
		} else {
			break
		}
	}

	return (len(past) + 1) / 2
}

func filter(ori []string, filt string) []string {
	temp := []string{}
	for _, c := range ori {
		if strings.Contains(filt, c) {
			temp = append(temp, c)
		}
	}
	return temp
}

func AMazeZing() int {
	sc, file := utils.Read("./day10/input.txt")
	// sc, file := utils.Read("./day10/input_2_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	line := ""
	inputs := []string{}
	maze := [][]Node{}
	var tmpNode Node
	var start Coor
	for sc.Scan() {
		line = sc.Text()
		inputs = append(inputs, line)
	}

	height := len(inputs)
	length := len(inputs[0])
	possibleS := []string{"-", "|", "F", "J", "L", "7"}
	for x := 0; x < height; x++ {
		tmpRow := []Node{}
		for y := 0; y < length; y++ {
			tmpNode = Node{}
			if string(inputs[x][y]) == "|" {
				if x != 0 {
					tmpNode.links = append(tmpNode.links, Coor{x - 1, y})
				}
				if x < height-1 {
					tmpNode.links = append(tmpNode.links, Coor{x + 1, y})
				}
			}
			if string(inputs[x][y]) == "-" {
				if y != 0 {
					tmpNode.links = append(tmpNode.links, Coor{x, y - 1})
				}
				if y < length-1 {
					tmpNode.links = append(tmpNode.links, Coor{x, y + 1})
				}
			}
			if string(inputs[x][y]) == "F" {
				if x < height-1 {
					tmpNode.links = append(tmpNode.links, Coor{x + 1, y})
				}
				if y < length-1 {
					tmpNode.links = append(tmpNode.links, Coor{x, y + 1})
				}
			}
			if string(inputs[x][y]) == "L" {
				if x != 0 {
					tmpNode.links = append(tmpNode.links, Coor{x - 1, y})
				}
				if y < length-1 {
					tmpNode.links = append(tmpNode.links, Coor{x, y + 1})
				}
			}
			if string(inputs[x][y]) == "J" {
				if x != 0 {
					tmpNode.links = append(tmpNode.links, Coor{x - 1, y})
				}
				if y != 0 {
					tmpNode.links = append(tmpNode.links, Coor{x, y - 1})
				}
			}
			if string(inputs[x][y]) == "7" {
				if x < height-1 {
					tmpNode.links = append(tmpNode.links, Coor{x + 1, y})
				}
				if y != 0 {
					tmpNode.links = append(tmpNode.links, Coor{x, y - 1})
				}
			}
			if string(inputs[x][y]) == "S" {
				if x != 0 && (string(inputs[x-1][y]) == "|" || string(inputs[x-1][y]) == "7" || string(inputs[x-1][y]) == "F") {
					tmpNode.links = append(tmpNode.links, Coor{x - 1, y})
					possibleS = filter(possibleS, "|JL")
				}
				if x < height-1 && (string(inputs[x+1][y]) == "|" || string(inputs[x+1][y]) == "L" || string(inputs[x+1][y]) == "J") {
					tmpNode.links = append(tmpNode.links, Coor{x + 1, y})
					possibleS = filter(possibleS, "|F7")
				}
				if y != 0 && (string(inputs[x][y-1]) == "-" || string(inputs[x][y-1]) == "L" || string(inputs[x][y-1]) == "F") {
					tmpNode.links = append(tmpNode.links, Coor{x, y - 1})
					possibleS = filter(possibleS, "-J7")
				}
				if y < length-1 && (string(inputs[x][y+1]) == "-" || string(inputs[x][y+1]) == "J" || string(inputs[x][y+1]) == "7") {
					tmpNode.links = append(tmpNode.links, Coor{x, y + 1})
					possibleS = filter(possibleS, "-LF")
				}
				start = Coor{x, y}
			}
			tmpRow = append(tmpRow, tmpNode)
		}
		maze = append(maze, tmpRow)
	}

	inputs[start.r] = strings.ReplaceAll(inputs[start.r], "S", possibleS[0])

	past := []Coor{start}
	currNode := maze[start.r][start.c].links[0]

	for currNode.r != start.r || currNode.c != start.c {
		tmpNode = maze[currNode.r][currNode.c]
		past = append(past, currNode)
		if notIn(tmpNode.links[0], past) {
			currNode = tmpNode.links[0]
		} else if notIn(tmpNode.links[1], past) {
			currNode = tmpNode.links[1]
		} else {
			break
		}
	}

	for r, line := range inputs {
		tmp := ""
		for c, ch := range line {
			if notIn(Coor{r, c}, past) {
				tmp += "."
			} else {
				tmp += string(ch)
			}
		}
		inputs[r] = tmp
	}
	// outside := []Coor{}

	cnt := 0

	for r, row := range inputs {
		inside := false
		goUp := "unset"
		for c, col := range row {
			if col == '|' && goUp == "unset" {
				inside = !inside
			} else if strings.Contains("LF", string(col)) {
				if col == 'L' {
					goUp = "true"
				} else {
					goUp = "false"
				}
			} else if strings.Contains("J7", string(col)) {
				if goUp == "true" && col != 'J' || goUp == "false" && col != '7' {
					inside = !inside
				}
				goUp = "unset"
			}

			if !inside {
				// outside = append(outside, Coor{r, c})
				// fmt.Print("#")
			} else if notIn(Coor{r, c}, past) {
				cnt += 1
				// 	fmt.Print(".")
				// } else{
				// 	fmt.Print("#")
			}
		}
		// fmt.Println()
	}

	return cnt
}
