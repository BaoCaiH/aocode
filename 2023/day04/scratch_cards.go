package day04

import (
	"math"
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

func isPartOf(val string, data []string) bool {
	for _, v := range data {
		if val == v {
			return true
		}
	}
	return false
}

func Scratch() int {
	sc, file := utils.Read("./day04/input.txt")
	// sc, file := utils.Read("./day04/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	total, cnt := 0, -1
	line := ""
	var lineParts, winning, yours []string
	for sc.Scan() {
		line = sc.Text()
		lineParts = strings.Split(strings.Split(line, ":")[1], "|")
		winning = strings.Split(strings.ReplaceAll(strings.Trim(lineParts[0], " "), "  ", " "), " ")
		yours = strings.Split(strings.ReplaceAll(strings.Trim(lineParts[1], " "), "  ", " "), " ")

		for _, num := range yours {
			if isPartOf(num, winning) {
				cnt += 1
			}
		}
		if cnt != -1 {
			total += int(math.Pow(float64(2), float64(cnt)))
		}
		cnt = -1
	}

	return total
}

func MoreScratch() int {
	sc, file := utils.Read("./day04/input.txt")
	// sc, file := utils.Read("./day04/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	total, cnt, mult := 0, -1, 1
	line := ""
	var bonus []int
	var lineParts, winning, yours []string
	for sc.Scan() {
		line = sc.Text()
		lineParts = strings.Split(strings.Split(line, ":")[1], "|")
		winning = strings.Split(strings.ReplaceAll(strings.Trim(lineParts[0], " "), "  ", " "), " ")
		yours = strings.Split(strings.ReplaceAll(strings.Trim(lineParts[1], " "), "  ", " "), " ")

		for _, num := range yours {
			if isPartOf(num, winning) {
				cnt += 1
			}
		}

		total += 1
		if len(bonus) > 0 {
			total += bonus[0]
			mult = bonus[0] + 1
			bonus = bonus[1:]
		}

		for i := 0; i <= cnt; i++ {
			if len(bonus) <= i {
				bonus = append(bonus, 0)
			}
			bonus[i] += 1 * mult
		}

		cnt = -1
		mult = 1
	}

	return total
}
