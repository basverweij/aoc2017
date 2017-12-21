package main

func fractalize(b *bitgrid, rules map[sizedKey]*bitgrid) *bitgrid {
	d := 2
	if b.size%2 == 1 {
		d = 3
	}

	n := b.size / d
	fractal := newBitgrid(n * (d + 1))

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			k := b.key(i*d, j*d, d)
			r := rules[sizedKey{d, k}]

			fractal.copyFrom(r, i*(d+1), j*(d+1))
		}
	}

	return fractal
}
