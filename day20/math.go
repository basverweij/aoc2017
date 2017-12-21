package main

import (
	"errors"
	"math"
)

var (
	errNoRoots = errors.New("no roots")
)

// qroots returns the intercepts for a quadratic function defined as
// ax^2 + bx + c.
func qroots(a, b, c int) (float64, float64, error) {
	if a == 0 {
		// not a quadratic function :)
		if b == 0 {
			return 0, 0, errNoRoots
		}

		t := float64(c / b)

		return t, t, nil
	}

	disc := float64(b*b - 4*a*c)
	if disc < 0 {
		return 0, 0, errNoRoots
	}

	sqrt := math.Sqrt(float64(disc))
	t0 := (float64(-b) - sqrt) / float64(2*a)
	t1 := (float64(-b) + sqrt) / float64(2*a)

	if t0 > t1 {
		return t1, t0, nil
	}

	return t0, t1, nil
}
