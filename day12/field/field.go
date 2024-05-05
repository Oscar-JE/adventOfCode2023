package field

import "day12/springs"

type Field struct {
	layout []springs.Condition
}

func Init(rep string) Field {
	runes := []rune(rep)
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

func (f Field) PossiblePlacements(block int) []int {
	possiblePlacements := []int{}
	for i := range f.layout {
		leftSpacing := i - 1
		rightSpacing := i + block
		if f.isSeparator(leftSpacing) && f.isSeparator(rightSpacing) {
			if f.possibleBrokenBetween(leftSpacing, rightSpacing) {
				possiblePlacements = append(possiblePlacements, i)
			}
		}
		if f.isDamaged(i) {
			break
		}
	}
	return possiblePlacements
}

func (f Field) possibleBrokenBetween(left int, right int) bool {
	allPossiblyOccupied := true
	for middelPoint := left + 1; middelPoint < right; middelPoint++ {
		allPossiblyOccupied = allPossiblyOccupied && f.possiblyOccupied(middelPoint)
	}
	return allPossiblyOccupied
}

func (f Field) isUnknown(index int) bool {
	if index < 0 || index >= len(f.layout) {
		return false
	}
	return f.layout[index] == springs.Unknown()
}

func (f Field) isDamaged(index int) bool {
	if index < 0 || index >= len(f.layout) {
		return false
	}
	return f.layout[index] == springs.Damaged()
}

func (f Field) possiblyOccupied(index int) bool {
	if index < 0 || index >= len(f.layout) {
		return false
	}
	return f.layout[index] == springs.Unknown() || f.layout[index] == springs.Damaged()
}

func (f Field) isSeparator(index int) bool {
	if index < 0 || index >= len(f.layout) {
		return true
	}
	return f.layout[index] == springs.Operational() || f.layout[index] == springs.Unknown()
}

func (f Field) SubField(index int) Field {
	if index < len(f.layout) {
		return Field{layout: f.layout[index:]}
	}
	return Field{layout: []springs.Condition{}}
}
