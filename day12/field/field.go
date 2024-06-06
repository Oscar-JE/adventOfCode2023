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

func (f Field) HasDamaged() bool {
	for _, s := range f.layout {
		if s == springs.Damaged() {
			return true
		}
	}
	return false
}

func (f *Field) Unfold(times int) { // denna komme tas bort
	length := len(f.layout)
	for i := 0; i < times; i++ {
		f.layout = append(f.layout, springs.Unknown())
		for j := 0; j < length; j++ {
			f.layout = append(f.layout, f.layout[j])
		}

	}
}

func (f Field) RestrictFromLeftAndRight(leftTaken int, righTaken int) (Field, bool) {
	f, leftOk := f.RestritLeft(leftTaken)
	f, rightOk := f.RestritRight(righTaken)
	return f, leftOk && rightOk
}

func (f Field) RestritLeft(leftTaken int) (Field, bool) {
	if leftTaken >= len(f.layout) {
		return f, false
	}
	if leftTaken < 0 { // kan försöka lägga in betydelse i negativa tal men inte just nu
		return f, true
	}
	// hantering av mellanrummet
	if !(f.layout[leftTaken] == springs.Operational() || f.layout[leftTaken] == springs.Unknown()) {
		return f, false
	} else {
		f.layout[leftTaken] = springs.Operational()
	}
	possible := true
	for i := leftTaken - 1; i >= 0; i-- {
		loopSpring := f.layout[i]
		if loopSpring == springs.Damaged() {
			continue
		} else if loopSpring == springs.Unknown() {
			f.layout[i] = springs.Damaged()
		} else {
			possible = false
			break
		}
	}
	return f, possible
}

func (f Field) RestritRight(rightTaken int) (Field, bool) {
	if rightTaken >= len(f.layout) {
		return f, false
	}
	if rightTaken < 0 { // kan försöka lägga in betydelse i negativa tal men inte just nu
		return f, true
	}
	// hantering av mellanrummet
	lowestIndex := len(f.layout) - 1 - rightTaken
	if !(f.layout[lowestIndex] == springs.Operational() || f.layout[lowestIndex] == springs.Unknown()) {
		return f, false
	} else {
		f.layout[lowestIndex] = springs.Operational()
	}
	possible := true
	for i := lowestIndex + 1; i < len(f.layout); i++ {
		loopSpring := f.layout[i]
		if loopSpring == springs.Damaged() {
			continue
		} else if loopSpring == springs.Unknown() {
			f.layout[i] = springs.Damaged()
		} else {
			possible = false
			break
		}
	}
	return f, possible
}

func (f Field) Len() int {
	return len(f.layout)
}
