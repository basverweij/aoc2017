package main

func (cs components) strongestBridge() int {
	lengths := make(map[int][]int, len(cs))
	for i, c := range cs {
		lengths[c.portA] = append(lengths[c.portA], i)
		if c.portA != c.portB {
			lengths[c.portB] = append(lengths[c.portB], i)
		}
	}

	return cs.strongestBridgeFrom(0, lengths, make([]direction, len(cs)), 0)
}

func (cs components) strongestBridgeFrom(length int, lengths map[int][]int, usedBefore []direction, strengthBefore int) int {
	used := make([]direction, len(cs))
	copy(used, usedBefore)

	s := strengthBefore
	for _, i := range lengths[length] {
		if used[i] != none {
			continue
		}

		if cs[i].portA == length {
			used[i] = fromA
			s = max(s, cs.strongestBridgeFrom(cs[i].portB, lengths, used, strengthBefore+cs[i].strength()))
		} else if cs[i].portB == length {
			used[i] = fromB
			s = max(s, cs.strongestBridgeFrom(cs[i].portA, lengths, used, strengthBefore+cs[i].strength()))
		}

		used[i] = none
	}

	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
