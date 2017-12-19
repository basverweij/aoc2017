package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %s\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() string {
	s, _ := follow(input())
	return s
}

func puzzle2() int {
	_, n := follow(input())
	return n
}
