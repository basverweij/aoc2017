package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	c, _ := cubeCoordsFromMoves(input...)
	return c.distance()
}

func puzzle2() int {
	_, maxDistance := cubeCoordsFromMoves(input...)
	return maxDistance
}
