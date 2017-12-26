package main

func (cs components) strongestBridge() int {
	lengths := make(map[int][]int, len(cs))
	for i, c := range cs {
		lengths[c.portA] = append(lengths[c.portA], i)
		if c.portA != c.portB {
			lengths[c.portB] = append(lengths[c.portB], i)
		}
	}

	s, _ := cs.strongestBridgeFrom(0, lengths, make([]direction, len(cs)), 0, 0)

	return s
}

func (cs components) strongestBridgeFrom(fromLength int, lengths map[int][]int, usedBefore []direction, strengthBefore, lengthBefore int) (strength, length int) {
	used := make([]direction, len(cs))
	copy(used, usedBefore)

	strength = strengthBefore
	length = lengthBefore

	var s, l int
	for _, i := range lengths[fromLength] {
		if used[i] != none {
			continue
		}

		if cs[i].portA == fromLength {
			used[i] = fromA
			s, l = cs.strongestBridgeFrom(cs[i].portB, lengths, used, strengthBefore+cs[i].strength(), lengthBefore+1)
		} else { // cs[i].portB == fromLength
			used[i] = fromB
			s, l = cs.strongestBridgeFrom(cs[i].portA, lengths, used, strengthBefore+cs[i].strength(), lengthBefore+1)
		}

		used[i] = none

		if l > length {
			length = l
			strength = s
		} else if l == length && s > strength {
			strength = s
		}
	}

	return
}
