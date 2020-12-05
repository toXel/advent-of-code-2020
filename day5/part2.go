package day5

import (
	"sort"

	"github.com/toXel/advent-of-code-2020/utils"
)

func SolvePart2() int {
	input := utils.ReadInput("day5/input.txt", "\n")

	seatIDs := make([]int, 0)
	for _, seat := range input {
		if seat == "" {
			continue
		}

		lowRow := 0
		highRow := 127
		lowCol := 0
		highCol := 7

		// Calculate row
		for i := 0; i <= 6; i++ {
			half := seat[i]

			switch half {
			case 'F':
				highRow = ((highRow - lowRow) / 2) + lowRow
			case 'B':
				lowRow = highRow - ((highRow - lowRow) / 2)
			}
		}

		// Calculate column
		for i := 7; i <= 9; i++ {
			half := seat[i]

			switch half {
			case 'L':
				highCol = ((highCol - lowCol) / 2) + lowCol
			case 'R':
				lowCol = highCol - ((highCol - lowCol) / 2)
			}
		}

		seatID := highRow*8 + highCol

		seatIDs = append(seatIDs, seatID)
	}

	sort.Ints(seatIDs)

	for i := 0; i <= len(seatIDs); i++ {
		if seatIDs[i+1] != seatIDs[i]+1 {
			return seatIDs[i] + 1
		}
	}

	return 0
}
