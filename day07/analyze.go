package main

import (
	"errors"
	"fmt"
)

var (
	errMultipleRoots = errors.New("multiple roots found")
	errNoRoot        = errors.New("no root found")
)

func root(programs []*program) (*program, error) {
	subs := make(map[string]*program, len(programs))

	for _, p := range programs {
		subs[p.Name] = p
	}

	for _, p := range programs {
		for _, s := range p.SubPrograms {
			delete(subs, s)
		}
	}

	if len(subs) > 1 {
		return nil, errMultipleRoots
	}

	for _, p := range subs {
		return p, nil
	}

	return nil, errNoRoot
}

func unbalanced(programs []*program) (*program, int, error) {
	root, err := root(programs)
	if err != nil {
		return nil, 0, err
	}

	r := newRuntime(programs)

	rootw := 0
	for {
		p, w, err := unbalancedSubProgram(r, root)
		if err != nil {
			return nil, 0, err
		}

		if p == nil {
			// root is unbalanced
			return root, rootw, nil
		}

		root = p
		rootw = w
	}
}

func unbalancedSubProgram(r *runtime, p *program) (*program, int, error) {
	if len(p.SubPrograms) == 0 {
		return nil, 0, nil
	}

	weights := make(map[int][]string)
	for _, n := range p.SubPrograms {
		w := r.totalWeight(n)
		weights[w] = append(weights[w], n)
	}

	if len(weights) == 1 {
		// no unbalanced sub programs
		return nil, 0, nil
	}

	if len(weights) > 2 {
		return nil, 0, fmt.Errorf("too many different weights: %v", weights)
	}

	var unb *program
	unbWeight := 0
	balWeight := 0

	for w, pbw := range weights {
		if len(pbw) == 1 {
			unb = r.program(pbw[0])
			unbWeight = w
			continue
		}

		balWeight = w
	}

	return unb, unb.Weight - unbWeight + balWeight, nil
}
