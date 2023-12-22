package day02

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const elfChoices string = "ABC"
const myChoices string = "XYZ"

func RockScissorsPaper() int {
	// fmt.Println("Play rock scissors paper..")

	file, err := os.Open("./day02/input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	var total, score int = 0, 0
	for sc.Scan() {
		choices := strings.Split(sc.Text(), " ")
		elf, me := strings.Index(elfChoices, choices[0]), strings.Index(myChoices, choices[1])
		switch me - elf {
		case 0:
			score = me + 4
		case 1, -2:
			score = me + 7
		default:
			score = me + 1
		}
		total += score
	}

	return total
}

func RockScissorsPaperSneaky() int {
	// fmt.Println("Play rock scissors paper easy mode..")

	file, err := os.Open("./day02/input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	var total, score int = 0, 0
	for sc.Scan() {
		choices := strings.Split(sc.Text(), " ")
		elf, me := strings.Index(elfChoices, choices[0]), strings.Index(myChoices, choices[1])
		switch me {
		case 0:
			score = ((elf + 2) % 3) + 1
		case 1:
			score = elf + 4
		case 2:
			score = ((elf + 1) % 3) + 7
		default:
			score = 0
		}
		total += score
	}

	return total
}
