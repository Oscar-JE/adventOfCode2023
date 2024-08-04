package reflection

import "errors"

type IndexCount struct {
	counts []int
}

func InitIndexCount(nrOfIndices int) IndexCount {
	counts := []int{}
	i := 0
	for i < nrOfIndices {
		counts = append(counts, 0)
		i++
	}
	return IndexCount{counts: counts}
}

func (i *IndexCount) Align(values []int) {
	for j := range len(i.counts) {
		if IsReflektionIndex(j, values) {
			i.counts[j]++
		}
	}
}

func (c IndexCount) IndexWithNrMatches(nrMatches int) (int, error) {
	for i, el := range c.counts {
		if el == nrMatches {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func IsReflectionWithOneException(index int, values []int) bool { // kan bli av med kodupprepning om man bryter ut innehållet med exception som input
	exception := 1
	loweLimit := 0
	upperLimit := len(values)
	loweRefIndex := index
	upperRefINdex := index + 1
	if loweRefIndex < loweLimit || upperLimit <= upperRefINdex {
		return false
	}
	for loweLimit <= loweRefIndex && upperRefINdex < upperLimit {
		if values[loweRefIndex] != values[upperRefINdex] {
			if exception <= 0 {
				return false
			} else {
				exception--
			}
		}
		loweRefIndex--
		upperRefINdex++
	}
	return true
}
