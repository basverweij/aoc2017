package main

func judge(genA, genB *gen, iterations int) int {
	n := 0

	for i := 0; i < iterations; i++ {
		a, b := genA.next(), genB.next()

		if a&0xffff == b&0xffff {
			n++
		}
	}

	return n
}
