package main

import (
	"fmt"

	"github.com/basverweij/aoc2017/day10"
)

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %s\n", puzzle2())
}

func puzzle1() int {
	l := day10.Knot(256, input)

	return l[0] * l[1]
}

func puzzle2() string {
	return day10.Hash(256, inputText)
}
