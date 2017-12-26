package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle2() int {
	return input.strongestBridge()
}
