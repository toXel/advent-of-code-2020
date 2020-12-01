package day1

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func SolvePart2() int {
	rawInput, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")

	var result int

	for _, first := range input {
		firstInt, _ := strconv.Atoi(first)

		for _, second := range input {
			secondInt, _ := strconv.Atoi(second)

			for _, third := range input {
				thirdInt, _ := strconv.Atoi(third)

				if firstInt+secondInt+thirdInt == 2020 {
					result = firstInt * secondInt * thirdInt
					break
				}

				if result > 0 {
					break
				}
			}

			if result > 0 {
				break
			}
		}

		if result > 0 {
			break
		}
	}

	return result
}
