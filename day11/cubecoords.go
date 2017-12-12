// https://www.redblobgames.com/grids/hexagons/#coordinates-cube
package main

type cubecoords struct {
	x, y, z int
}

func cubeCoordsFromMoves(dirs ...direction) (*cubecoords, int) {
	c := &cubecoords{}
	maxDistance := 0

	for _, dir := range dirs {
		c.move(dir)

		maxDistance = max(maxDistance, c.distance())
	}

	return c, maxDistance
}

var moves = map[direction][3]int{
	n:  [3]int{0, 1, -1},
	ne: [3]int{1, 0, -1},
	se: [3]int{1, -1, 0},
	s:  [3]int{0, -1, 1},
	sw: [3]int{-1, 0, 1},
	nw: [3]int{-1, 1, 0},
}

func (c *cubecoords) move(dir direction) {
	m := moves[dir]
	c.x += m[0]
	c.y += m[1]
	c.z += m[2]
}

func (c *cubecoords) distance() int {
	return max(max(abs(c.x), abs(c.y)), abs(c.z))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
