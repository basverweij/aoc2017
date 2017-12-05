package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1("input.txt"))
}

func puzzle1(path string) int {
	f, _ := os.Open(path)
	js, _ := parse(f)

	return js.StepsToExit()
}
