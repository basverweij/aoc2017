package main

// validateAll returns the number of valid passphrases
func validateAll(ps []passphrase, validateFunc func(passphrase) bool) int {
	n := 0

	for _, p := range ps {
		if validateFunc(p) {
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

func isValidWithAnagrams(p passphrase) bool {
	seen := make(map[string]struct{})

	for _, w := range p {
		if _, found := seen[w]; found {
			return false
		}

		for _, a := range anagrams(w) {
			seen[a] = struct{}{}
		}
	}

	return true
}
