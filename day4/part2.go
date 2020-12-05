package day4

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func SolvePart2() int {
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

		if len(keyValues) == 8 && validatePassport(keyValues) {
			validPassports++
		} else if len(keyValues) == 7 && !contains(keys, "cid") && validatePassport(keyValues) {
			validPassports++
		}
	}

	return validPassports
}

var colorHexRegex = regexp.MustCompile(`^#[a-f0-9]{6}$`)
var pidRegex = regexp.MustCompile(`^\d{9}$`)
var eclList = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func validatePassport(data map[string]string) bool {
	byr := data["byr"]
	if len(byr) == 4 {
		value, _ := strconv.Atoi(byr)
		if value < 1920 || value > 2002 {
			return false
		}
	} else {
		return false
	}

	iyr := data["iyr"]
	if len(iyr) == 4 {
		value, _ := strconv.Atoi(iyr)
		if value < 2010 || value > 2020 {
			return false
		}
	} else {
		return false
	}

	eyr := data["eyr"]
	if len(eyr) == 4 {
		value, _ := strconv.Atoi(eyr)
		if value < 2020 || value > 2030 {
			return false
		}
	} else {
		return false
	}

	hgt := data["hgt"]
	if strings.HasSuffix(hgt, "cm") || strings.HasSuffix(hgt, "in") {
		unit := hgt[len(hgt)-2:]
		value, _ := strconv.Atoi(strings.Split(hgt, unit)[0])

		if unit == "cm" && (value < 150 || value > 193) {
			return false
		} else if unit == "in" && (value < 59 || value > 76) {
			return false
		}
	} else {
		return false
	}

	hcl := data["hcl"]
	if len(hcl) == 7 {
		if !colorHexRegex.MatchString(hcl) {
			return false
		}
	} else {
		return false
	}

	ecl := data["ecl"]
	if !contains(eclList, ecl) {
		return false
	}

	pid := data["pid"]
	if !pidRegex.MatchString(pid) {
		return false
	}

	return true
}
