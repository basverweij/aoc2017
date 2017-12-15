package main

const mod = 2147483647

type gen struct {
	factor int
	value  int
}

func newGen(factor, value int) *gen {
	return &gen{factor, value}
}

func (g *gen) next() int {
	g.value *= g.factor
	g.value %= mod

	return g.value
}
