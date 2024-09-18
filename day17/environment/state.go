package environment

import (
	"day16/vec"
)

const maxStepInline int = 3

type State struct {
	position  vec.Vec2d
	direction vec.Vec2d
	stepsTake int
}

func Init(position vec.Vec2d, direction vec.Vec2d, stepsTaken int) State {
	return State{position: position, direction: direction, stepsTake: stepsTaken}
}

func (s State) NextPossibleStates() []State {
	possibleNext := []State{}
	for _, dir := range directions {
		loopNext := s
		if loopNext.direction.InnerProduct(dir) == -1 {
			continue
		}
		if loopNext.direction == dir {
			if loopNext.stepsTake != maxStepInline-1 {
				loopNext.stepsTake++
			} else {
				continue
			}

		} else if loopNext.direction != dir {
			loopNext.stepsTake = 0
		}
		loopNext.direction = dir
		loopNext.position = vec.Add(loopNext.position, loopNext.direction)
		possibleNext = append(possibleNext, loopNext)
	}
	return possibleNext
}

func (s State) hashCode(rows int) int {
	positionHash := s.position.GetX()*rows + s.position.GetY()
	directionHash := directionHash(s.direction)
	stepsInLineHash := s.stepsTake
	hash := positionHash*(len(directions)*maxStepInline) + directionHash*maxStepInline + stepsInLineHash
	return hash
}

func WinnigStates(end vec.Vec2d) []State {
	ret := []State{}
	for step := range 3 {
		for _, dir := range directions {
			ret = append(ret, State{stepsTake: step, position: end, direction: dir})
		}
	}
	return ret
}

func StartState() State {
	return State{stepsTake: 0, position: vec.Init(0, 0), direction: east}
}
