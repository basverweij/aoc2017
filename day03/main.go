package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1(361527))
}

func puzzle1(n int) int {
	return distance(spiralCoords(n))
}
