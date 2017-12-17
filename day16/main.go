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

const total = 1000000000
const recurrence = 36

func puzzle2() string {
	moves := input()
	l := newLine(16)

	remaining := total % recurrence

	for i := 0; i < remaining; i++ {
		dance(l, moves)
	}

	// not: igbdhmcpknafjelo

	return l.String()
}
