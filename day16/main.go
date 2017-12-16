package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %s\n", puzzle1())

	fmt.Printf("Puzzle 2: %s\n", puzzle2())
}

func puzzle1() string {
	l := newLine(16)
	dance(l, input())

	return l.String()
}

func puzzle2() string {
	moves := input()

	l := newLine(16)

	// 1000000000
	for i := 0; i < 1000; i++ {
		dance(l, moves)
	}

	return l.String()
}
