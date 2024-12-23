package main

import (
	"sknoslo/aoc2016/utils"
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
	lines := strings.Split(input, "\n")

	counts := make([][26]int, len(lines[0]))

	for _, line := range lines {
		for i, c := range line {
			counts[i][c-'a']++
		}
	}

	answer := []rune(lines[0])
	maxes := make([]int, len(answer))

	for i := range len(answer) {
		for j, c := range counts[i] {
			if c > maxes[i] {
				maxes[i] = c
				answer[i] = rune(j + 'a')
			}
		}
	}

	return string(answer)
}

func parttwo() string {
	lines := strings.Split(input, "\n")

	counts := make([][26]int, len(lines[0]))

	for _, line := range lines {
		for i, c := range line {
			counts[i][c-'a']++
		}
	}

	answer := []rune(lines[0])
	mins := make([]int, len(answer))

	for i := range len(answer) {
		for j, c := range counts[i] {
			if mins[i] == 0 || c < mins[i] {
				mins[i] = c
				answer[i] = rune(j + 'a')
			}
		}
	}

	return string(answer)
}
