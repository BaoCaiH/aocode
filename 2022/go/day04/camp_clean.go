package day04

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func TasksOverlapping() int {
	// fmt.Println("Checking tasks..")

	file, err := os.Open("./day04/input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	var tasks, tasks1, tasks2 []string
	var counter int
	for sc.Scan() {
		tasks = strings.Split(sc.Text(), ",")
		tasks1, tasks2 = strings.Split(tasks[0], "-"), strings.Split(tasks[1], "-")
		a, _ := strconv.Atoi(tasks1[0])
		b, _ := strconv.Atoi(tasks1[1])
		x, _ := strconv.Atoi(tasks2[0])
		y, _ := strconv.Atoi(tasks2[1])
		minLeft := math.Min(float64(a), float64(x))
		maxRight := math.Max(float64(b), float64(y))
		if (a == int(minLeft) && b == int(maxRight)) ||
			(x == int(minLeft) && y == int(maxRight)) {
			counter += 1
		}
	}

	return counter
}

func TasksOverlappingAbsolute() int {
	// fmt.Println("Checking tasks strictly..")

	file, err := os.Open("./day04/input.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	var tasks, tasks1, tasks2 []string
	var counter int
	for sc.Scan() {
		tasks = strings.Split(sc.Text(), ",")
		tasks1, tasks2 = strings.Split(tasks[0], "-"), strings.Split(tasks[1], "-")
		a, _ := strconv.Atoi(tasks1[0])
		b, _ := strconv.Atoi(tasks1[1])
		x, _ := strconv.Atoi(tasks2[0])
		y, _ := strconv.Atoi(tasks2[1])
		if !(a > y || x > b) {
			counter += 1
		}
	}

	return counter
}
