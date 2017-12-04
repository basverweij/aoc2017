package main

func distance(c coords) int {
	return abs(c.x) + abs(c.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
