package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1(361527))

	fmt.Printf("Puzzle 2: %d\n", puzzle2(361527))
}

func puzzle1(n int) int {
	return distance(spiralCoords(n))
}

func puzzle2(n int) int {
	return fill(n)
}
