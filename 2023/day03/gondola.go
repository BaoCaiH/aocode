package day03

import (
	"fmt"
	"strconv"

	"github.com/BaoCaiH/aocode/2023/utils"
)

func isIn(element string, data []string) bool {
	for _, v := range data {
		if element == v {
			return true
		}
	}
	return false
}

func isSymbol(element string) bool {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	return element != "." && !isIn(element, digits)
}

func adjacent(x, y int, grid []string) bool {
	width := len(grid[0])
	height := len(grid)
	if x-1 > 0 && y-1 > 0 && isSymbol(string(grid[x-1][y-1])) {
		return true
	}
	if y-1 > 0 && isSymbol(string(grid[x][y-1])) {
		return true
	}
	if x+1 < width && y-1 > 0 && isSymbol(string(grid[x+1][y-1])) {
		return true
	}
	if x-1 > 0 && isSymbol(string(grid[x-1][y])) {
		return true
	}
	if x+1 < width && isSymbol(string(grid[x+1][y])) {
		return true
	}
	if x-1 > 0 && y+1 < height && isSymbol(string(grid[x-1][y+1])) {
		return true
	}
	if y+1 < height && isSymbol(string(grid[x][y+1])) {
		return true
	}
	if x+1 < width && y+1 < height && isSymbol(string(grid[x+1][y+1])) {
		return true
	}
	return false
}

func gearRatio(x, y int, grid []string) string {
	width := len(grid[0])
	height := len(grid)
	if x-1 > 0 && y-1 > 0 && grid[x-1][y-1] == '*' {
		return fmt.Sprintf("%c:%c", x-1, y-1)
	}
	if y-1 > 0 && grid[x][y-1] == '*' {
		return fmt.Sprintf("%c:%c", x, y-1)
	}
	if x+1 < width && y-1 > 0 && grid[x+1][y-1] == '*' {
		return fmt.Sprintf("%c:%c", x+1, y-1)
	}
	if x-1 > 0 && grid[x-1][y] == '*' {
		return fmt.Sprintf("%c:%c", x-1, y)
	}
	if x+1 < width && grid[x+1][y] == '*' {
		return fmt.Sprintf("%c:%c", x+1, y)
	}
	if x-1 > 0 && y+1 < height && grid[x-1][y+1] == '*' {
		return fmt.Sprintf("%c:%c", x-1, y+1)
	}
	if y+1 < height && grid[x][y+1] == '*' {
		return fmt.Sprintf("%c:%c", x, y+1)
	}
	if x+1 < width && y+1 < height && grid[x+1][y+1] == '*' {
		return fmt.Sprintf("%c:%c", x+1, y+1)
	}
	return ""
}

func Gondola() int {
	sc, file := utils.Read("./day03/input.txt")
	// sc, file := utils.Read("./day03/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	var grid []string
	for sc.Scan() {
		grid = append(grid, sc.Text())
	}

	numberStr := ""
	numberMarker := false
	total := 0
	var tmp int
	for x, line := range grid {
		for y, char := range line {
			if char == '.' || isSymbol(string(char)) {
				if numberStr != "" && numberMarker {
					tmp, _ = strconv.Atoi(numberStr)
					total += tmp
				}
				numberStr = ""
				numberMarker = false
			} else if char != '.' && !isSymbol(string(char)) {
				numberStr += string(char)
				numberMarker = adjacent(x, y, grid) || numberMarker
			}
		}
		if numberStr != "" && numberMarker {
			tmp, _ = strconv.Atoi(numberStr)
			total += tmp
		}
		numberStr = ""
		numberMarker = false
	}

	return total
}

func GearRatio() int {
	sc, file := utils.Read("./day03/input.txt")
	// sc, file := utils.Read("./day03/input_1_example.txt")
	fmt.Print("")
	if sc == nil {
		return 0
	}
	defer file.Close()

	var grid []string
	for sc.Scan() {
		grid = append(grid, sc.Text())
	}

	numberStr := ""
	numberMarker := ""
	gears := map[string]int{}
	total := 0
	var tmp int
	for x, line := range grid {
		for y, char := range line {
			if char == '.' || isSymbol(string(char)) {
				if numberStr != "" && numberMarker != "" {
					tmp, _ = strconv.Atoi(numberStr)
					if gears[numberMarker] != 0 {
						total += gears[numberMarker] * tmp
					} else {
						gears[numberMarker] = tmp
					}
				}
				numberStr = ""
				numberMarker = ""
			} else if char != '.' && !isSymbol(string(char)) {
				numberStr += string(char)
				if numberMarker == "" {
					numberMarker = gearRatio(x, y, grid)
				}
			}
		}
		if numberStr != "" && numberMarker != "" {
			tmp, _ = strconv.Atoi(numberStr)
			if gears[numberMarker] != 0 {
				total += gears[numberMarker] * tmp
			} else {
				gears[numberMarker] = tmp
			}
		}
		numberStr = ""
		numberMarker = ""
	}

	return total
}
