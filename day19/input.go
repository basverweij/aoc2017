package main

import (
	"io/ioutil"
	"strings"
)

func input() [][]byte {
	b, _ := ioutil.ReadFile("input.txt")
	return parse(string(b))
}

func parse(input string) [][]byte {
	lines := strings.Split(input, "\r\n")
	tubes := make([][]byte, 0)

	for _, line := range lines {
		if line == "" {
			break
		}

		tubes = append(tubes, []byte(line))
	}

	return tubes
}
