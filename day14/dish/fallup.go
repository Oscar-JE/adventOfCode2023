package dish

import "day14/dish/pice"

// varför göra det komplicerat

func fallUpp(pices []pice.Pice) []pice.Pice {
	for i, el := range pices {
		if el == pice.Movable {
			newIndex := fallUppOne(i, pices)
			pices[i] = pice.Empty
			pices[newIndex] = pice.Movable
		}
	}
	return pices
}

// långsamt men enkelt
func fallUppOne(start int, pices []pice.Pice) int {
	if start == 0 {
		return 0
	} else if pices[start-1] != pice.Empty {
		return start
	} else {
		return fallUppOne(start-1, pices)
	}
}
