package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Puzzle 1: %s\n", puzzle1())
}

func puzzle1() string {
	f, _ := os.Open("input.txt")
	programs, _ := parse(f)
	root, _ := root(programs)

	return root.Name
}
