package day08

import (
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

type Instr struct {
	l, r string
}

func LeftRight() int {
	sc, file := utils.Read("./day08/input.txt")
	// sc, file := utils.Read("./day08/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	sc.Scan()
	instruction := sc.Text()

	line := ""
	var tmpStrArr, tmpLR []string
	mapInstr := map[string]Instr{}

	for sc.Scan() {
		line = sc.Text()
		if len(line) == 0 {
			continue
		}

		tmpStrArr = strings.Split(line, " = ")
		tmpLR = strings.Split(tmpStrArr[1], ", ")
		mapInstr[tmpStrArr[0]] = Instr{strings.Trim(tmpLR[0], "("), strings.Trim(tmpLR[1], ")")}
	}

	nodeCurr := "AAA"
	cnt := 0

	for nodeCurr != "ZZZ" {
		for _, instr := range instruction {
			if string(instr) == "L" {
				nodeCurr = mapInstr[nodeCurr].l
			} else {
				nodeCurr = mapInstr[nodeCurr].r
			}
			cnt += 1
		}
	}

	return cnt
}

func allArrived(nodes []string) bool {
	for _, node := range nodes {
		if string(node[2]) != "Z" {
			return false
		}
	}
	return true
}

func LeftRightGhost() int {
	sc, file := utils.Read("./day08/input.txt")
	// sc, file := utils.Read("./day08/input_2_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	sc.Scan()
	instruction := sc.Text()

	line := ""
	var tmpStrArr, tmpLR, nodes []string
	mapInstr := map[string]Instr{}

	for sc.Scan() {
		line = sc.Text()
		if len(line) == 0 {
			continue
		}

		tmpStrArr = strings.Split(line, " = ")
		tmpLR = strings.Split(tmpStrArr[1], ", ")
		mapInstr[tmpStrArr[0]] = Instr{strings.Trim(tmpLR[0], "("), strings.Trim(tmpLR[1], ")")}
		if string(tmpStrArr[0][2]) == "A" {
			nodes = append(nodes, tmpStrArr[0])
		}
	}

	cnt := 0
	total := []int{}

	for _, node := range nodes {
		currNode := node
		for string(currNode[2]) != "Z" {
			for _, instr := range instruction {
				if string(instr) == "L" {
					currNode = mapInstr[currNode].l
				} else {
					currNode = mapInstr[currNode].r
				}
				cnt += 1
			}
		}
		total = append(total, cnt)
		cnt = 0
	}

	return utils.LCM(total[0], total[1], total[2:]...)
}
