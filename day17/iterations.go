package main

type buffer interface {
	move(int)
	insert(int)
}

func iterations(b buffer, iterationCount, moveAmount int) {
	for i := 1; i <= iterationCount; i++ {
		b.move(moveAmount)
		b.insert(i)
	}
}
