package environment

import (
	"day16/matrix"
	"day16/vec"
	"day17/dijkstras"
	"strconv"
)

type cost struct {
	v int
}

func (c cost) toInt() int {
	return c.v
}

func (c cost) String() string {
	return strconv.Itoa(c.v)
}

type Environment struct {
	costs matrix.Matrix[cost]
}

func InitEnv(data []int, rows int, cols int) Environment {
	costs := []cost{}
	for _, val := range data {
		costs = append(costs, cost{val})
	}
	matCosts := matrix.Init(costs, rows, cols)
	return Environment{matCosts}
}

func (e Environment) NextStates(current State) []State {
	unRestricted := current.NextPossibleStates()
	restricted := []State{}
	for _, s := range unRestricted {
		if e.inArea(s.position) {
			restricted = append(restricted, s)
		}
	}
	return restricted
}

func (e Environment) TransFer(current State) []dijkstras.StateAndCost[State] {
	nextStastes := e.NextStates(current)
	statesAndCosts := []dijkstras.StateAndCost[State]{}
	for _, state := range nextStastes {
		cost := e.costs.Get(state.position.GetX(), state.position.GetY())
		statesAndCosts = append(statesAndCosts, dijkstras.InitStateAndCost[State](state, cost.toInt()))
	}
	return statesAndCosts
}

func (e Environment) inArea(p vec.Vec2d) bool {
	x := p.GetX()
	y := p.GetY()
	return 0 <= x && x < e.costs.GetRows() && 0 <= y && y < e.costs.GetCols()
}

func (e Environment) Cost(position vec.Vec2d) int {
	return e.costs.Get(position.GetX(), position.GetY()).v
}

func (e Environment) String() string {
	return e.costs.String()
}
