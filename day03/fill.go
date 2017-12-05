package main

// fill returns the first value that is larger than n
func fill(n int) int {
	// determine dimensions
	c := spiralCoords(n)
	offset := max(c.x, c.y) + 2
	dim := 2*offset + 1

	g := newGrid(dim, offset)

	// set center value
	g.set(0, 0, 1)

	for ring := 1; ; ring++ {
		size := spiralSize(ring) / 4

		for side := 0; side < 4; side++ {
			for idx := 0; idx < size; idx++ {
				c = newCoords(ring, side, idx)
				s := g.sumNeighbours(c.x, c.y)

				if s > n {
					return s
				}

				g.set(c.x, c.y, s)
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
