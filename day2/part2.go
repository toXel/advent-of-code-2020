package day2

import (
	"io/ioutil"
	"strings"
)

func SolvePart2() int {
	rawInput, err := ioutil.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n")

	var validPwCount int
	for _, row := range input {
		if row == "" {
			continue
		}

		policy, pw := parsePolicy(row)
		charOne := pw[policy.MinCount-1]
		charTwo := pw[policy.MaxCount-1]

		if (charOne == policy.CharToAppear[0]) != (charTwo == policy.CharToAppear[0]) {
			validPwCount++
		}
	}

	return validPwCount
}
