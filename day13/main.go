package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	s, _ := severity(input, 0, false)
	return s
}

func puzzle2() int {
	return clearPass(input)
}
