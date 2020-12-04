package day4

import (
	"io/ioutil"
	"strings"
)

func SolvePart1() int {
	rawInput, err := ioutil.ReadFile("day4/input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(rawInput), "\n\n")

	validPassports := 0
	for _, passport := range input {
		keyValues := parsePassport(passport)

		keys := make([]string, 0)
		for key, _ := range keyValues {
			keys = append(keys, key)
		}

		if len(keyValues) == 8 {
			validPassports++
		} else if len(keyValues) == 7 && !contains(keys, "cid") {
			validPassports++
		}
	}

	return validPassports
}

func parsePassport(passport string) map[string]string {
	result := make(map[string]string)
	passport = strings.TrimSpace(strings.ReplaceAll(passport, "\n", " "))

	for _, pair := range strings.Split(passport, " ") {
		splitted := strings.Split(pair, ":")
		result[splitted[0]] = splitted[1]
	}

	return result
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
