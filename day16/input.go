package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func rawInput() []string {
	f, _ := os.Open("input.txt")
	b, _ := ioutil.ReadAll(f)
	return strings.Split(string(b), ",")
}

func input() []move {
	r := rawInput()
	moves := make([]move, len(r))

	var m move
	for i, s := range r {
		if s[0] == 's' {
			// spin
			m = newSpin(toInt(s[1:]))
		} else if s[0] == 'x' {
			// exchange
			j := strings.Index(s, "/")
			m = newExchange(toInt(s[1:j]), toInt(s[j+1:]))
		} else {
			// partner
			m = newPartner(s[1], s[3])
		}

		moves[i] = m
	}

	return moves
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
