package main

import (
	"math"
)

type trajectory struct {
	// the ticks at which this trajectory starts c.q. ends
	tStart, tEnd int
	// contain the signum values (-1/1) for the three axes
	xSgn, ySgn, zSgn int
}

var partQroots = []func(*particle) (float64, float64, error){
	func(p *particle) (float64, float64, error) { return qroots(p.a.x, p.v.x, p.p.x) },
	func(p *particle) (float64, float64, error) { return qroots(p.a.y, p.v.y, p.p.y) },
	func(p *particle) (float64, float64, error) { return qroots(p.a.z, p.v.z, p.p.z) },
}

var partPos = []func(*particle, int) float64{
	func(p *particle, t int) float64 { return float64(p.a.x*t*t + p.v.x*t + p.p.x) },
	func(p *particle, t int) float64 { return float64(p.a.y*t*t + p.v.y*t + p.p.y) },
	func(p *particle, t int) float64 { return float64(p.a.z*t*t + p.v.z*t + p.p.z) },
}

// finalTrajectory returns the final trajectory after which
// no more signum changes occur on the axes
func finalTrajectory(p *particle) trajectory {
	tFinal := math.NaN()
	for i := 0; i < 3; i++ {
		_, t1, err := partQroots[i](p)
		if err != nil {
			// no roots
			continue
		}

		if math.IsNaN(tFinal) || t1 > tFinal {
			tFinal = t1
		}
	}

	if math.IsNaN(tFinal) {
		tFinal = 0
	}

	tStart := int(math.Floor(tFinal + 1.0))

	j := trajectory{
		tStart: tStart,
		tEnd:   math.MaxInt32,
		xSgn:   signum(math.Signbit(partPos[0](p, tStart))),
		ySgn:   signum(math.Signbit(partPos[1](p, tStart))),
		zSgn:   signum(math.Signbit(partPos[2](p, tStart))),
	}

	return j
}

func signum(signbit bool) int {
	if signbit {
		return -1
	}

	return 1
}
