package environment

import "day16/vec"

var north vec.Vec2d = vec.Init(-1, 0)
var east vec.Vec2d = vec.Init(0, 1)
var south vec.Vec2d = vec.Init(1, 0)
var west vec.Vec2d = vec.Init(0, -1)

var directions = []vec.Vec2d{north, east, south, west}
