package day6

import (
	"strings"

	"github.com/toXel/advent-of-code-2020/utils"
)

func SolvePart2() int {
	groups := utils.ReadInput("day6/input.txt", "\n\n")

	sum := 0
	for _, group := range groups {
		group = strings.TrimSpace(group)
		if group == "" {
			continue
		}

		personCount := len(strings.Split(group, "\n"))

		group = strings.ReplaceAll(group, "\n", "")

		uniqueMap := make(map[rune]struct{})
		for _, char := range group {
			if _, ok := uniqueMap[char]; !ok {
				uniqueMap[char] = struct{}{}
			}
		}

		uniqueAnswers := make([]rune, 0)
		for key := range uniqueMap {
			uniqueAnswers = append(uniqueAnswers, key)
		}

		answersCount := 0
		for _, answer := range uniqueAnswers {
			if strings.Count(group, string(answer)) == personCount {
				answersCount++
			}
		}

		sum += answersCount
	}

	return sum
}
