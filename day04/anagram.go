package main

// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func anagrams(s string) []string {
	as := []string{s}

	a := []byte(s)
	n := len(s)
	c := make([]int, n)

	for i := 0; i < n; {
		if c[i] < i {
			if i%2 == 0 {
				a[0], a[i] = a[i], a[0]
			} else {
				a[c[i]], a[i] = a[i], a[c[i]]
			}

			as = append(as, string(a))

			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}

	return as
}
