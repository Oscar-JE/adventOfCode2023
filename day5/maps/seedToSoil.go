package maps

type Mapping struct {
	fromStart int
	toStart   int
	length    int
}

func InitMapping(fromStart int, toStart int, length int) Mapping {
	return Mapping{fromStart: fromStart, toStart: toStart, length: length}
}

func (m Mapping) Do(in int) int {
	if m.fromStart <= in && in < (m.fromStart+m.length) {
		return in + (m.toStart - m.fromStart)
	}
	return in
}

func Atlas(myMappings []Mapping, in int) int {
	toMap := in
	for _, myMap := range myMappings {
		toMap = myMap.Do(toMap)
	}
	return toMap
}
