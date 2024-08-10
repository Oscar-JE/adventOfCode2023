package pice

type Pice int

const Empty = 0
const Fixed = 1
const Movable = 2

// kan göras snyggare
func Init(nr int) Pice {
	if nr == Fixed {
		return Pice(Fixed)
	} else if nr == Movable {
		return Pice(Movable)
	} else {
		return Pice(Empty)
	}
}

func InitFromRune(r rune) Pice {
	if r == '#' {
		return Pice(Fixed)
	} else if r == 'O' {
		return Pice(Movable)
	} else {
		return Pice(Empty)
	}
}

func (p Pice) IsEmpty() bool {
	return p == Empty
}

func (p Pice) IsFixed() bool {
	return p == Fixed
}

func (p Pice) IsMovavle() bool {
	return p == Movable
}

func (p Pice) ToRune() rune {
	if p == Fixed {
		return '#'
	} else if p == Movable {
		return 'O'
	} else {
		return '.'
	}
}

func (p Pice) String() string {
	return string(p.ToRune())
}
