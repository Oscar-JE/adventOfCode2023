package math

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func GCD(a int, b int) int {
	if a <= 1 || b <= 1 {
		return 1
	}
	lesser := Min(a, b)
	greater := Max(a, b)
	for lesser != 0 {
		lesser_next := modulus(greater, lesser)
		greater = lesser
		lesser = lesser_next
	}
	return greater
}

func modulus(a int, modulus int) int {
	quotient := a / modulus
	reminder := a - quotient*modulus
	return reminder
}
