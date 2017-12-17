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

func puzzle2() string {
	moves := input()
	moves = moves[:3] // po/k, x4/0, s12
	moves = []move{
		newPartner('o', 'k'),
		newExchange(4, 0),
		newSpin(12),
	}

	l := newLine(16)
	dance(l, moves)

	// create effective translation table for this dance
	t := make([]int, 16)
	for i := range t {
		t[int(l.get(i))] = i
	}

	fmt.Printf("%#v\n", t)

	l = newLine(16)
	a := []byte(l.String())
	b := make([]byte, len(a))

	fmt.Printf("l: %s\n", l.String())
	fmt.Printf("a: %s\n", string(a))

	for i := 0; i < 3; i++ {
		dance(l, moves)

		for j, k := range t {
			b[k] = a[j]
		}
		a, b = b, a

		fmt.Printf("-- %d --\n", i)
		fmt.Printf("l: %s\n", l.String())
		fmt.Printf("a: %s\n", string(a))
		// fmt.Printf("b: %s\n", string(b))
	}

	// not: igbdhmcpknafjelo

	return string(a)
}
