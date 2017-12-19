package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle2() int {
	code := input()

	chan0 := make(chan int, 1000)
	chan1 := make(chan int, 1000)

	p0 := newProgram(0, code, chan1, chan0)
	p1 := newProgram(1, code, chan0, chan1)

	go p0.run(p1)
	p1.run(p0)

	// 127 is too low
	return p1.sndCount
}
