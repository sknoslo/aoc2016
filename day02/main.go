package main

import (
	"sknoslo/aoc2016/grids"
	"sknoslo/aoc2016/utils"
	"sknoslo/aoc2016/vec2"
	"strings"
)

var input string

func init() {
	input = utils.MustReadInput("input.txt")
}

func main() {
	utils.Run(1, partone)
	utils.Run(2, parttwo)
}

func partone() string {
	pad := grids.New(3, 3, []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'})
	pos := vec2.New(1, 1)
	lines := strings.Split(input, "\n")
	res := make([]byte, len(lines))

	for i, line := range lines {
		for _, inst := range line {
			dir := vec2.FromRune(inst)
			n := pos.Add(dir)
			if pad.InGrid(n) {
				pos = n
			}
		}
		res[i] = pad.CellAt(pos)
	}

	return string(res)
}

func parttwo() string {
	pad := grids.New(5, 5, []byte{
		'_', '_', '1', '_', '_',
		'_', '2', '3', '4', '_',
		'5', '6', '7', '8', '9',
		'_', 'A', 'B', 'C', '_',
		'_', '_', 'D', '_', '_',
	})
	pos := vec2.New(1, 1)
	lines := strings.Split(input, "\n")
	res := make([]byte, len(lines))

	for i, line := range lines {
		for _, inst := range line {
			dir := vec2.FromRune(inst)
			n := pos.Add(dir)
			if pad.InGrid(n) && pad.CellAt(n) != '_' {
				pos = n
			}
		}
		res[i] = pad.CellAt(pos)
	}

	return string(res)
}
