package main

func optimized() int {
	c := 126900
	primes := genPrimes(c)

	h := 0

	for b := 109900; b <= c; b += 17 {
		if _, isPrime := primes[b]; !isPrime {
			h++
		}
	}

	return h
}

func genPrimes(max int) map[int]struct{} {
	var primes []int

	for i := 2; i <= max; i++ {
		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	primeMap := make(map[int]struct{}, len(primes))
	for _, p := range primes {
		primeMap[p] = struct{}{}
	}

	return primeMap
}
