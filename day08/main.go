package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())
}

func puzzle1() int {
	f, _ := os.Open("input.txt")
	is, _ := parse(f)

	regs := make(registers)

	for _, i := range is {
		regs.apply(i)
	}

	return regs.largestValue()
}
