package day6

import (
	"strings"

	"github.com/toXel/advent-of-code-2020/utils"
)

func SolvePart1() int {
	groups := utils.ReadInput("day6/input.txt", "\n\n")

	sum := 0
	for _, group := range groups {
		group = strings.ReplaceAll(group, "\n", "")
		group = strings.TrimSpace(group)

		if group == "" {
			continue
		}

		uniqueMap := make(map[rune]struct{})
		for _, char := range group {
			if _, ok := uniqueMap[char]; !ok {
				uniqueMap[char] = struct{}{}
			}
		}

		result := make([]rune, 0)
		for key := range uniqueMap {
			result = append(result, key)
		}

		sum += len(result)
	}

	return sum
}
