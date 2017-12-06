package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	steps, _ := redistribute(input)
	return steps
}

func puzzle2() int {
	_, loop := redistribute(input)
	return loop
}
