package day12

import (
	"strconv"
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

type Setup struct {
	condition string
	check     []int
}

func isStillPossible(cond string, check []int) bool {
	cnt := 0
	setI := 0
	checkI := check[setI]
	sets := []int{}
	for _, c := range cond {
		if c == '#' {
			cnt += 1
			continue
		}
		if c == '?' && cnt <= checkI && len(sets) <= len(check) {
			return true
		}
		if c == '.' && cnt > 0 {
			if cnt != checkI {
				return false
			}
			sets = append(sets, cnt)
			cnt = 0
			if setI+1 == len(check) {
				checkI = -1
			} else {
				setI += 1
				checkI = check[setI]
			}
			continue
		}
	}
	if cnt > 0 {
		if cnt != checkI {
			return false
		}
		sets = append(sets, cnt)
	}
	if len(sets) != len(check) {
		return false
	}
	return true
}

func Spring() int {
	sc, file := utils.Read("./day12/input.txt")
	// sc, file := utils.Read("./day12/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	line := ""
	tmpStrParts := []string{}
	tmpIntParts := []int{}
	tmpInt := 0
	records := []Setup{}
	for sc.Scan() {
		line = sc.Text()
		tmpStrParts = strings.Split(line, " ")
		tmpIntParts = []int{}
		for _, part := range strings.Split(tmpStrParts[1], ",") {
			tmpInt, _ = strconv.Atoi(part)
			tmpIntParts = append(tmpIntParts, tmpInt)
		}
		records = append(records, Setup{tmpStrParts[0], tmpIntParts})
	}

	total := 0
	possibleRecords, tmpRecords := []Setup{}, []Setup{}
	for _, record := range records {
		for j, c := range record.condition {
			if j == 0 {
				possibleRecords = append(possibleRecords, record)
			}
			if c == '.' || c == '#' {
				continue
			}
			for _, pos := range possibleRecords {
				newCondition := pos.condition[:j] + "#" + pos.condition[j+1:]
				// Set #
				if isStillPossible(newCondition, pos.check) {
					tmpRecords = append(tmpRecords, Setup{newCondition, pos.check})
				}
				// Not set #
				newCondition = pos.condition[:j] + "." + pos.condition[j+1:]
				if isStillPossible(newCondition, pos.check) {
					tmpRecords = append(tmpRecords, Setup{newCondition, pos.check})
				}
			}
			possibleRecords = tmpRecords
			tmpRecords = []Setup{}
			// fmt.Println(possibleRecords)
		}
		total += len(possibleRecords)
		possibleRecords = []Setup{}
	}

	return total
}

type Pos struct {
	condI, checkI, block int
}

var tracking = map[Pos]int{}

func traverse(cond string, check []int, track Pos) int {
	// This is special
	// This is in case for `.??..??...?##. 1,1,3`
	// and we hit `.#...??...?##. 1,1,3` and `..#..??...?##. 1,1,3` at index 4
	// The results of these 2 possibilities are basically the same
	// since the problem now is only `..??...?##. 1,3`
	if n, ok := tracking[track]; ok {
		return n
	}

	// If at the end of string, exit condition
	if track.condI == len(cond) {
		if track.checkI == len(check) && track.block == 0 {
			return 1
		}
		if track.checkI == len(check)-1 && check[track.checkI] == track.block {
			return 1
		}
		return 0
	}

	tmp := 0
	if cond[track.condI] == '.' || cond[track.condI] == '?' {
		// If it's a full stop (or possible full stop)
		if track.block == 0 {
			// If there's no # block then just move the cursor
			tmp += traverse(cond, check, Pos{track.condI + 1, track.checkI, track.block})
		} else if track.block > 0 && track.checkI < len(check) && check[track.checkI] == track.block {
			// If there's a # block then stop the block, move the cursor for both the string and the checks
			tmp += traverse(cond, check, Pos{track.condI + 1, track.checkI + 1, 0})
		}
	}
	if cond[track.condI] == '#' || cond[track.condI] == '?' {
		// In case it's a # (or possible #), move the cursor and add 1 to # block
		tmp += traverse(cond, check, Pos{track.condI + 1, track.checkI, track.block + 1})
	}

	tracking[track] = tmp
	return tmp
}
func SpringRabbit() int {
	sc, file := utils.Read("./day12/input.txt")
	// sc, file := utils.Read("./day12/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	line := ""
	tmpStrParts := []string{}
	tmpCond := ""
	tmpCheck := ""
	tmpIntParts := []int{}
	tmpInt := 0
	records := []Setup{}
	for sc.Scan() {
		line = sc.Text()
		tmpStrParts = strings.Split(line, " ")
		tmpCond = tmpStrParts[0]
		tmpCheck = tmpStrParts[1]
		for i := 0; i < 4; i++ {
			tmpCond += "?" + tmpStrParts[0]
			tmpCheck += "," + tmpStrParts[1]
		}
		tmpStrParts = []string{tmpCond, tmpCheck}
		tmpIntParts = []int{}
		for _, part := range strings.Split(tmpStrParts[1], ",") {
			tmpInt, _ = strconv.Atoi(part)
			tmpIntParts = append(tmpIntParts, tmpInt)
		}
		records = append(records, Setup{tmpStrParts[0], tmpIntParts})
	}

	total := 0
	for _, record := range records {
		// fmt.Println(tracking)
		tracking = map[Pos]int{}
		n := traverse(record.condition, record.check, Pos{0, 0, 0})
		// fmt.Println(n)
		total += n
	}

	return total
}
