package point

type VecList struct {
	vecs []Vec
}

func InitVecList(vecs []Vec) VecList {
	return VecList{vecs: vecs}
}

func (vs VecList) Get(index int) Vec {
	internalIndex := index % vs.Len()
	return vs.vecs[internalIndex]
}

func (vs VecList) Len() int {
	return len(vs.vecs)
}

func (vs VecList) Eq(other VecList) bool {
	if vs.Len() != other.Len() {
		return false
	}
	equal := true
	for i := range vs.Len() {
		equal = equal && vs.vecs[i] == other.vecs[i]
		if !equal {
			break
		}
	}
	return equal
}

func (vs *VecList) Append(p Vec) {
	vs.vecs = append(vs.vecs, p)
}

func (vs *VecList) Popp() Vec {
	retValue := vs.vecs[0]
	vs.vecs = vs.vecs[1:]
	return retValue
}

func (vs VecList) IsEmpty() bool {
	return vs.Len() == 0
}

func (vs VecList) Has(v Vec) bool {
	for i := range vs.vecs {
		if vs.vecs[i] == v {
			return true
		}
	}
	return false
}

func (vs VecList) FindCLosestTo(v Vec) (Vec, int) {
	index := 0
	dist := L2DistanceSquared(v, vs.vecs[index])
	vec := vs.vecs[index]
	for i := 1; i < vs.Len(); i++ {
		ldist := L2DistanceSquared(v, vs.vecs[i])
		if ldist < dist {
			dist = ldist
			vec = vs.vecs[i]
			index = i
		}
	}
	return vec, index
}
