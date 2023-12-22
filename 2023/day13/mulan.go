package day13

import (
	"github.com/BaoCaiH/aocode/2023/utils"
)

func verticalMirror(m []string, i int) bool {
	length := len(m[0])
	if i == length-1 {
		return false
	}
	availableLeft := i + 1
	remainingRight := length - (i + 1)
	nToCheck := 0
	if availableLeft < remainingRight {
		nToCheck = availableLeft
	} else {
		nToCheck = remainingRight
	}
	for _, line := range m {
		for j := 0; j < nToCheck; j++ {
			if line[i-j] != line[availableLeft+j] {
				return false
			}
		}
	}
	return true
}

func horizontalMirror(m []string, i int) bool {
	length := len(m)
	if i == length-1 {
		return false
	}
	availableUp := i + 1
	remainingDown := length - (i + 1)
	nToCheck := 0
	if availableUp < remainingDown {
		nToCheck = availableUp
	} else {
		nToCheck = remainingDown
	}
	for j := 0; j < nToCheck; j++ {
		if m[i-j] != m[availableUp+j] {
			return false
		}
	}
	return true
}

func Reflection() int {
	sc, file := utils.Read("./day13/input.txt")
	// sc, file := utils.Read("./day13/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	mirrors := [][]string{}
	currMirror := []string{}

	line := ""
	for sc.Scan() {
		line = sc.Text()
		if len(line) == 0 {
			mirrors = append(mirrors, currMirror)
			currMirror = []string{}
			continue
		}

		currMirror = append(currMirror, line)
	}
	if len(currMirror) > 0 {
		mirrors = append(mirrors, currMirror)
	}

	total := 0
	for _, m := range mirrors {
		// Horizontal
		for i := range m {
			if horizontalMirror(m, i) {
				total += 100 * (i + 1)
			}
		}

		// Verticle
		for i := range m[0] {
			if verticalMirror(m, i) {
				total += i + 1
			}
		}
	}

	return total
}

func verticalMirrorSmudge(m []string, i int) bool {
	length := len(m[0])
	if i == length-1 {
		return false
	}
	allowSmudge := 1
	availableLeft := i + 1
	remainingRight := length - (i + 1)
	nToCheck := 0
	if availableLeft < remainingRight {
		nToCheck = availableLeft
	} else {
		nToCheck = remainingRight
	}
	for _, line := range m {
		for j := 0; j < nToCheck; j++ {
			if line[i-j] != line[availableLeft+j] {
				if allowSmudge == 0 {
					return false
				}
				allowSmudge -= 1
			}
		}
	}
	if allowSmudge != 0 {
		return false
	}
	return true
}

func horizontalMirrorSmudge(m []string, i int) bool {
	length := len(m)
	if i == length-1 {
		return false
	}
	allowSmudge := 1
	availableUp := i + 1
	remainingDown := length - (i + 1)
	nToCheck := 0
	if availableUp < remainingDown {
		nToCheck = availableUp
	} else {
		nToCheck = remainingDown
	}
	for j := 0; j < nToCheck; j++ {
		if m[i-j] != m[availableUp+j] {
			for k := range m[i-j] {
				if m[i-j][k] != m[availableUp+j][k] {
					if allowSmudge == 0 {
						return false
					}
					allowSmudge -= 1
				}
			}
		}
	}
	if allowSmudge != 0 {
		return false
	}
	return true
}

func ReflectionSmudge() int {
	sc, file := utils.Read("./day13/input.txt")
	// sc, file := utils.Read("./day13/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	mirrors := [][]string{}
	currMirror := []string{}

	line := ""
	for sc.Scan() {
		line = sc.Text()
		if len(line) == 0 {
			mirrors = append(mirrors, currMirror)
			currMirror = []string{}
			continue
		}

		currMirror = append(currMirror, line)
	}
	if len(currMirror) > 0 {
		mirrors = append(mirrors, currMirror)
	}

	total := 0
	for _, m := range mirrors {
		// Horizontal
		for i := range m {
			if horizontalMirrorSmudge(m, i) {
				total += 100 * (i + 1)
			}
		}

		// Verticle
		for i := range m[0] {
			if verticalMirrorSmudge(m, i) {
				total += i + 1
			}
		}
	}

	return total
}
