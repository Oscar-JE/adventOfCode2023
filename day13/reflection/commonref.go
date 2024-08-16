package reflection

type CommonRef struct {
	refIndex []int
}

func FindPossibleRef(values []int) CommonRef {
	refindxes := []int{}
	for i := range values {
		if IsReflektionIndex(i, values) {
			refindxes = append(refindxes, i)
		}
	}
	return CommonRef{refIndex: refindxes}
}

func (c CommonRef) HasReflection() bool {
	return len(c.refIndex) > 0
}

func (c *CommonRef) Align(values []int) {
	newIndexes := []int{}
	for _, index := range c.refIndex {
		if IsReflektionIndex(index, values) {
			newIndexes = append(newIndexes, index)
		}
	}
	c.refIndex = newIndexes
}

func IsReflektionIndex(index int, values []int) bool {
	loweLimit := 0
	upperLimit := len(values)
	loweRefIndex := index
	upperRefINdex := index + 1
	if loweRefIndex < loweLimit || upperLimit <= upperRefINdex {
		return false
	}
	for loweLimit <= loweRefIndex && upperRefINdex < upperLimit {
		if values[loweRefIndex] != values[upperRefINdex] {
			return false
		}
		loweRefIndex--
		upperRefINdex++
	}
	return true
}

func (c CommonRef) GetIndex() int {
	if len(c.refIndex) != 1 {
		panic("no unique reflection found")
	}
	return c.refIndex[0]
}
