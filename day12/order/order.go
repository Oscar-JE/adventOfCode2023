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

func (o Order) Len() int {
	return len(o.ordering)
}

func (o *Order) Pop() int {
	if len(o.ordering) == 0 {
		return -1 // bättre med error här kanske
	}
	element := o.ordering[0]
	o.ordering = o.ordering[1:]
	return element
}

func (o *Order) Unfold(times int) {
	length := o.Len()
	for i := 0; i < times; i++ {
		for j := 0; j < length; j++ {
			o.ordering = append(o.ordering, o.ordering[j])
		}
	}
}
