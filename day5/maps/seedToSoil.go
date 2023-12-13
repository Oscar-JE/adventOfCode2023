package maps

import "fmt"

type Transform struct {
	subTransforms []SubTransform
}

func InitTransform(subT []SubTransform) Transform {
	return Transform{subT}
}

func (t Transform) Do(in int) int {
	out := in
	mapped := false
	for _, subTrans := range t.subTransforms {
		out, mapped = subTrans.Do(out)
		if mapped {
			break
		}
	}
	return out
}

type SubTransform struct {
	fromStart int
	toStart   int
	length    int
}

func InitSubTransform(fromStart int, toStart int, length int) SubTransform {
	return SubTransform{fromStart: fromStart, toStart: toStart, length: length}
}

func (m SubTransform) Do(in int) (int, bool) {
	if m.fromStart <= in && in < (m.fromStart+m.length) {
		return in + (m.toStart - m.fromStart), true
	}
	return in, false
}

func Atlas(myMappings []Transform, in int) int {
	toMap := in
	for _, myMap := range myMappings {
		toMap = myMap.Do(toMap)
		fmt.Printf("%d ", toMap)
	}
	return toMap
}
