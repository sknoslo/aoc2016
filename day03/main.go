package main

import (
	"sknoslo/aoc2016/utils"
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

func isTri(a, b, c int) bool {
	return a + b > c && a + c > b && c + b > a
}

func partone() string {
	valid := 0
	for _, line := range strings.Split(input, "\n") {
		t := utils.MustSplitInts(line)

		if isTri(t[0], t[1], t[2]){
			valid++
		}
	}
	return strconv.Itoa(valid)
}

func parttwo() string {
	valid := 0
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i += 3 {
		a := utils.MustSplitInts(lines[i])
		b := utils.MustSplitInts(lines[i+1])
		c := utils.MustSplitInts(lines[i+2])

		for j := range 3 {
			if isTri(a[j], b[j], c[j]) {
				valid++
			}
		}
	}

	return strconv.Itoa(valid)
}
