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

func Trun90upp(v1 Vec2d) Vec2d { /// kanske inte kommer behöva denna funktion väntar med att skriva den
	// se right hand rule
	return Init(0, 0) // enbart för att kunna köra och skriva våra tester
}
