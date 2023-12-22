package day15

import (
	"strconv"
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

func hash(code []rune) int {
	curr := 0
	for _, char := range code {
		curr += int(char)
		curr *= 17
		curr = curr % 256
	}
	return curr
}

func Asciii() int {
	sc, file := utils.Read("./day15/input.txt")
	// sc, file := utils.Read("./day15/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	sc.Scan()
	line := sc.Text()
	inputs := [][]rune{}
	for _, code := range strings.Split(line, ",") {
		inputs = append(inputs, []rune(code))
	}

	total := 0
	for _, code := range inputs {
		total += hash(code)
	}

	return total
}

type Lens struct {
	label string
	fl    int
}

func dropLens(box []Lens, label string) []Lens {
	for i, content := range box {
		if content.label == label {
			return append(box[:i], box[i+1:]...)
		}
	}
	return box
}

func changeLens(box []Lens, label string, fl int) []Lens {
	for i, content := range box {
		if content.label == label {
			box[i] = Lens{label, fl}
			return box
		}
	}
	return append(box, Lens{label, fl})
}

func Focal() int {
	sc, file := utils.Read("./day15/input.txt")
	// sc, file := utils.Read("./day15/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	sc.Scan()
	line := sc.Text()
	inputs := strings.Split(line, ",")

	boxes := [][]Lens{}
	for i := 0; i < 256; i++ {
		boxes = append(boxes, []Lens{})
	}

	total := 0
	label := ""
	labelInt := -1
	flStr := ""
	fl := -1
	for _, code := range inputs {
		if strings.Contains(code, "-") {
			label = strings.Split(code, "-")[0]
			labelInt = hash([]rune(label))
			boxes[labelInt] = dropLens(boxes[labelInt], label)
			continue
		}
		label = strings.Split(code, "=")[0]
		flStr = strings.Split(code, "=")[1]
		labelInt = hash([]rune(label))
		fl, _ = strconv.Atoi(flStr)
		boxes[labelInt] = changeLens(boxes[labelInt], label, fl)
	}

	for i, box := range boxes {
		for j, content := range box {
			total += (i + 1) * (j + 1) * content.fl
		}
	}

	return total
}
