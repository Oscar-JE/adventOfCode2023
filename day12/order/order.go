package order

type Order struct {
	ordering []int
}

func Init(ordering []int) Order {
	return Order{ordering: ordering}
}

func (o Order) Eq(other Order) bool {
	if len(o.ordering) != len(other.ordering) {
		return false
	}
	eq := true
	for i := range o.ordering {
		eq = eq && o.ordering[i] == other.ordering[i]
	}
	return eq
}
