package main

import "fmt"

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())

	fmt.Printf("Puzzle 2: %d\n", puzzle2())
}

func puzzle1() int {
	fractal, rules := input()
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: %v\n", i, fractal)
		fractal = fractalize(fractal, rules)
	}

	fmt.Printf("5: %v\n", fractal)

	return fractal.countOn()
}

func puzzle2() int {
	fractal, rules := input()
	for i := 0; i < 18; i++ {
		fmt.Printf("%02d: size=%d\n", i, fractal.size)
		fractal = fractalize(fractal, rules)
	}

	fmt.Printf("18: size=%d\n", fractal.size)

	// 72363 is too low
	return fractal.countOn()
}
