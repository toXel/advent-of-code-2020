package day2

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var policyRegex = regexp.MustCompile(`(\d*)-(\d*) (\w): (.*)`)

func SolvePart1() int {
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
		charCount := strings.Count(pw, policy.CharToAppear)

		if charCount >= policy.MinCount && charCount <= policy.MaxCount {
			validPwCount++
		}
	}

	return validPwCount
}

func parsePolicy(value string) (policy, string) {
	parsed := policyRegex.FindStringSubmatch(value)
	min, _ := strconv.Atoi(parsed[1])
	max, _ := strconv.Atoi(parsed[2])
	policy := policy{min, max, parsed[3]}
	return policy, parsed[4]
}

type policy struct {
	MinCount     int
	MaxCount     int
	CharToAppear string
}
