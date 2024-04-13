package tile

import "day10/point"

type Bend struct {
	orientation BendOrientation
	position    point.Vec
}

func NortEastBend(p point.Vec) Bend {
	return Bend{NortEast(), p}
}

func NortWestBend(p point.Vec) Bend {
	return Bend{NortWest(), p}
}

func SoutWestBend(p point.Vec) Bend {
	return Bend{SoutWest(), p}
}

func EastSouthBend(p point.Vec) Bend {
	return Bend{EastSouth(), p}
}

func (b Bend) Ways() point.VecSet {
	m := b.movement()
	return m.Translate(b.position)
}

func (b Bend) GetPos() point.Vec {
	return b.position
}

func (b Bend) movement() point.VecSet {
	set := point.EmptySet() // default till NortEast
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
