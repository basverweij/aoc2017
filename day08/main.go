package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	f, _ := os.Open("input.txt")
	defer f.Close()

	is, _ := parse(f)

	regs := newRegisters()

	for _, i := range is {
		regs.apply(i)
	}

	return regs.largestValue()
}

func puzzle2() int {
	f, _ := os.Open("input.txt")
	defer f.Close()

	is, _ := parse(f)

	regs := newRegisters()

	for _, i := range is {
		regs.apply(i)
	}

	return regs.largestValueEver()
}
