package dijkstras

import (
	"day17/dijkstras/priorityqueue"
)

// borde inte vara comparable
type Environment[E comparable] interface {
	TransFer(state E) []StateAndCost[E]
}

type StateAndCost[E comparable] struct {
	state E
	cost  int
}

func InitStateAndCost[E comparable](state E, cost int) StateAndCost[E] {
	return StateAndCost[E]{state: state, cost: cost}
}

func (s StateAndCost[E]) GetState() E {
	return s.state
}

func (s StateAndCost[E]) GetCost() int {
	return s.cost
}

type StateCach[E comparable] interface {
	SetValue(state E, pathCost int)
	GetValue(state E) int
	Has(state E) bool
}

type Dijkstras[E comparable] struct {
	cach    StateCach[E]
	env     Environment[E]
	visited priorityqueue.PriorityQueue[E]
}

func Init[E comparable](env Environment[E], cach StateCach[E]) Dijkstras[E] {
	visited := priorityqueue.Init[E]()
	return Dijkstras[E]{cach: cach, env: env, visited: visited}
}

func (d *Dijkstras[E]) findPaths(startState E, startCost int) {
	d.visited.Add(startState, startCost)
	for d.visited.HasElement() {
		state, trueDist := d.visited.Pop()
		d.cach.SetValue(state, trueDist)
		nextGenStateAndCost := d.env.TransFer(state)
		for _, stateAndCost := range nextGenStateAndCost {
			if d.cach.Has(stateAndCost.GetState()) {
				continue
			}
			costCandidate := trueDist + stateAndCost.cost
			d.visited.Update(stateAndCost.GetState(), costCandidate)
		}
	}
}

func (this Dijkstras[E]) SmallestDist(endStates []E, startState E, startCost int) int {
	this.findPaths(startState, startCost)
	if len(endStates) == 0 {
		return 0
	}
	min := this.cach.GetValue(endStates[0])
	for i := 1; i < len(endStates); i++ {
		cached := this.cach.GetValue(endStates[i])
		if cached < min {
			min = cached
		}
	}
	return min
}
