package day1

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func SolvePart1() int {
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

			if firstInt+secondInt == 2020 {
				result = firstInt * secondInt
				break
			}
		}
	}

	return result
}
