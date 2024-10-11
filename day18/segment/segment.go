package segment

import (
	"day16/vec"
	"day18/instruction"
)

type Segment struct {
	start     vec.Vec2d
	direction vec.Vec2d
	maxLen    int
}

func Init(start vec.Vec2d, len int, direction vec.Vec2d) Segment {
	return Segment{start: start, direction: direction, maxLen: len}
}

func (s Segment) GetDirection() vec.Vec2d {
	return s.direction
}

func InitFromInstruction(start vec.Vec2d, inst instruction.Instruction) Segment {
	return Segment{start: start, direction: inst.GetDirection(), maxLen: inst.GetDistance()}
}

func (s Segment) IsHorizontal() bool {
	return s.direction == vec.Init(0, 1) || s.direction == vec.Init(0, -1)
}

func (s Segment) Translate(translationDirection vec.Vec2d) []vec.Vec2d {
	points := []vec.Vec2d{}
	startPrim := vec.Add(s.start, translationDirection)
	for i := 1; i <= s.maxLen; i++ {
		startPrim = vec.Add(s.direction, startPrim)
		points = append(points, startPrim)
	}
	return points
}

func (s Segment) In(point vec.Vec2d) bool {
	referencePoint := vec.Subtract(point, s.start)
	if referencePoint.GetX() == 0 {
		return within(s.direction.GetY(), s.maxLen, referencePoint.GetY())
	}
	if referencePoint.GetY() == 0 {
		return within(s.direction.GetX(), s.maxLen, referencePoint.GetX())
	}
	return false
}

func within(sign int, length int, controlledValue int) bool {
	positiveControl := controlledValue * sign
	return 0 < positiveControl && positiveControl <= length
}

func (s Segment) MinRow() int {
	return min(s.start.GetX(), s.end().GetX())
}

func (s Segment) MaxRow() int {
	return max(s.start.GetX(), s.end().GetX())
}

func (s Segment) MinCol() int {
	return min(s.start.GetY(), s.end().GetY())
}

func (s Segment) MaxCol() int {
	return max(s.start.GetY(), s.end().GetY())
}

func (s Segment) end() vec.Vec2d {
	return vec.Add(s.start, s.direction.ScalarMultiplication(s.maxLen))
}

func (s Segment) NrOfPoints() int {
	return s.maxLen
}
