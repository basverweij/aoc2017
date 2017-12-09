package main

import (
	"fmt"
	"strings"
)

type group struct {
	parent   *group
	children []*group
}

func newGroup(parent *group) *group {
	return &group{
		parent:   parent,
		children: make([]*group, 0),
	}
}

func (g *group) String() string {
	cs := make([]string, len(g.children))
	for i, c := range g.children {
		cs[i] = c.String()
	}

	return fmt.Sprintf("{%s}", strings.Join(cs, ","))
}
