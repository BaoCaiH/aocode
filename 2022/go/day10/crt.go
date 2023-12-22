package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CathodeRayTube() int {
	// file, err := os.Open("./day10/example.txt")
	file, err := os.Open("./day10/input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	cnt, x, sumSignals := 0, 1, 0
	signals := []int{}
	for sc.Scan() {
		instr := sc.Text()
		cnt++
		if instr == "noop" {
			if r, s := calcSignal(cnt, x); r {
				sumSignals += s
				signals = append(signals, s)
			}
		} else {
			if r, s := calcSignal(cnt, x); r {
				sumSignals += s
				signals = append(signals, s)
			}
			cnt++
			if r, s := calcSignal(cnt, x); r {
				sumSignals += s
				signals = append(signals, s)
			}
			addx, _ := strconv.Atoi(strings.Split(instr, " ")[1])
			x += addx
		}
	}
	fmt.Println(signals)
	return sumSignals
}

func CathodeRayTubePixel() {
	// file, err := os.Open("./day10/example.txt")
	file, err := os.Open("./day10/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	cnt, x := 0, 1
	for sc.Scan() {
		instr := sc.Text()
		cnt++
		if instr == "noop" {
			printPixel(&cnt, x)
		} else {
			printPixel(&cnt, x)
			cnt++
			printPixel(&cnt, x)
			addx, _ := strconv.Atoi(strings.Split(instr, " ")[1])
			x += addx
		}
	}
	fmt.Println()
}

func calcSignal(cnt, x int) (bool, int) {
	if cnt == 20 || (cnt-20)%40 == 0 {
		return true, cnt * x
	}
	return false, 0
}

func printPixel(pos *int, cur int) {
	if *pos == 41 {
		*pos = 1
		fmt.Println()
	}
	if (*pos%41 >= cur) && (*pos%41 <= cur+2) {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}
