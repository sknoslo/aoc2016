package main

import (
	"sknoslo/aoc2016/utils"
	"sort"
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

type CharCount struct {
	char  rune
	count int
}

func partone() string {
	sum := 0

mainloop:
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")

		name := strings.Join(parts[0:len(parts)-1], "")
		parts = strings.Split(parts[len(parts)-1], "[")
		sectorId := utils.MustAtoi(parts[0])
		checksum := parts[1][0 : len(parts[1])-1]

		counts := make([]CharCount, 26)

		for _, c := range name {
			counts[c-'a'].char = c
			counts[c-'a'].count++
		}

		sort.Slice(counts, func(i, j int) bool {
			ic, jc := counts[i], counts[j]
			return ic.count > jc.count || ic.count == jc.count && ic.char < jc.char
		})

		for i, c := range checksum {
			if counts[i].char != c {
				continue mainloop
			}
		}

		sum += sectorId
	}
	return strconv.Itoa(sum)
}

func parttwo() string {
mainloop:
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")

		name := strings.Join(parts[0:len(parts)-1], "-")
		parts = strings.Split(parts[len(parts)-1], "[")
		sectorId := utils.MustAtoi(parts[0])
		checksum := parts[1][0 : len(parts[1])-1]

		counts := make([]CharCount, 26)

		for _, c := range name {
			if c == '-' {
				continue
			}
			counts[c-'a'].char = c
			counts[c-'a'].count++
		}

		sort.Slice(counts, func(i, j int) bool {
			ic, jc := counts[i], counts[j]
			return ic.count > jc.count || ic.count == jc.count && ic.char < jc.char
		})

		for i, c := range checksum {
			if counts[i].char != c {
				continue mainloop
			}
		}

		decrypt := make([]rune, len(name))

		for i, c := range name {
			if c == '-' {
				decrypt[i] = ' '
			} else {
				decrypt[i] = rune((int(c - 'a') + sectorId) % 26 + int('a'))
			}
		}

		name = string(decrypt)

		if name == "northpole object storage" {
			return strconv.Itoa(sectorId)
		}
	}
	return "not found"
}
