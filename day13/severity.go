package main

func severity(depthsAndRanges map[int]int) int {
	f := newFirewall(depthsAndRanges)

	severity := 0

	for d := 0; d <= f.depth; d++ {
		if f.scanner(d) == 0 {
			severity += d * depthsAndRanges[d]
		}

		f.advance()
	}

	return severity
}
