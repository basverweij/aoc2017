package main

func iterations(b *buf, iterationCount, moveAmount int) int {
	for i := 1; i <= iterationCount; i++ {
		b.move(moveAmount)
		b.insert(i)
	}

	// get value after last insert
	b.move(1)

	return b.d[b.pos]
}
