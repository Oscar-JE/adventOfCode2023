package tile

import "day10/point"

type Bend struct {
	orientation BendOrientation
}

func nortEastBend() Bend {
	return Bend{NortEast()}
}

func nortWestBend() Bend {
	return Bend{NortWest()}
}

func soutWestBend() Bend {
	return Bend{SoutWest()}
}

func eastSouthBend() Bend {
	return Bend{EastSouth()}
}

func (b Bend) moves() point.VecSet {
	set := point.EmptySet()
	set.Add(point.Init(-1, 0))
	set.Add(point.Init(0, 1))
	for i := 0; i < b.orientation.Key; i++ {
		set.Rotate90counterclockwise()
	}
	return set
}

type BendOrientation struct { // får hålla koll på rotattionerna senare
	Key int
}

func NortEast() BendOrientation {
	return BendOrientation{0}
}

func NortWest() BendOrientation {
	return BendOrientation{1}
}

func SoutWest() BendOrientation {
	return BendOrientation{2}
}

func EastSouth() BendOrientation {
	return BendOrientation{3}
}
