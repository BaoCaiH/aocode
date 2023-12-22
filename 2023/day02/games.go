package day02

import (
	"strconv"
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

func PossibleGames() int {
	sc, file := utils.Read("./day02/input.txt")
	// sc, file := utils.Read("./day02/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	var line, tmpStr string
	var idCurr, setupNumber, total int
	var lineParts, tmpSetup []string
	var ids []int
	limits := map[string]int{"red": 12, "green": 13, "blue": 14}
	var possibility bool

	for sc.Scan() {
		line = sc.Text()
		if len(line) == 0 {
			continue
		}
		lineParts = strings.Split(line, ":")
		tmpStr = strings.Split(lineParts[0], " ")[1]
		idCurr, _ = strconv.Atoi(tmpStr)
		lineParts = strings.Split(lineParts[1], ";")
		possibility = true
		for _, part := range lineParts {
			for _, setup := range strings.Split(part, ",") {
				tmpSetup = strings.Split(strings.Trim(setup, " "), " ")
				setupNumber, _ = strconv.Atoi(tmpSetup[0])
				if setupNumber > limits[tmpSetup[1]] {
					possibility = false
				}
			}
		}
		if possibility {
			ids = append(ids, idCurr)
			total += idCurr
		}
	}

	return total
}

func PowerGames() int {
	sc, file := utils.Read("./day02/input.txt")
	// sc, file := utils.Read("./day02/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	var line string
	var setupNumber, total int
	var lineParts, tmpSetup []string
	var gameSetup map[string]int

	for sc.Scan() {
		line = sc.Text()
		if len(line) == 0 {
			continue
		}
		lineParts = strings.Split(line, ":")
		lineParts = strings.Split(lineParts[1], ";")
		gameSetup = map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, part := range lineParts {
			for _, setup := range strings.Split(part, ",") {
				tmpSetup = strings.Split(strings.Trim(setup, " "), " ")
				setupNumber, _ = strconv.Atoi(tmpSetup[0])
				if gameSetup[tmpSetup[1]] == 0 || gameSetup[tmpSetup[1]] < setupNumber {
					gameSetup[tmpSetup[1]] = setupNumber
				}
			}
		}
		total += gameSetup["red"] * gameSetup["green"] * gameSetup["blue"]
	}

	return total
}
