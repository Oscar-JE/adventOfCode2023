package field

import "day12/springs"

type Field struct {
	layout []springs.Condition
}

func Init(runes []rune) Field {
	layout := []springs.Condition{}
	for _, r := range runes {
		layout = append(layout, springs.Init(r))
	}
	return Field{layout: layout}
}

func InitFromSprings(springs []springs.Condition) Field {
	return Field{layout: springs}
}

func (f Field) Eq(other Field) bool {
	if len(f.layout) != len(other.layout) {
		return false
	}
	eq := true
	for i := range f.layout {
		eq = eq && f.layout[i] == other.layout[i]
	}
	return eq
}
