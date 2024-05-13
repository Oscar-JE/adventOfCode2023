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

func (o Order) RolingRsize(left int, right int) Order {
	// häri gömmer vi logiken medvilken ordning och liknande
	orginalLista := o.ordering
	// expanderande operationer först
	res := orginalLista
	if left < 0 {
		expandLeft := []int{}
		for i := -1; i >= left; i-- {
			j := (i + 10*o.Len()) % o.Len() // 10 an för att undvika negativa tal
			expandLeft = append([]int{orginalLista[j]}, expandLeft...)
		}
		res = append(expandLeft, res...)
	}
	if right > 0 {
		expandRight := []int{}
		for i := range right {
			j := i % o.Len()
			expandRight = append(expandRight, o.ordering[j])
		}
		res = append(res, expandRight...)
	}
	if left >= 0 {
		res = res[left:]
	}
	if right <= 0 {
		res = res[:len(res)+right] // frågan är om det blir en minus etta här
	}

	// lite längre men jag tycker att den är lättare att förstå
	return Init(res)

}
