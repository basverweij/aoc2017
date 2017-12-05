package main

// validateAll returns the number of valid passphrases
func validateAll(ps []passphrase) int {
	n := 0

	for _, p := range ps {
		if isValid(p) {
			n++
		}
	}

	return n
}

func isValid(p passphrase) bool {
	seen := make(map[string]struct{})

	for _, w := range p {
		if _, found := seen[w]; found {
			return false
		}

		seen[w] = struct{}{}
	}

	return true
}
