package main

import "math"

func checksum(data [][]int) int {
	s := 0

	for _, row := range data {
		min, max := minMax(row)
		s += max - min
	}

	return s
}

func minMax(row []int) (min int, max int) {
	min = math.MaxInt32
	max = math.MinInt32

	for _, value := range row {
		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	return
}
