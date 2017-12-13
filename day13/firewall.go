package main

type firewall struct {
	depth  int
	layers map[int]*layer
}

func newFirewall(depthsAndRanges map[int]int) *firewall {
	f := &firewall{
		depth:  0,
		layers: make(map[int]*layer, len(depthsAndRanges)),
	}

	for d, r := range depthsAndRanges {
		f.layers[d] = &layer{rng: r}

		if d > f.depth {
			f.depth = d
		}
	}

	return f
}

func (f *firewall) advance() {
	for _, l := range f.layers {
		l.advance()
	}
}

func (f *firewall) scanner(depth int) int {
	l, found := f.layers[depth]
	if !found {
		return -1
	}

	return l.scanner
}

type layer struct {
	rng     int
	scanner int
	forward bool
}

func (l *layer) advance() {
	if l.rng == 1 {
		return
	}

	if l.forward {
		l.scanner++
		if l.scanner == l.rng {
			l.forward = false
			l.scanner = l.rng - 2
		}
	} else {
		l.scanner--
		if l.scanner < 0 {
			l.forward = true
			l.scanner = 1
		}
	}
}
