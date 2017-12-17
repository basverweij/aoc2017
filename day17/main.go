package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	b := newBuf(2018)
	iterations(b, 2017, input)

	// get value after last insert
	b.move(1)

	return b.d[b.pos]
}

func puzzle2() int {
	b := newZeroBuf()
	iterations(b, 50000000, input)

	return b.valueAfterZero
}
