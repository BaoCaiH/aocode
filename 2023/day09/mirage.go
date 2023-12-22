package day09

import (
	"strconv"
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

func allZero(arr []int) bool {
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

func Mirage() int {
	sc, file := utils.Read("./day09/input.txt")
	// sc, file := utils.Read("./day09/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	line := ""
	var tmpInt, length int
	var tmpStrArr []string
	var tmpIntArr, currIntArr []int
	total := 0

	for sc.Scan() {
		line = sc.Text()
		tmpStrArr = strings.Split(line, " ")
		currIntArr = []int{}
		for _, v := range tmpStrArr {
			tmpInt, _ = strconv.Atoi(v)
			currIntArr = append(currIntArr, tmpInt)
		}

		for !allZero(currIntArr) {
			tmpIntArr = []int{}
			length = len(currIntArr)
			for i, v := range currIntArr {
				if i == length-1 {
					total += v
					continue
				}
				tmpIntArr = append(tmpIntArr, currIntArr[i+1]-v)
			}
			currIntArr = tmpIntArr
		}

	}

	return total
}

func Egarim() int {
	sc, file := utils.Read("./day09/input.txt")
	// sc, file := utils.Read("./day09/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	line := ""
	var tmpInt, length int
	var tmpStrArr []string
	var tmpIntArr, currIntArr, firsts []int
	total := 0

	for sc.Scan() {
		line = sc.Text()
		tmpStrArr = strings.Split(line, " ")
		currIntArr = []int{}
		for _, v := range tmpStrArr {
			tmpInt, _ = strconv.Atoi(v)
			currIntArr = append(currIntArr, tmpInt)
		}

		firsts = []int{}
		for !allZero(currIntArr) {
			tmpIntArr = []int{}
			length = len(currIntArr)
			for i, v := range currIntArr {
				if i == 0 {
					firsts = append(firsts, v)
				}
				if i == length-1 {
					continue
				}
				tmpIntArr = append(tmpIntArr, currIntArr[i+1]-v)
			}
			currIntArr = tmpIntArr
		}

		for i := len(firsts) - 1; i >= 0; i-- {
			if i == len(firsts)-1 {
				tmpInt = firsts[i]
				continue
			}
			if i == 0 {
				total += firsts[i] - tmpInt
				continue
			}
			tmpInt = firsts[i] - tmpInt
		}
	}

	return total
}
