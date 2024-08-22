package beam

import "day16/vec"

type Beam struct {
	position  vec.Vec2d
	direction vec.Vec2d
}

func Init(position vec.Vec2d, direction vec.Vec2d) Beam {
	return Beam{position: position, direction: direction}
}

func (b *Beam) Travel() {
	b.position = vec.Add(b.position, b.direction)
}

func (b Beam) GetPosition() vec.Vec2d {
	return b.position
}

func (b Beam) GetDirection() vec.Vec2d {
	return b.direction
}
