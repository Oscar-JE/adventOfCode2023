package linearalgebra

type Vec struct {
	x int
	y int
}

func InitVec(x int, y int) Vec {
	return Vec{x: x, y: y}
}

func (v Vec) GetX() int {
	return v.x
}

func (v Vec) GetY() int {
	return v.y
}

func Add(v1 Vec, v2 Vec) Vec {
	return InitVec(v1.x+v2.x, v1.y+v2.y)
}

func (v Vec) Multiply(multiple int) Vec {
	return InitVec(v.x*multiple, v.y*multiple)
}

func ManhattanDistans(v1 Vec, v2 Vec) int {
	return abs(v2.x-v1.x) + abs(v2.y-v1.y)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
