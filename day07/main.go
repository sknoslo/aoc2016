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

func isABBA(s string) bool {
	return s[0] == s[3] && s[1] == s[2] && s[0] != s[1]
}

func partone() string {
	ips := strings.Split(input, "\n")
	tlsIps := 0

mainloop:
	for _, ip := range ips {
		inBrackets := false
		tlsIp := false
		for i, c := range ip {
			switch c {
			case '[':
				inBrackets = true
			case ']':
				inBrackets = false
			default:
				if i+3 < len(ip) {
					s := ip[i : i+4]
					if !strings.ContainsAny(s, "[]") && isABBA(s) {
						if inBrackets {
							tlsIp = false
							continue mainloop
						} else {
							tlsIp = true
						}
					}
				}
			}
		}
		if tlsIp {
			tlsIps++
		}
	}

	return strconv.Itoa(tlsIps)
}

func isABA(s string) bool {
	return s[0] == s[2] && s[0] != s[1]
}

func parttwo() string {
	ips := strings.Split(input, "\n")
	sslIps := 0

mainloop:
	for _, ip := range ips {
		abas, babs := make([]string, 0, 10), make([]string, 0, 10)
		inBrackets := false
		for i, c := range ip {
			switch c {
			case '[':
				inBrackets = true
			case ']':
				inBrackets = false
			default:
				if i+2 < len(ip) {
					s := ip[i : i+3]
					if !strings.ContainsAny(s, "[]") && isABA(s) {
						if inBrackets {
							babs = append(babs, s)
						} else {
							abas = append(abas, s)
						}
					}
				}
			}
		}

		for _, aba := range abas {
			for _, bab := range babs {
				if aba[0] == bab[1] && bab[0] == aba[1] {
					sslIps++
					continue mainloop
				}
			}
		}
	}

	return strconv.Itoa(sslIps)
}
