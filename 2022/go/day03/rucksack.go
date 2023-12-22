package day03

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RucksackReorg() int {
	// fmt.Println("Sorting items..")

	file, err := os.Open("./day03/input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	var commonItem, contents, com1, com2 string
	var prio int = 0
	for sc.Scan() {
		commonItem = ""
		contents = sc.Text()
		com1, com2 = contents[:len(contents)/2], contents[len(contents)/2:]
		for _, char := range com1 {
			if strings.Contains(com2, string(char)) && !strings.Contains(commonItem, string(char)) {
				commonItem += string(char)
				if char > 96 {
					prio += (int(char) - 96)
				} else {
					prio += (int(char) - 38)
				}
			}
		}
	}
	return prio
}
func scanThree(sc *bufio.Scanner) (bool, []string) {
	var counter int = 0
	var bags [3]string
	for sc.Scan() {
		bags[counter] = sc.Text()
		counter += 1
		if counter > 2 {
			break
		}
	}
	if counter != 3 {
		return false, nil
	}
	return true, bags[:]
}

func RucksackGroup() int {
	// fmt.Println("Checking groups..")

	file, err := os.Open("./day03/input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	var commonItem, bag1, bag2, bag3 string
	var prio int = 0
	cont, bags := scanThree(sc)
	for cont {
		commonItem = ""
		bag1, bag2, bag3 = bags[0], bags[1], bags[2]
		for _, char := range bag1 {
			if strings.Contains(bag2, string(char)) &&
				strings.Contains(bag3, string(char)) &&
				!strings.Contains(commonItem, string(char)) {
				commonItem += string(char)
				if char > 96 {
					prio += (int(char) - 96)
				} else {
					prio += (int(char) - 38)
				}
			}
		}
		cont, bags = scanThree(sc)
	}
	return prio
}
