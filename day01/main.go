package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	s := strings.TrimSpace(string(input))

	fmt.Println(captcha(s))
}
