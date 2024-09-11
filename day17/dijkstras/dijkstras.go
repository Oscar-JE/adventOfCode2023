package dijkstras

import "day17/dijkstras/priorityqueue"

// borde inte vara any
type Environment[E any] interface {
	TransFer(state E) []StateAndCost[E]
}

type StateAndCost[E any] interface {
	GetState() E
	GetCost() int
}

type StateCach[E any] interface {
	SetValue(state E, pathCost int)
	GetValue(state E) int
}

type Dijkstras[E any] struct {
	cach    StateCach[E]
	env     Environment[E]
	visited priorityqueue.PriorityQueue[E]
}

func Init[E any](env Environment[E], cach StateCach[E]) Dijkstras[E] {
	visited := priorityqueue.Init[E]()
	return Dijkstras[E]{cach: cach, env: env, visited: visited}
}

func (d *Dijkstras[E]) findPaths(startState E, startCost int) {
	d.cach.SetValue(startState, startCost)
	nextStates := d.env.TransFer(startState)
	for _, stateAndCost := range nextStates {
		d.visited.Add(stateAndCost.GetState(), stateAndCost.GetCost())
	}
	for d.visited.HasElement() { // kan det här leda till en förtidig stop ?

	}
}
