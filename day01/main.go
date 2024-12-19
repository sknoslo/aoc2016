package main

import (
	"fmt"
	"sknoslo/aoc2016/utils"
	"sknoslo/aoc2016/vec2"
	"strconv"
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
	dir := vec2.North
	pos := vec2.New(0, 0)

	for _, inst := range strings.Split(input, ", ") {
		var d rune
		var l int
		fmt.Sscanf(inst, "%c%d", &d, &l)

		if d == 'L' {
			dir = dir.RotateCardinalCCW()
		} else {
			dir = dir.RotateCardinalCW()
		}

		pos = pos.Add(dir.Mul(l))
	}

	return strconv.Itoa(utils.Abs(pos.X) + utils.Abs(pos.Y))
}

func parttwo() string {
	dir := vec2.North
	pos := vec2.New(0, 0)
	s := make(map[vec2.Vec2]bool, 256)

mainloop:
	for _, inst := range strings.Split(input, ", ") {
		var d rune
		var l int
		fmt.Sscanf(inst, "%c%d", &d, &l)

		if d == 'L' {
			dir = dir.RotateCardinalCCW()
		} else {
			dir = dir.RotateCardinalCW()
		}

		for range l {
			pos = pos.Add(dir)
			if s[pos] {
				break mainloop
			}
			s[pos] = true
		}
	}

	return strconv.Itoa(utils.Abs(pos.X) + utils.Abs(pos.Y))
}
