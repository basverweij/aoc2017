package main

import (
	"fmt"
	"math"
)

func closest(ps []*particle) int {
	js := make([]trajectory, len(ps))

	tStartMax := math.MinInt32
	for i, p := range ps {
		js[i] = finalTrajectory(p)

		if js[i].tStart > tStartMax {
			tStartMax = js[i].tStart
		}
	}

	fmt.Println("tStartMax", tStartMax)

	originVelocityMin := math.MaxInt32
	closestIndex := -1
	for i, j := range js {
		v := abs(ps[i].originVelocity(j, tStartMax))
		fmt.Printf("#%03d: %6v @ %6v -> %6d\n", i, ps[i], j, v)

		if v < originVelocityMin {
			originVelocityMin = v
			closestIndex = i
		}
	}

	return closestIndex
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
