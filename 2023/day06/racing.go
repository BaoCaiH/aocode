package day06

import (
	"math"
	"strconv"
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

func Press() int {
	sc, file := utils.Read("./day06/input.txt")
	// sc, file := utils.Read("./day06/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	var tmpArr []string
	var times, distance []int

	line := ""
	sc.Scan()
	line = sc.Text()
	tmpArr = strings.Split(strings.Split(line, ":")[1], " ")
	for _, v := range tmpArr {
		if len(v) == 0 {
			continue
		}
		n, _ := strconv.Atoi(strings.Trim(v, " "))
		times = append(times, n)
	}

	sc.Scan()
	line = sc.Text()
	tmpArr = strings.Split(strings.Split(line, ":")[1], " ")
	for _, v := range tmpArr {
		if len(v) == 0 {
			continue
		}
		n, _ := strconv.Atoi(strings.Trim(v, " "))
		distance = append(distance, n)
	}

	total := 1

	for i := 0; i < len(times); i++ {
		x := float64(times[i])
		z := distance[i]
		tmp := float64(x*x) - float64(4*z)
		left := (float64(x) - math.Sqrt(tmp)) / 2
		right := (float64(x) + math.Sqrt(tmp)) / 2
		total *= int(right) - int(left)
	}
	return total
}

func PressLong() int {
	sc, file := utils.Read("./day06/input.txt")
	// sc, file := utils.Read("./day06/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	var tmpStr string
	var time, distance int

	line := ""
	sc.Scan()
	line = sc.Text()
	tmpStr = strings.Split(line, ":")[1]
	time, _ = strconv.Atoi(strings.ReplaceAll(tmpStr, " ", ""))
	sc.Scan()
	line = sc.Text()
	tmpStr = strings.Split(line, ":")[1]
	distance, _ = strconv.Atoi(strings.ReplaceAll(tmpStr, " ", ""))

	x := float64(time)
	z := distance
	tmp := float64(x*x) - float64(4*z)
	left := (float64(x) - math.Sqrt(tmp)) / 2
	right := (float64(x) + math.Sqrt(tmp)) / 2
	return int(right) - int(left)
}
