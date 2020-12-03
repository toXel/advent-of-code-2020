package day3

import (
	"io/ioutil"
	"strings"
)

func SolvePart2() int {
	rawInput, err := ioutil.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")

	stepConfigs := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result := 1
	for _, stepConfig := range stepConfigs {
		treeCount := countTrees(input, stepConfig[0], stepConfig[1])
		result *= treeCount
	}

	return result
}

func countTrees(input []string, colStepSize int, rowStepSize int) int {
	treeCount := 0
	column := 0
	for rowIndex := 0; rowIndex < len(input); rowIndex += rowStepSize {
		row := input[rowIndex]

		if row == "" || rowIndex == 0 {
			continue
		}

		column = (column + colStepSize) % len(row)

		char := row[column]

		if char == '#' {
			treeCount++
		}
	}
	return treeCount
}
