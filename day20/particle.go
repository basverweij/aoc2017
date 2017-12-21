package main

type d3 struct {
	x, y, z int
}

type particle struct {
	p, v, a d3
}

func (p particle) originVelocity(j trajectory, t int) int {
	if t < j.tStart || t > j.tEnd {
		panic("invalid ticks for trajectory")
	}

	return p.v.x + j.xSgn*2*p.a.x*t +
		p.v.y + j.ySgn*2*p.a.y*t +
		p.v.z + j.zSgn*2*p.a.z*t
}
