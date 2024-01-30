package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type cell struct {
	value      int
	candidates []int
	cellNumber int
}

type board [9][9]cell

func (b board) String() string {
	var builder strings.Builder

	for y, line := range b {
		for x := range line {
			builder.WriteString(fmt.Sprintf("[%v] ", b[x][y].value))
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}

type sudoku struct {
	page *rod.Page
}

func newSudoku() sudoku {
	page := rod.New().ControlURL(
		launcher.New().Headless(false).MustLaunch(),
	).MustConnect().MustPage("https://www.nytimes.com/puzzles/sudoku/easy").MustWindowMaximize()
	page.MustWaitIdle()
	el := page.MustElementR("button", "Reject all")
	el.Click(proto.InputMouseButtonLeft, 1)
	return sudoku{page}
}

var defaultCandidates = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func (s sudoku) board() board {
	var b board
	elements := s.page.MustElements(".su-board [data-cell]")
	x := 0
	y := 0

	for i, el := range elements {
		inputValue := *el.MustAttribute("aria-label")

		if inputValue == "empty" {
			inputValue = "0"
		}
		value, err := strconv.Atoi(inputValue)

		if err != nil {
			panic(err)
		}

		b[x][y] = cell{
			value,
			defaultCandidates,
			i,
		}

		x++
		if x == len(b[y]) {
			x = 0
			y++
		}
	}

	return b
}

func (s sudoku) applySolution(b board) {
	for y, column := range b {
		for x, cell := range column {
			el := s.page.MustElement(fmt.Sprintf(".su-board [data-cell=\"%d\"]", cell.cellNumber))

			if x != 0 || y != 0 {
				el.Click(proto.InputMouseButtonLeft, 1)
			} else {
				el.Focus()
			}

			s.page.KeyActions().Press(input.ControlLeft).Type(input.Key(fmt.Sprintf("%v", cell.value)[0])).Do()
		}
	}
}
