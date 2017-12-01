package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	s := strings.TrimSpace(string(input))

	fmt.Printf("Puzzle 1: %d\n", captcha(s))

	fmt.Printf("Puzzle 2: %d\n", captcha2(s))
}
