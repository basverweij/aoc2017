package main

func captcha2(s string) int {
	digits, n := toDigits([]byte(s))

	sum := 0

	for i := 0; i < n; i++ {
		if digits[i] == digits[(i+n/2)%n] {
			sum += int(digits[i])
		}
	}

	return sum
}
