package point

import "sort"

type VecSet struct {
	vectors []Vec
}

func (v VecSet) GetVectors() []Vec {
	return v.vectors
}

func (set VecSet) In(v Vec) bool {
	for _, el := range set.vectors {
		if el == v {
			return true
		}
	}
	return false
}

func (set *VecSet) Add(v Vec) {
	set.vectors = append(set.vectors, v)
}

func (set VecSet) IsEmpty() bool {
	return len(set.vectors) == 0
}

func EmptySet() VecSet {
	return VecSet{[]Vec{}}
}

func InitVecSet(vectors []Vec) VecSet {
	return VecSet{vectors: vectors}
}

func (set VecSet) Translate(v Vec) VecSet {
	s := EmptySet()
	for _, vec := range set.vectors {
		s.Add(vec.Add(v))
	}
	return s
}

func (set *VecSet) Rotate90counterclockwise() {
	for i := range set.vectors {
		set.vectors[i].Rotate90counterclockwise()
	}
}

func (set VecSet) Eq(other VecSet) bool {
	if len(set.vectors) != len(other.vectors) {
		return false
	}
	eq := true
	sort.Sort(set)
	sort.Sort(other)
	for i := range set.vectors {
		eq = eq && set.vectors[i] == other.vectors[i]
	}
	return eq
}

func (set VecSet) SharesElements(other VecSet) bool {
	for _, v := range set.vectors {
		if other.Has(v) {
			return true
		}
	}
	return false
}

func (set VecSet) Has(v Vec) bool {
	for _, el := range set.vectors {
		if el == v {
			return true
		}
	}
	return false
}

// sort demanded functions
func (set VecSet) Len() int           { return len(set.vectors) }
func (set VecSet) Swap(i, j int)      { set.vectors[i], set.vectors[j] = set.vectors[j], set.vectors[i] }
func (set VecSet) Less(i, j int) bool { return set.vectors[i].Less(set.vectors[j]) }
