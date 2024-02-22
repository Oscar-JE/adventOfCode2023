package direction

type Direction int

const (
	LEFT  Direction = 0
	RIGHT Direction = 1
)

func (d Direction) String() string {
	if d == 0 {
		return "L"
	} else {
		return "H"
	}
}
