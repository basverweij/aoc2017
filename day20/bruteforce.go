package main

import (
	"math"
)

func bruteForce(ps []*particle, destroyOnCollision bool) (int, int) {
	for t := 0; t < 1000; t++ {
		for _, p := range ps {
			if p.destroyed {
				continue
			}

			p.v.x += p.a.x
			p.v.y += p.a.y
			p.v.z += p.a.z

			p.p.x += p.v.x
			p.p.y += p.v.y
			p.p.z += p.v.z
		}

		if destroyOnCollision {
			// fmt.Printf("t=%d, particles left: %d\n", t, particlesLeft(ps))

			partPos := make(map[d3]int)

			for i, p := range ps {
				if p.destroyed {
					continue
				}

				if first, found := partPos[p.p]; found {
					// fmt.Printf("destroying %v and %v\n", ps[first], p)
					ps[first].destroyed = true
					p.destroyed = true
				} else {
					partPos[p.p] = i
				}
			}
		}
	}

	closestIndex := -1
	closestDistance := math.MaxInt64
	particlesLeft := 0
	for i, p := range ps {
		if p.destroyed {
			continue
		}

		particlesLeft++

		d := abs(p.p.x) + abs(p.p.y) + abs(p.p.z)

		if d < closestDistance {
			closestDistance = d
			closestIndex = i
		}
	}

	return closestIndex, particlesLeft
}

func particlesLeft(ps []*particle) int {
	left := 0
	for _, p := range ps {
		if !p.destroyed {
			left++
		}
	}

	return left
}
