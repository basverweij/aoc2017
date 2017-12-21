package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())
}

func puzzle1() int {
	fractal, rules := input()
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: %v\n", i, fractal)
		fractal = fractalize(fractal, rules)
	}

	fmt.Printf("5: %v\n", fractal)

	// 110 is too low, 378 is too high
	return fractal.countOn()
}
