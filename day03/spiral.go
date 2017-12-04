package main

type coords struct {
	x, y int
}

var (
	sides = []coords{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
)

func spiralCoords(n int) coords {
	if n == 1 {
		return coords{}
	}

	// remove 0-ring and make n zero based
	n -= 2

	// get ring
	ring := 1
	var size int
	for ; n >= 0; ring++ {
		size = spiralSize(ring)

		if n < size {
			break
		}

		n -= size
	}

	// get side of ring and index within side
	side := n / (size / 4)
	idx := n % (size / 4)

	return coords{
		sides[side].x*ring + sides[side].y*(idx+1-ring),
		sides[side].y*ring + sides[side].x*(idx+1-ring),
	}
}

func spiralSize(ring int) int {
	if ring == 0 {
		return 1
	}

	return 8 * ring
}
