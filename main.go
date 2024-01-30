package main

import (
	"time"
)

func main() {
	s := newSudoku()
	solution := solve(s.board())

	s.applySolution(solution)

	time.Sleep(10 * time.Second)
}
