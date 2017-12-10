package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %s\n", puzzle2())
}

func puzzle1() int {
	l := knot(256, input)

	return l[0] * l[1]
}

func puzzle2() string {
	return hash(256, inputText)
}
