package middle

func IsEven(n int) bool {
	return (n/2)*2 == n
}

func MiddlePointFromZero(n int) int {
	if IsEven(n) {
		return n / 2
	} else {
		return n/2 + 1
	}
}
