package day11

import (
	"strconv"
	"strings"

	"github.com/BaoCaiH/aocode/utils"
)

type monkey struct {
	items     []int
	left      string
	operation string
	right     string
	test      int
	if_true   int
	if_false  int
	inspects  int
}

func parseMonkeys() []monkey {
	// sc, file := utils.Read("day11/example.txt")
	sc, file := utils.Read("day11/input.txt")
	if sc == nil {
		return nil
	}
	defer file.Close()

	monkeys := []monkey{}
	i := -1
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, "Monkey") {
			i++
			monkeys = append(monkeys, monkey{})
		} else if strings.Contains(line, "Starting") {
			monkeys[i].items = apply(strings.Split(line[18:], ", "), toInt)
		} else if strings.Contains(line, "Operation") {
			operations := strings.Split(line[19:], " ")
			monkeys[i].left = operations[0]
			monkeys[i].operation = operations[1]
			monkeys[i].right = operations[2]
		} else if strings.Contains(line, "Test") {
			monkeys[i].test = toInt(line[21:])
		} else if strings.Contains(line, "If true") {
			monkeys[i].if_true = toInt(line[29:])
		} else if strings.Contains(line, "If false") {
			monkeys[i].if_false = toInt(line[30:])
		}
	}
	return monkeys
}

func MonkeyBusiness() int {
	monkeys := parseMonkeys()
	if monkeys == nil {
		return 0
	}

	for i := 0; i < 20; i++ {
		for j, monkey := range monkeys {
			for _, item := range monkey.items {
				newItem, left, right := 0, 0, 0

				if monkey.left == "old" {
					left = item
				} else {
					left = toInt(monkey.left)
				}
				if monkey.right == "old" {
					right = item
				} else {
					right = toInt(monkey.right)
				}

				if monkey.operation == "+" {
					newItem = (left + right) / 3
				} else {
					newItem = (left * right) / 3
				}

				if newItem%monkey.test == 0 {
					monkeys[monkey.if_true].items = append(monkeys[monkey.if_true].items, newItem)
				} else {
					monkeys[monkey.if_false].items = append(monkeys[monkey.if_false].items, newItem)
				}

				monkeys[j].inspects++
			}

			monkeys[j].items = []int{}
		}
	}

	most, second := 0, 0
	for _, m := range monkeys {
		if m.inspects > most {
			second = most
			most = m.inspects
		} else if m.inspects > second {
			second = m.inspects
		}
	}
	return most * second
}
func MonkeyBusinessChaos() int {
	monkeys := parseMonkeys()
	if monkeys == nil {
		return 0
	}

	manage := 1
	for _, m := range monkeys {
		manage *= m.test
	}

	for i := 0; i < 10000; i++ {
		for j, monkey := range monkeys {
			for _, item := range monkey.items {
				newItem, left, right := 0, 0, 0

				if monkey.left == "old" {
					left = item
				} else {
					left = toInt(monkey.left)
				}
				if monkey.right == "old" {
					right = item
				} else {
					right = toInt(monkey.right)
				}

				if monkey.operation == "+" {
					newItem = (left + right) % manage
				} else {
					newItem = (left * right) % manage
				}

				if newItem%monkey.test == 0 {
					monkeys[monkey.if_true].items = append(monkeys[monkey.if_true].items, newItem)
				} else {
					monkeys[monkey.if_false].items = append(monkeys[monkey.if_false].items, newItem)
				}

				monkeys[j].inspects++
			}

			monkeys[j].items = []int{}
		}
	}

	most, second := 0, 0
	for _, m := range monkeys {
		if m.inspects > most {
			second = most
			most = m.inspects
		} else if m.inspects > second {
			second = m.inspects
		}
	}
	return most * second
}
func apply[T comparable, V comparable](things []T, f func(T) V) []V {
	returns := make([]V, len(things))
	for i, thing := range things {
		returns[i] = f(thing)
	}
	return returns
}

func toInt(s string) int {
	d, _ := strconv.Atoi(s)
	return d
}
