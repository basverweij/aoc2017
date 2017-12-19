package main

type instruction func(executable)

type value func(executable) int

func newSnd(value value) instruction {
	return func(x executable) {
		x.send(value(x))
	}
}

func newRcv(reg string) instruction {
	return func(x executable) {
		x.set(reg, x.receive())
	}
}

func newSet(reg string, value value) instruction {
	return func(x executable) {
		x.set(reg, value(x))
	}
}

func newAdd(reg string, value value) instruction {
	return func(x executable) {
		x.set(reg, x.get(reg)+value(x))
	}
}

func newMul(reg string, value value) instruction {
	return func(x executable) {
		x.set(reg, x.get(reg)*value(x))
	}
}

func newMod(reg string, value value) instruction {
	return func(x executable) {
		x.set(reg, x.get(reg)%value(x))
	}
}

func newJgz(condition value, value value) instruction {
	return func(x executable) {
		if condition(x) > 0 {
			x.jump(value(x) - 1) // subtract 1 because program automatically adds 1
		}
	}
}
