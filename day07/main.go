package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Puzzle 1: %s\n", puzzle1())

	fmt.Printf("Puzzle 2: %s\n", puzzle2())
}

func puzzle1() string {
	f, _ := os.Open("input.txt")
	programs, _ := parse(f)
	root, _ := root(programs)

	return root.Name
}

func puzzle2() string {
	f, _ := os.Open("input.txt")
	programs, _ := parse(f)
	unb, w, _ := unbalanced(programs)

	return fmt.Sprintf("%s: %d", unb.Name, w)
}
