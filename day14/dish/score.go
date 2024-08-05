package dish

import "day14/dish/pice"

func arrScore(pices []pice.Pice) int {
	scoreSum := 0
	for i, el := range pices {
		if el == pice.Movable {
			scoreSum += len(pices) - i
		}
	}
	return scoreSum
}
