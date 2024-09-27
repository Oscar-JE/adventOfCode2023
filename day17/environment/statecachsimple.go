package environment

import (
	"day16/vec"
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

func (c SimpleCach) FindLowestWithPosition(position vec.Vec2d) int {
	allMatchedValues := []int{}
	for _, el := range c.l {
		if el.GetState().position == position {
			allMatchedValues = append(allMatchedValues, el.GetCost())
		}
	}
	if len(allMatchedValues) == 0 {
		return 1000
	}
	minimum := allMatchedValues[0]
	for i := 1; i < len(allMatchedValues); i++ {
		if minimum > allMatchedValues[i] {
			minimum = allMatchedValues[i]
		}
	}
	return minimum
}
