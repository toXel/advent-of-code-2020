package day3

import (
	"io/ioutil"
	"strings"
)

func SolvePart1() int {
	rawInput, err := ioutil.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")

	treeCount := 0
	column := 0
	for index, row := range input {
		if row == "" || index == 0 {
			continue
		}

		column = (column + 3) % len(row)

		char := row[column]

		if char == '#' {
			treeCount++
		}
	}

	return treeCount
}
