package main

import "sort"

func checksum2(data [][]int) int {
	s := 0

	for _, row := range data {
		num, denom := divs(row)
		s += denom / num
	}

	return s
}

func divs(row []int) (num int, denom int) {
	sort.Ints(row)

	for d := len(row) - 1; d > 0; d-- {
		denom = row[d]

		for n := 0; n < d; n++ {
			num = row[n]

			if denom%num == 0 {
				return
			}
		}
	}

	return
}
