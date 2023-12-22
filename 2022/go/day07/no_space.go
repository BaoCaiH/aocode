package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type path struct {
	name     string
	parent   *path
	size     int
	isFile   bool
	contains map[string]*path
}

func traverse() (*path, []*path) {
	file, err := os.Open("./day07/input.txt")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	tree := path{name: "/", contains: map[string]*path{}}
	dirs := []*path{&tree}
	var current *path
	for sc.Scan() {
		line := sc.Text()
		if line == "$ cd /" {
			current = &tree
		} else if line == "$ cd .." {
			if current.parent != nil {
				current = current.parent
			}
		} else if line[:4] == "$ cd" {
			current = current.contains[line[5:]]
		} else if line[:3] == "dir" {
			newDir := path{name: line[4:], parent: current, contains: map[string]*path{}}
			current.contains[line[4:]] = &newDir
			dirs = append(dirs, &newDir)
		} else if line == "$ ls" {
			continue
		} else {
			stuffs := strings.Split(line, " ")
			s, _ := strconv.Atoi(stuffs[0])
			current.contains[stuffs[1]] = &path{name: stuffs[1], size: s, isFile: true, parent: current}
		}
	}

	setOrSetDirSize(&tree)

	return &tree, dirs
}

func setOrSetDirSize(dir *path) int {
	if dir.size != 0 {
		return dir.size
	}

	totalSize := 0
	for _, f := range dir.contains {
		if f.isFile {
			totalSize += f.size
		} else {
			totalSize += setOrSetDirSize(f)
		}
	}

	dir.size = totalSize
	return totalSize
}

func NoSpace() int {
	_, dirs := traverse()

	totalSize := 0
	for _, dir := range dirs {
		// fmt.Printf("Dir %s size %d\n", dir.name, dir.size)
		if dir.size <= 100000 {
			totalSize += dir.size
		}
	}

	return totalSize
}

func FreeUp() int {
	root, dirs := traverse()

	neededSize := 30000000 - (70000000 - root.size)
	currMin := 30000000

	for _, dir := range dirs {
		if dir.size >= neededSize && dir.size < currMin {
			currMin = dir.size
		}
	}

	return currMin
}
