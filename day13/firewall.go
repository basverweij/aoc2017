package main

type firewall struct {
	depth  int
	layers map[int]*layer
	t      int
}

func newFirewall(depthsAndRanges map[int]int) *firewall {
	f := &firewall{
		depth:  0,
		layers: make(map[int]*layer, len(depthsAndRanges)),
	}

	for d, r := range depthsAndRanges {
		f.layers[d] = &layer{rng: r, mod: r*2 - 2}

		if d > f.depth {
			f.depth = d
		}
	}

	return f
}

func (f *firewall) advance() {
	f.t++
}

func (f *firewall) scanner(depth int) int {
	l, found := f.layers[depth]
	if !found {
		return -1
	}

	s := f.t % l.mod
	if s < l.rng {
		return s
	}

	return l.mod - s
}

type layer struct {
	rng int
	mod int
}
