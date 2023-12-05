package ticket

// oklart om dubbelmatch på ett tal är möjligt
func NrOfMatches(winningNumbers []int, holdingNumbers []int) int {
	matched := 0
	for i := range winningNumbers {
		for j := range holdingNumbers {
			if winningNumbers[i] == holdingNumbers[j] {
				matched++
			}
		}
	}
	return matched
}

func Winnings(matches int) int {
	if matches < 1 {
		return 0
	}
	return powerOftwo(matches - 1)
}

func powerOftwo(exponent int) int {
	if exponent <= 0 {
		return 1
	}
	return 2 * powerOftwo(exponent-1)
}
