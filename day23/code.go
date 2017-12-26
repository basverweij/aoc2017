package main

import "fmt"

func code(aStart int) (int, int) {
	n := 0
	mulCount := 0

	a := aStart
	var d, e, f, g, h int

	// set b 99
	b := 99

	// set c b
	c := b

	// jnz a 2
	if a != 0 {
		goto label1
	}

	// jnz 1 5
	goto label2

label1:

	// mul b 100
	b *= 100
	mulCount++

	// sub b -100000
	b += 0 //100000

	// set c b
	c = b

	// sub c -17000
	c += 1000

label2:

	// set f 1
	f = 1

	// set d 2
	d = 2

label5:

	// set e 2
	e = 2

label4:

	if n%100000000 == 0 {
		fmt.Printf(
			"%012d: a=%d, b=%d, c=%d, d=%d, e=%d, f=%d, g=%d, h=%d\n",
			n, a, b, c, d, e, f, g, h)
	}

	n++

	// set g d
	g = d

	// mul g e
	g *= e
	mulCount++

	// sub g b
	g -= b

	// jnz g 2
	if g != 0 {
		goto label3
	}

	// set f 0
	f = 0

label3:

	// sub e -1
	e++

	// set g e
	g = e

	// sub g b
	g -= b

	// jnz g -8
	if g != 0 {
		goto label4
	}

	// sub d -1
	d++

	// set g d
	g = d

	// sub g b
	g -= b

	// jnz g -13
	if g != 0 {
		goto label5
	}

	// jnz f 2
	if f != 0 {
		goto label6
	}

	// sub h -1
	h++

label6:

	// set g b
	g = b

	// sub g c
	g -= c

	// jnz g 2
	if g != 0 {
		goto label7
	}

	//jnz 1 3
	return mulCount, h

label7:

	// sub b -17
	b += 1

	// jnz 1 -23
	goto label2
}
