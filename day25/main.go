package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Printf("Puzzle 1: %d\n", puzzle1())
}

func puzzle1() int {
	f, _ := ioutil.ReadFile("input.txt")
	m := input(string(f))
	d := runDiagnostics(m)

	return d
}
