package main

var nibbleCounts = []int{
	0, 1, 1, 2,
	1, 2, 2, 3,
	1, 2, 2, 3,
	2, 3, 3, 4,
}

func bitCount(buf []byte) int {
	n := 0

	for _, b := range buf {
		n += nibbleCounts[b&0xf] + nibbleCounts[b>>4]
	}

	return n
}
