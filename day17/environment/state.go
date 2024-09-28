package environment

import (
	"day16/vec"
)

const maxStepInLinePart1 int = 3
const minStepInLinePart1 int = 0

const maxStepInLinePart2 int = 10
const minStepInLinePart2 int = 4

type State struct {
	position  vec.Vec2d
	direction vec.Vec2d
	stepsTake int
	poll      policy
}

type policy interface {
	NextPossibleStates(State) []State
	maxNrSteps() int // har denna gått sönder helt och fullt?
}

type policy1 struct {
}

func (p policy1) NextPossibleStates(s State) []State {
	return NextStatesWithLimits(s, minStepInLinePart1, maxStepInLinePart1)
}

func (p policy1) maxNrSteps() int {
	return maxStepInLinePart1
}

type policy2 struct {
}

func (p policy2) maxNrSteps() int {
	return maxStepInLinePart2
}

func (p policy2) NextPossibleStates(s State) []State {
	return NextStatesWithLimits(s, minStepInLinePart2, maxStepInLinePart2)
}

func NextStatesWithLimits(s State, lowerStepLimit int, upperStepLimit int) []State {
	possibleNext := []State{}
	if s.stepsTake < lowerStepLimit {
		nextState := s
		nextState.position = vec.Add(nextState.position, nextState.direction)
		nextState.stepsTake++
		possibleNext = append(possibleNext, nextState)
		return possibleNext
	}
	for _, dir := range directions {
		loopNext := s
		if loopNext.direction.InnerProduct(dir) == -1 {
			continue
		}
		if loopNext.direction == dir {
			if loopNext.stepsTake != upperStepLimit {
				loopNext.stepsTake++
			} else {
				continue
			}

		} else if loopNext.direction != dir {
			loopNext.stepsTake = 1
		}
		loopNext.direction = dir
		loopNext.position = vec.Add(loopNext.position, loopNext.direction)
		possibleNext = append(possibleNext, loopNext)
	}
	return possibleNext
}

func Init(position vec.Vec2d, direction vec.Vec2d, stepsTaken int) State {
	return State{position: position, direction: direction, stepsTake: stepsTaken, poll: policy1{}}
}

func Init2(position vec.Vec2d, direction vec.Vec2d, stepsTaken int) State {
	return State{position: position, direction: direction, stepsTake: stepsTaken, poll: policy2{}}
}

func (s State) NextPossibleStates() []State {
	return s.poll.NextPossibleStates(s)
}

func (s State) hashCode(rows int) int {
	positionHash := s.position.GetX()*rows + s.position.GetY()
	directionHash := directionHash(s.direction)
	stepsInLineHash := s.stepsTake
	hash := positionHash*(len(directions)*s.poll.maxNrSteps()) + directionHash*s.poll.maxNrSteps() + stepsInLineHash
	return hash
}

func (s State) MaxHash(rows int, cols int) int {
	return rows * cols * len(directions) * s.poll.maxNrSteps()
}

func WinningStates(end vec.Vec2d) []State {
	ret := []State{}
	for step := range maxStepInLinePart1 {
		for _, dir := range directions {
			ret = append(ret, Init(end, dir, step))
		}
	}
	return ret
}

func WinningStates2(end vec.Vec2d) []State {
	ret := []State{}
	for step := range maxStepInLinePart2 {
		for _, dir := range directions {
			ret = append(ret, Init2(end, dir, step))
		}
	}
	return ret
}

func StartState() State {
	return Init(vec.Init(0, 0), east, 0)
}

func StartState2() State {
	return Init2(vec.Init(0, 0), east, 0)
}
