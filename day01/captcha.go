package main

func captcha(s string) int {
	digits, n := toDigits([]byte(s))

	sum := 0

	for i := 0; i < n; i++ {
		if digits[i] == digits[(i+1)%n] {
			sum += int(digits[i])
		}
	}

	return sum
}

func toDigits(s []byte) ([]byte, int) {
	n := len(s)
	d := make([]byte, n)

	for i := 0; i < n; i++ {
		d[i] = s[i] - '0'
	}

	return d, n
}
