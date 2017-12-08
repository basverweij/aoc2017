package main

import "errors"

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
