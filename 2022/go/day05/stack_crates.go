package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BaoCaiH/aocode/utils"
)

func readSetup() ([]utils.Stack, *bufio.Scanner, *os.File) {
	stacks := [9]utils.Stack{}

	file, err := os.Open("./day05/input.txt")
	if err != nil {
		fmt.Println(err)
		return stacks[:], nil, file
	}
	sc := bufio.NewScanner(file)

	var holder []string
	for i := 0; i < 9; i++ {
		if sc.Scan() {
			holder = append(holder, sc.Text())
		}
	}
	sc.Scan()

	for i := 7; i >= 0; i-- {
		tmp := holder[i]
		stacks[0].Push(string(tmp[1]))
		stacks[1].Push(string(tmp[5]))
		stacks[2].Push(string(tmp[9]))
		stacks[3].Push(string(tmp[13]))
		stacks[4].Push(string(tmp[17]))
		stacks[5].Push(string(tmp[21]))
		stacks[6].Push(string(tmp[25]))
		stacks[7].Push(string(tmp[29]))
		stacks[8].Push(string(tmp[33]))
	}

	return stacks[:], sc, file
}

func StackCrate() string {
	// fmt.Println("Stacking crates..")
	stacks, sc, file := readSetup()
	defer file.Close()
	for sc.Scan() {
		line := sc.Text()
		count, _ := strconv.Atoi(line[strings.Index(line, "move")+5 : strings.Index(line, "from")-1])
		from, _ := strconv.Atoi(line[strings.Index(line, "from")+5 : strings.Index(line, "to")-1])
		to, _ := strconv.Atoi(line[strings.Index(line, "to")+3:])
		for i := 0; i < count; i++ {
			stacks[to-1].Push(stacks[from-1].Pop())
		}
	}

	topCrates := ""
	for _, stack := range stacks {
		topCrates += stack.Pop()
	}

	return topCrates
}

func StackCrateFast() string {
	// fmt.Println("Stacking crates faster..")
	stacks, sc, file := readSetup()
	defer file.Close()
	for sc.Scan() {
		line := sc.Text()
		count, _ := strconv.Atoi(line[strings.Index(line, "move")+5 : strings.Index(line, "from")-1])
		from, _ := strconv.Atoi(line[strings.Index(line, "from")+5 : strings.Index(line, "to")-1])
		to, _ := strconv.Atoi(line[strings.Index(line, "to")+3:])
		stacks[to-1].PushMultiple(stacks[from-1].PopMultiple(count))
	}

	topCrates := ""
	for _, stack := range stacks {
		topCrates += stack.Pop()
	}

	return topCrates
}
