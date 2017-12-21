package main

import "math"

func bruteForce(ps []*particle) int {
	closestIndex := -1

	for t := 1; t < 10000; t++ {
		closestDistance := math.MaxInt64

		for i, p := range ps {
			p.v.x += p.a.x
			p.v.y += p.a.y
			p.v.z += p.a.z

			p.p.x += p.v.x
			p.p.y += p.v.y
			p.p.z += p.v.z

			d := abs(p.p.x) + abs(p.p.y) + abs(p.p.z)

			if d < closestDistance {
				closestDistance = d
				closestIndex = i
			}
		}
	}

	return closestIndex
}
