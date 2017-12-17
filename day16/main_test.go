package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	main()
}

func TestFindRecurrence(t *testing.T) {
	moves := input()
	l := newLine(16)

	seen := make(map[string]int)

	for i := 0; i < 100; i++ {
		dance(l, moves)

		s := l.String()
		if idx, found := seen[s]; found {
			fmt.Printf("Found %s previously at %d, now at %d (recurrence = %d)\n", s, idx, i, i-idx)
			break
		}

		seen[s] = i
	}
}
