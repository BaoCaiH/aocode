package day01

import (
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

func indexOf(element string, data []string) int {
	for i, v := range data {
		if element == v {
			return i
		}
	}
	return -1
}

func Calibrate() int {
	sc, file := utils.Read("./day01/input.txt")
	// sc, file := utils.Read("./day01/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	total, first, firstIndex, last, lastIndex, tmpIndex := 0, -1, -1, -1, -1, -1
	var line string
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for sc.Scan() {
		line = sc.Text()
		for _, digit := range digits {
			tmpIndex = strings.Index(line, digit)
			if (tmpIndex != -1 && firstIndex == -1) || (tmpIndex != -1 && tmpIndex < firstIndex) {
				firstIndex = tmpIndex
				first = indexOf(digit, digits)
			}
			tmpIndex = strings.LastIndex(line, digit)
			if (tmpIndex != -1 && lastIndex == -1) || (tmpIndex != -1 && tmpIndex > lastIndex) {
				lastIndex = tmpIndex
				last = indexOf(digit, digits)
			}
		}
		total += first*10 + last
		firstIndex, lastIndex = -1, -1
	}
	return total
}

func CalibrateText() int {
	sc, file := utils.Read("./day01/input.txt")
	// sc, file := utils.Read("./day01/input_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	digitStrings := []string{"||", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	total, first, firstIndex, last, lastIndex, tmpIndex := 0, -1, -1, -1, -1, -1
	var line string

	for sc.Scan() {
		line = sc.Text()
		for _, digit := range digits {
			tmpIndex = strings.Index(line, digit)
			if (tmpIndex != -1 && firstIndex == -1) || (tmpIndex != -1 && tmpIndex < firstIndex) {
				firstIndex = tmpIndex
				first = indexOf(digit, digits)
			}
			tmpIndex = strings.LastIndex(line, digit)
			if (tmpIndex != -1 && lastIndex == -1) || (tmpIndex != -1 && tmpIndex > lastIndex) {
				lastIndex = tmpIndex
				last = indexOf(digit, digits)
			}
		}
		for _, digitString := range digitStrings {
			tmpIndex = strings.Index(line, digitString)
			if (tmpIndex != -1 && firstIndex == -1) || (tmpIndex != -1 && tmpIndex < firstIndex) {
				firstIndex = tmpIndex
				first = indexOf(digitString, digitStrings)
			}
			tmpIndex = strings.LastIndex(line, digitString)
			if (tmpIndex != -1 && lastIndex == -1) || (tmpIndex != -1 && tmpIndex > lastIndex) {
				lastIndex = tmpIndex
				last = indexOf(digitString, digitStrings)
			}
		}

		total += first*10 + last
		firstIndex, lastIndex = -1, -1
	}

	return total

}
