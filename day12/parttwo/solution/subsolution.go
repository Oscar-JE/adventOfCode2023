package solution

import "day12/order"

type SubProblem struct {
	left     Connection
	right    Connection
	subOrder order.Order
}

func InitSubProblem(left Connection, subOrder order.Order, right Connection) SubProblem {
	return SubProblem{left: left, subOrder: subOrder, right: right}
}

func (s SubProblem) SolutionKey() SolutionKey {
	return SolutionKey{LeftDemand: s.left.RightDemand(), RightDemand: s.right.LeftDemand(), SubOrder: s.subOrder}
}
