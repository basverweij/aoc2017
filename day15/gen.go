package main

const mod = 2147483647

type gen struct {
	factor    int
	value     int
	multiples int
}

func newGen(factor, value, multiples int) *gen {
	return &gen{factor, value, multiples}
}

func (g *gen) next() int {
	for {
		g.value *= g.factor
		g.value %= mod

		if g.value%g.multiples == 0 {
			return g.value
		}
	}
}
