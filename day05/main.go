package main

import (
	"crypto/md5"
	"encoding/hex"
	"sknoslo/aoc2016/utils"
	"strconv"
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
	var pword [8]byte
	i := 0
	j := 0

	for j < 8 {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		sum := hex.EncodeToString(hash[:4])
		i++

		if sum[0:5] != "00000" {
			continue
		}

		pword[j] = sum[5]
		j++
	}
	return string(pword[:])
}

func parttwo() string {
	var pword [8]byte
	i := 0
	j := 0

	for j < 8 {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		i++

		sum := hex.EncodeToString(hash[:4])

		if sum[0:5] != "00000" {
			continue
		}

		k := sum[5] - '0'
		if k >= 0 && k < 8 && pword[k] == 0 {
			pword[k] = sum[6]
			j++
		}
	}
	return string(pword[:])
}
