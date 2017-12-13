package main

func severity(depthsAndRanges map[int]int, delay int, stopAfterDetection bool) (int, bool) {
	f := newFirewall(depthsAndRanges)

	f.t = delay

	severity, detected := 0, false

	for d := 0; d <= f.depth; d++ {
		if f.scanner(d) == 0 {
			severity += d * depthsAndRanges[d]
			detected = true

			if stopAfterDetection {
				return severity, detected
			}
		}

		f.advance()
	}

	return severity, detected
}

func clearPass(depthsAndRanges map[int]int) int {
	for delay := 0; ; delay++ {
		_, detected := severity(depthsAndRanges, delay, true)
		if !detected {
			return delay
		}
	}
}
