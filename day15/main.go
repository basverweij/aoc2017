package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())
}

func puzzle1() int {
	genA := newGen(factorGenA, inputGenA)
	genB := newGen(factorGenB, inputGenB)

	return judge(genA, genB, 40000000)
}
