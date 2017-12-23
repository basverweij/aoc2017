package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle2() int {
	v := newVirus(parse(rawInput))

	for i := 0; i < 10000000; i++ {
		v.burst()
	}

	return v.infections
}
