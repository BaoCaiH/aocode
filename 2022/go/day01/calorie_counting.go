package day01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func CountCalorie() int {
	// fmt.Println("Counting calories..")

	file, err := os.Open("./day01/input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	var curr, currMax int = 0, 0
	for sc.Scan() {
		i, err := strconv.Atoi(sc.Text())
		curr += i
		if err != nil { // Failed to convert the empty line to number, default return 0
			if curr > currMax {
				currMax = curr
			}
			curr = 0
		}
	}

	return currMax
}

func TotalTopThree() int {
	// fmt.Println("Getting top 3 Elves..")

	file, err := os.Open("./day01/input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	var curr int = 0
	topThree := []int{0, 0, 0}
	for sc.Scan() {
		i, err := strconv.Atoi(sc.Text())
		curr += i
		if err != nil { // Failed to convert the empty line to number, default return 0
			if curr > topThree[0] {
				topThree[0] = curr
				sort.Ints(topThree)
			}
			curr = 0
		}
	}

	// fmt.Printf("Top 3 are %v\n", topThree)

	sum := 0
	for _, num := range topThree {
		sum += num
	}

	return sum
}
