package main

type direction int

const (
	none direction = iota
	fromA
	fromB
)

type component struct {
	portA, portB int
}

type components []component

func (c component) start(dir direction) int {
	if dir == fromA {
		return c.portA
	} else if dir == fromB {

		return c.portB
	}

	panic("invalid direction")
}

func (c component) end(dir direction) int {
	if dir == fromA {
		return c.portB
	} else if dir == fromB {
		return c.portA
	}

	panic("invalid direction")
}

func (c component) strength() int {
	return c.portA + c.portB
}
