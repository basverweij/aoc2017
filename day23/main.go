package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	//mulCount, _ := newProgram(input()).run(nil)

	mulCount, _ := code(0)

	return mulCount
}

func puzzle2() int {
	start := map[rune]int{'a': 1}
	_, h := newProgram(input()).run(start)

	// 887 is too low
	// 999 is too high
	// 896 is not correct
	// _, h := code(1)

	return h
}
