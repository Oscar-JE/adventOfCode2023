package environment

import "day16/vec"

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
			if loopNext.stepsTake != 2 {
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

// kan invertera denna funktion också för att köra uträkningen snabbare
// men det borde inte behövas för att köra en weighted dixtras
