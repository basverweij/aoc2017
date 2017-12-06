package main

import "fmt"

func redistribute(banks []int) int {
	var n int

	seen := make(map[string]struct{})

	for n = 1; ; n++ {
		max, bank := max(banks)
		banks[bank] = 0

		for ; max > 0; max-- {
			bank++
			if bank >= len(banks) {
				bank = 0
			}

			banks[bank]++
		}

		s := fmt.Sprintf("%#v", banks)
		if _, found := seen[s]; found {
			break
		}

		seen[s] = struct{}{}
	}

	return n
}

func max(banks []int) (max, bank int) {
	max = banks[bank]

	for i := 1; i < len(banks); i++ {
		if banks[i] > max {
			bank = i
			max = banks[i]
		}
	}

	return
}
