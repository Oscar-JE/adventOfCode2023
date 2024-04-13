package point

type Vec struct {
	x int
	y int
}

func Init(x int, y int) Vec {
	return Vec{x: x, y: y}
}

func (v Vec) GetX() int {
	return v.x
}

func (v Vec) GetY() int {
	return v.y
}

func (v Vec) Add(other Vec) Vec {
	return Init(v.x+other.x, v.y+other.y)
}

func (v Vec) Negate() Vec {
	return Init(-v.x, -v.y)
}

func (v *Vec) Rotate90counterclockwise() {
	newX := -v.y
	newY := v.x
	v.x = newX
	v.y = newY
}

func (v Vec) Less(other Vec) bool {
	if v.x < other.x {
		return true
	} else if v.x == other.x && v.y < other.y {
		return true
	}
	return false
}
