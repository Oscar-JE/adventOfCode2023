package springs

type Condition struct {
	condition int
}

func Operational() Condition {
	return Condition{0}
}

func Damaged() Condition {
	return Condition{1}
}

func Unknown() Condition {
	return Condition{2}
}

func (c Condition) String() string {
	runeMapping := []rune{'.', '#', '?'}
	return string(runeMapping[c.condition])
}

func Init(r rune) Condition {
	if r == '.' {
		return Operational()
	} else if r == '#' {
		return Damaged()
	} else {
		return Unknown()
	}
}
