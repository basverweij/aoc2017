package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())
}

func puzzle1() int {
	v := newVirus(parse(rawInput))

	for i := 0; i < 10000; i++ {
		v.burst()
	}

	return v.infections
}
