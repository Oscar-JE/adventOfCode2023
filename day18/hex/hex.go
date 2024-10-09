package hex

func Parse(hex string) int {
	symbols := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	lenSymbols := len(symbols)
	base := 1
	hexRep := []rune(hex)
	hexLen := len(hexRep)
	sum := 0
	for i := range hexRep {
		relevantRune := hexRep[hexLen-1-i]
		sum += indexOf(symbols, relevantRune) * base
		base *= lenSymbols
	}
	return sum
}

func indexOf(symbols []rune, symbol rune) int {
	for i, r := range symbols {
		if r == symbol {
			return i
		}
	}
	panic("Illegal symbol given to hex parsing")
}
