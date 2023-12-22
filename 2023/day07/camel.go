package day07

import (
	"slices"
	"strconv"
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

type Hand struct {
	cards string
	bid   int
}

var CardStrengths = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
var CardStrengthsJoker = []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}

func (h Hand) GetType() int {
	tmpMap := map[string]int{}

	for _, c := range h.cards {
		if _, ok := tmpMap[string(c)]; !ok {
			tmpMap[string(c)] = 0
		}
		tmpMap[string(c)] += 1
	}

	if len(tmpMap) == 1 {
		return 0
	}
	if len(tmpMap) == 2 {
		for _, v := range tmpMap {
			if v == 1 || v == 4 {
				return 1
			}
		}
		return 2
	}
	if len(tmpMap) == 3 {
		for _, v := range tmpMap {
			if v == 3 {
				return 3
			}
		}
		return 4
	}
	if len(tmpMap) == 4 {
		return 5
	}
	return 6
}

func (h Hand) GetTypeJoker() int {
	tmpMap := map[string]int{}
	joker := 0

	for _, c := range h.cards {
		if string(c) == "J" {
			joker += 1
			continue
		}
		if _, ok := tmpMap[string(c)]; !ok {
			tmpMap[string(c)] = 0
		}
		tmpMap[string(c)] += 1
	}

	if joker > 0 {
		mostCard := ""
		cnt := 0
		for c, k := range tmpMap {
			if k > cnt {
				mostCard = c
				cnt = k
			}
		}
		tmpMap[mostCard] += joker
	}

	if len(tmpMap) == 1 {
		return 0
	}
	if len(tmpMap) == 2 {
		for _, v := range tmpMap {
			if v == 1 || v == 4 {
				return 1
			}
		}
		return 2
	}
	if len(tmpMap) == 3 {
		for _, v := range tmpMap {
			if v == 3 {
				return 3
			}
		}
		return 4
	}
	if len(tmpMap) == 4 {
		return 5
	}
	return 6
}

func Compare(h, k Hand) int {
	if h.GetType() < k.GetType() {
		return 1
	} else if h.GetType() > k.GetType() {
		return -1
	}

	for i := 0; i < 5; i++ {
		x := utils.IndexOf(string(h.cards[i]), CardStrengths)
		y := utils.IndexOf(string(k.cards[i]), CardStrengths)
		if x < y {
			return 1
		} else if x > y {
			return -1
		}
	}

	return 0
}

func CompareJoker(h, k Hand) int {
	if h.GetTypeJoker() < k.GetTypeJoker() {
		return 1
	} else if h.GetTypeJoker() > k.GetTypeJoker() {
		return -1
	}

	for i := 0; i < 5; i++ {
		x := utils.IndexOf(string(h.cards[i]), CardStrengthsJoker)
		y := utils.IndexOf(string(k.cards[i]), CardStrengthsJoker)
		if x < y {
			return 1
		} else if x > y {
			return -1
		}
	}

	return 0
}

func Camel() int {
	sc, file := utils.Read("./day07/input.txt")
	// sc, file := utils.Read("./day07/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	var hands []Hand
	var tmpStrParts []string
	var tmpInt int
	line := ""

	for sc.Scan() {
		line = sc.Text()
		tmpStrParts = strings.Split(line, " ")
		tmpInt, _ = strconv.Atoi(tmpStrParts[1])
		hands = append(hands, Hand{tmpStrParts[0], tmpInt})
	}

	slices.SortFunc(hands, Compare)

	total := 0

	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}

	return total
}

func CamelJoker() int {
	sc, file := utils.Read("./day07/input.txt")
	// sc, file := utils.Read("./day07/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	var hands []Hand
	var tmpStrParts []string
	var tmpInt int
	line := ""

	for sc.Scan() {
		line = sc.Text()
		tmpStrParts = strings.Split(line, " ")
		tmpInt, _ = strconv.Atoi(tmpStrParts[1])
		hands = append(hands, Hand{tmpStrParts[0], tmpInt})

	}

	slices.SortFunc(hands, CompareJoker)

	total := 0

	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}

	return total
}
