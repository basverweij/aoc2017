package main

import "fmt"

func redistribute(banks []int) (n, l int) {
	seen := make(map[string]int)

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
		if nSeen, found := seen[s]; found {
			l = n - nSeen
			return
		}

		seen[s] = n
	}
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
