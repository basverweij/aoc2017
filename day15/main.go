package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	genA := newGen(factorGenA, inputGenA, 1)
	genB := newGen(factorGenB, inputGenB, 1)

	return judge(genA, genB, 40000000)
}

func puzzle2() int {
	genA := newGen(factorGenA, inputGenA, multiplesGenA)
	genB := newGen(factorGenB, inputGenB, multiplesGenB)

	return judge(genA, genB, 5000000)
}
