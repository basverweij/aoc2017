package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	root, _ := scan(input)

	return score(root)
}

func puzzle2() int {
	_, garbageCount := scan(input)

	return garbageCount
}
