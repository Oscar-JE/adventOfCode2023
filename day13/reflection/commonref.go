package reflection

type CommonRef struct { // jobbar med antagandet att det är en unik reflektion
	refIndex []int
}

func FindPossibleRef(values []int) CommonRef { // börjar med den dummaste långsamma implementationen
	refindxes := []int{}
	for i := range values {
		if isReflektionIndex(i, values) {
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
		if isReflektionIndex(index, values) {
			newIndexes = append(newIndexes, index)
		}
	}
	c.refIndex = newIndexes
}

func isReflektionIndex(index int, values []int) bool { // index till höger om reflextionen är det som är viktigt
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
