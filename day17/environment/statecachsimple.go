package environment

import (
	"day17/dijkstras"
	"math"
)

type SimpleCach struct {
	l []dijkstras.StateAndCost[State]
}

func InitSimpleCach() SimpleCach {
	return SimpleCach{l: []dijkstras.StateAndCost[State]{}}
}

func (c SimpleCach) Has(s State) bool {
	for _, el := range c.l {
		if el.GetState() == s {
			return true
		}
	}
	return false
}

func (c *SimpleCach) SetValue(s State, cost int) {
	for index, el := range c.l {
		if el.GetState() == s {
			c.l[index] = dijkstras.InitStateAndCost(s, cost)
		}
	}
	c.l = append(c.l, dijkstras.InitStateAndCost[State](s, cost))
}

func (c SimpleCach) GetValue(s State) int {
	for _, el := range c.l {
		if el.GetState() == s {
			return el.GetCost()
		}
	}
	return math.MaxInt
}

func (c SimpleCach) ProcentageFilled() float64 {
	return 0.0
}
