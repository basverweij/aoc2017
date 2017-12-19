package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())
}

func puzzle1() int {
	code := input()

	p := newProgram(code)
	p.run()

	isRecovered, freq := p.isRecovered()
	if !isRecovered {
		return -1
	}

	return freq
}
