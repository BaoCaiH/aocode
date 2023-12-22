package day06

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Tuning(l int) int {
	file, err := os.Open("./day06/input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Scan()
	signal := sc.Text()

	signalLength := len(signal)
	for i := range signal {
		if signalLength-i < l {
			return 0
		}

		nextFour := signal[i+1 : i+l+1]
		diffChars := ""
		for _, c := range nextFour {
			if strings.Contains(diffChars, string(c)) {
				continue
			}
			diffChars += string(c)
		}
		if len(diffChars) == l {
			return i + l + 1
		}
	}

	return 0
}
