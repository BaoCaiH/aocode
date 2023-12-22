package day14

import (
	"strings"

	"github.com/BaoCaiH/aocode/2023/utils"
)

func JK() int {
	sc, file := utils.Read("./day14/input.txt")
	// sc, file := utils.Read("./day14/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	platform := []string{}
	for sc.Scan() {
		platform = append(platform, sc.Text())
	}

	height := len(platform)
	stopper := []int{}
	for range platform[0] {
		stopper = append(stopper, -1)
	}
	total := 0
	for i, line := range platform {
		for j, c := range line {
			if c == '#' {
				stopper[j] = i
				continue
			}
			if c == '.' {
				continue
			}
			total += height - (stopper[j] + 1)
			stopper[j] += 1
		}
	}

	return total
}

func rollUp(platform []string) []string {
	// height := len(platform)
	stopper := []int{}
	for range platform[0] {
		stopper = append(stopper, 0)
	}
	newPlatform := []string{}
	for i, line := range platform {
		newPlatform = append(newPlatform, strings.ReplaceAll(strings.ReplaceAll(line, "#", "."), "O", "."))
		for j, c := range line {
			if c == '#' {
				stopper[j] = i + 1
				tmp := []rune(newPlatform[i])
				tmp[j] = '#'
				newPlatform[i] = string(tmp)
				continue
			}
			if c == '.' {
				continue
			}
			tmp := []rune(newPlatform[stopper[j]])
			tmp[j] = 'O'
			newPlatform[stopper[j]] = string(tmp)
			stopper[j] += 1
		}
	}

	return newPlatform
}

func turnRight(platform []string) []string {
	newPlatform := []string{}
	width := len(platform[0])
	height := len(platform)
	tmp := []rune{}
	for i := 0; i < width; i++ {
		tmp = []rune{}
		for j := height - 1; j >= 0; j-- {
			tmp = append(tmp, rune(platform[j][i]))
		}
		newPlatform = append(newPlatform, string(tmp))
	}
	return newPlatform
}

func stress(platform []string) int {
	total := 0
	height := len(platform)
	for i, line := range platform {
		for _, c := range line {
			if c == 'O' {
				total += height - i
			}
		}
	}
	return total
}

func TurningHead() int {
	sc, file := utils.Read("./day14/input.txt")
	// sc, file := utils.Read("./day14/input_1_example.txt")
	if sc == nil {
		return 0
	}
	defer file.Close()

	platform := []string{}
	for sc.Scan() {
		platform = append(platform, sc.Text())
	}

	sets := []int{}
	patterns := map[int][]int{}
	found := false
	for i := 0; i < 1000000000; i++ {
		// North
		platform = rollUp(platform)
		sets = append(sets, stress(platform))
		platform = turnRight(platform)
		// West
		platform = rollUp(platform)
		sets = append(sets, stress(platform))
		platform = turnRight(platform)
		// South
		platform = rollUp(platform)
		sets = append(sets, stress(platform))
		platform = turnRight(platform)
		// East
		platform = rollUp(platform)
		sets = append(sets, stress(platform))
		platform = turnRight(platform)

		if found {
			continue
		}
		for k, pattern := range patterns {
			checker := true
			for j := range pattern {
				if pattern[j] != sets[j] {
					checker = false
					break
				}
			}
			if checker {
				patternLength := i - k
				patternRepeat := (1000000000 - k) / patternLength
				// Jump to
				i = patternRepeat*patternLength + k
				found = true
				break
			}
		}
		patterns[i] = sets
		sets = []int{}
	}

	return stress(platform)
}
