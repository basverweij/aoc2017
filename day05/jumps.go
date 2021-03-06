package main

type jumps []int

func (js jumps) StepsToExit() int {
	pc := 0

	for n := 0; ; n++ {
		if pc < 0 || pc >= len(js) {
			return n
		}

		js[pc]++
		pc += js[pc] - 1
	}
}

func (js jumps) StepsToExitWithDecrease() int {
	pc := 0

	var j int
	for n := 0; ; n++ {
		if pc < 0 || pc >= len(js) {
			return n
		}

		j = js[pc]
		if j < 3 {
			js[pc]++
		} else {
			js[pc]--
		}

		pc += j
	}
}
