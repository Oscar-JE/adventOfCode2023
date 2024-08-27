package vec

type Vec2d struct {
	x int
	y int
}

func Init(x int, y int) Vec2d {
	return Vec2d{x: x, y: y}
}

func Add(v1 Vec2d, v2 Vec2d) Vec2d {
	return Vec2d{x: v1.x + v2.x, y: v1.y + v2.y}
}

func (v Vec2d) GetX() int {
	return v.x
}

func (v Vec2d) GetY() int {
	return v.y
}

func (v Vec2d) ScalarMultiplication(a int) Vec2d {
	return Init(a*v.GetX(), a*v.GetY())
}

func (v Vec2d) ScalarDevision(a int) Vec2d {
	return Init(v.GetX()/a, v.GetY()/a)
}

func (v Vec2d) Flip() Vec2d {
	return v.ScalarMultiplication(-1)
}

func (v Vec2d) InnerProduct(other Vec2d) int {
	return v.x*other.x + v.y*other.y
}

func (v Vec2d) Orthogonal(other Vec2d) bool {
	return v.InnerProduct(other) == 0
}
