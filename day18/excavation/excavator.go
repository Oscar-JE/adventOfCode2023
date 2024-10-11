package excavation

import (
	"day16/vec"
	"day18/instruction"
	"day18/segment"
)

type excavator struct {
	position vec.Vec2d
}

func initExcavator() (excavator, vec.Vec2d) {
	return excavator{vec.Init(0, 0)}, vec.Init(0, 0)
}

func (e *excavator) dig(inst instruction.Instruction) []vec.Vec2d {
	direction := inst.GetDirection()
	nrSteps := inst.GetDistance()
	excevated := []vec.Vec2d{}
	for i := 0; i < nrSteps; i++ {
		e.position = vec.Add(e.position, direction)
		excevated = append(excevated, e.position)
	}
	return excevated
}

func Excavate(instructions []instruction.Instruction) []vec.Vec2d {
	exc, _ := initExcavator()
	holes := []vec.Vec2d{}
	for _, instruction := range instructions {
		newlyDugg := exc.dig(instruction)
		holes = append(holes, newlyDugg...)
	}
	return holes
}

func ExcavateInSegments(instructions []instruction.Instruction) []segment.Segment {
	currentPoint := vec.Init(0, 0)
	segments := []segment.Segment{}
	for _, inst := range instructions {
		seg := segment.InitFromInstruction(currentPoint, inst)
		segments = append(segments, seg)
		currentPoint = vec.Add(currentPoint, inst.GetDirection().ScalarMultiplication(inst.GetDistance()))
	}
	return segments
}
