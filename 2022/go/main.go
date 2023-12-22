package main

import (
	"fmt"

	"github.com/BaoCaiH/aocode/day01"
	"github.com/BaoCaiH/aocode/day02"
	"github.com/BaoCaiH/aocode/day03"
	"github.com/BaoCaiH/aocode/day04"
	"github.com/BaoCaiH/aocode/day05"
	"github.com/BaoCaiH/aocode/day06"
	"github.com/BaoCaiH/aocode/day07"
	"github.com/BaoCaiH/aocode/day08"
	"github.com/BaoCaiH/aocode/day09"
	"github.com/BaoCaiH/aocode/day10"
	"github.com/BaoCaiH/aocode/day11"
	"github.com/BaoCaiH/aocode/day12"
)

func main() {
	fmt.Println("Hello beautiful people..")

	fmt.Printf("Day 1 - Part 1: %d\n", day01.CountCalorie())
	fmt.Printf("Day 1 - Part 2: %d\n", day01.TotalTopThree())
	fmt.Printf("Day 2 - Part 1: %d\n", day02.RockScissorsPaper())
	fmt.Printf("Day 2 - Part 2: %d\n", day02.RockScissorsPaperSneaky())
	fmt.Printf("Day 3 - Part 1: %d\n", day03.RucksackReorg())
	fmt.Printf("Day 3 - Part 2: %d\n", day03.RucksackGroup())
	fmt.Printf("Day 4 - Part 1: %d\n", day04.TasksOverlapping())
	fmt.Printf("Day 4 - Part 2: %d\n", day04.TasksOverlappingAbsolute())
	fmt.Printf("Day 5 - Part 1: %s\n", day05.StackCrate())
	fmt.Printf("Day 5 - Part 2: %s\n", day05.StackCrateFast())
	fmt.Printf("Day 6 - Part 1: %d\n", day06.Tuning(4))
	fmt.Printf("Day 6 - Part 2: %d\n", day06.Tuning(14))
	fmt.Printf("Day 7 - Part 1: %d\n", day07.NoSpace())
	fmt.Printf("Day 7 - Part 2: %d\n", day07.FreeUp())
	fmt.Printf("Day 8 - Part 1: %d\n", day08.TreetopVisible())
	fmt.Printf("Day 8 - Part 2: %d\n", day08.TreetopScenic())
	fmt.Printf("Day 9 - Part 1: %d\n", day09.RopeBridge())
	fmt.Printf("Day 9 - Part 2: %d\n", day09.RopeBridgeSnap())
	fmt.Printf("Day 10 - Part 1: %d\n", day10.CathodeRayTube())
	fmt.Printf("Day 10 - Part 2:\n")
	day10.CathodeRayTubePixel()
	fmt.Printf("Day 11 - Part 1: %d\n", day11.MonkeyBusiness())
	fmt.Printf("Day 11 - Part 2: %d\n", day11.MonkeyBusinessChaos())
	fmt.Printf("Day 12 - Part 1: %d\n", day12.HillClimb())
	fmt.Printf("Day 12 - Part 2: %d\n", day12.HillClimbCheat())
}
