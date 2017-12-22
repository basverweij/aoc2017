package main

import (
	"strings"
)

func parse(raw string) *bitgrid {
	ss := strings.Split(raw, "/")
	b := newBitgrid(len(ss))

	for i, s := range ss {
		for j, r := range s {
			if r == '#' {
				b.set(j, i, true)
			}
		}
	}

	return b
}

var rotations = map[string]func(*bitgrid, int, int) bool{
	"rot000": func(b *bitgrid, x, y int) bool { return b.get(x, y) },
	"rot090": func(b *bitgrid, x, y int) bool { return b.get(y, b.size-1-x) },
	"rot180": func(b *bitgrid, x, y int) bool { return b.get(b.size-1-y, b.size-1-x) },
	"rot270": func(b *bitgrid, x, y int) bool { return b.get(b.size-1-y, x) },
}

var flips = map[string]func(*bitgrid, int, int) (int, int, bool){
	"none": func(b *bitgrid, x, y int) (int, int, bool) { return x, y, true },
	"vert": func(b *bitgrid, x, y int) (int, int, bool) { return x, b.size - 1 - y, true },
	"hori": func(b *bitgrid, x, y int) (int, int, bool) { return b.size - 1 - x, y, true },
	"both": func(b *bitgrid, x, y int) (int, int, bool) { return b.size - 1 - x, b.size - 1 - y, true },
}

func allTransforms(source *bitgrid) map[string][]*bitgrid {
	all := make(map[string][]*bitgrid)

	for _, r := range rotations {
		for _, f := range flips {
			b := newBitgrid(source.size)

			for i := 0; i < b.size; i++ {
				for j := 0; j < b.size; j++ {
					if r(source, i, j) {
						b.set(f(b, i, j))
					}
				}
			}

			k := b.key(0, 0, b.size)
			all[k] = append(all[k], b)
		}
	}

	return all
}
