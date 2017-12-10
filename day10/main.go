package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())
}

func puzzle1() int {
	l := knot(256, input)

	return l[0] * l[1]
}
