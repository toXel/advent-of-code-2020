package main

import (
	"fmt"

	"github.com/toXel/advent-of-code-2020/day1"
	"github.com/toXel/advent-of-code-2020/day2"
	"github.com/toXel/advent-of-code-2020/day3"
	"github.com/toXel/advent-of-code-2020/day4"
)

func main() {
	println("Day 1 Part 1: " + fmt.Sprint(day1.SolvePart1()))
	println("Day 1 Part 2: " + fmt.Sprint(day1.SolvePart2()))
	println("Day 2 Part 1: " + fmt.Sprint(day2.SolvePart1()))
	println("Day 2 Part 2: " + fmt.Sprint(day2.SolvePart2()))
	println("Day 3 Part 1: " + fmt.Sprint(day3.SolvePart1()))
	println("Day 3 Part 2: " + fmt.Sprint(day3.SolvePart2()))
	println("Day 4 Part 1: " + fmt.Sprint(day4.SolvePart1()))
	println("Day 4 Part 2: " + fmt.Sprint(day4.SolvePart2()))
}
