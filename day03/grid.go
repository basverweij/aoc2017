package main

import "fmt"

type grid struct {
	offset int
	data   [][]int
}

func newGrid(dim, offset int) *grid {
	g := &grid{
		offset: offset,
		data:   make([][]int, dim),
	}

	for i := 0; i < dim; i++ {
		g.data[i] = make([]int, dim)
	}

	return g
}

func (g *grid) set(x, y, v int) {
	g.data[y+g.offset][x+g.offset] = v

	return
}

func (g *grid) sumNeighbours(x, y int) int {
	return g.data[y+g.offset-1][x+g.offset-1] +
		g.data[y+g.offset][x+g.offset-1] +
		g.data[y+g.offset+1][x+g.offset-1] +
		g.data[y+g.offset-1][x+g.offset] +
		g.data[y+g.offset][x+g.offset] +
		g.data[y+g.offset+1][x+g.offset] +
		g.data[y+g.offset-1][x+g.offset+1] +
		g.data[y+g.offset][x+g.offset+1] +
		g.data[y+g.offset+1][x+g.offset+1]
}

func (g *grid) String() string {
	s := ""

	for i := 0; i < len(g.data); i++ {
		s += fmt.Sprintf("\n %3d: ", i)
		for j := 0; j < len(g.data[i]); j++ {
			s += fmt.Sprintf("%3d, ", g.data[i][j])
		}
	}

	return s
}
