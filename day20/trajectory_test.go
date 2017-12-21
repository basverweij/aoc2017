package main

import (
	"fmt"
	"testing"
)

func TestFinalTrajectory(t *testing.T) {
	j0 := finalTrajectory(testPart0)

	j1 := finalTrajectory(testPart1)

	tStart := j0.tStart
	if j1.tStart > tStart {
		tStart = j1.tStart
	}

	fmt.Printf("j0: %v (@t=%d: %d)\n",
		j0,
		tStart, testPart0.originVelocity(j0, tStart))

	fmt.Printf("j1: %v (@t=%d: %d)\n",
		j1,
		tStart, testPart1.originVelocity(j1, tStart))
}

func TestFinalTrajectory2(t *testing.T) {
	j2 := finalTrajectory(testPart2)
	fmt.Println(j2)
}
