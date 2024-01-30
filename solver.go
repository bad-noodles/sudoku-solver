package main

import "fmt"

func sectionIndex(index, relativeIndex int) int {
	return index - index%3 + relativeIndex
}

func solve(input board) board {
	skipped := 0
	x := 0
	y := 0
	loop := 0

	advance := func() {
		x++
		if x == len(input[x]) {
			x = 0
			y++

			if y == len(input[x]) {
				y = 0
				loop++
				if skipped < 9*9 {
					skipped = 0
				}
			}
		}
	}

	for skipped < 9*9 {
		fmt.Println(loop)
		cell := input[x][y]

		if cell.value > 0 {
			skipped++
			advance()
			continue
		}

		candidates := cell.candidates
		cell.candidates = []int{}

		for _, candidate := range candidates {
			hasCandidate := false

			// Checking the line
			for _, checkCell := range input[x] {
				if checkCell.value == candidate {
					hasCandidate = true
					break
				}
			}

			if hasCandidate {
				continue
			}

			// Checking the row
			for checkX := range input[x] {
				if input[checkX][y].value == candidate {
					hasCandidate = true
					break
				}
			}

			if hasCandidate {
				continue
			}

			// Checking the section
			for sx := 0; sx < 3; sx++ {
				for sy := 0; sy < 3; sy++ {
					if input[sectionIndex(x, sx)][sectionIndex(y, sy)].value == candidate {
						hasCandidate = true
						break
					}
				}
			}

			if !hasCandidate {
				cell.candidates = append(cell.candidates, candidate)
			}
		}

		if len(cell.candidates) == 1 {
			cell.value = cell.candidates[0]
		}

		input[x][y] = cell
		advance()
	}
	return input
}
