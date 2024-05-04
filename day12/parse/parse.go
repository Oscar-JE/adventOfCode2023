package parse

import (
	"day12/field"
	"day12/order"
	"strconv"
	"strings"
)

func Line(s string) (field.Field, order.Order) {
	fieldRepAndOrderRep := strings.Split(s, " ")
	fieldRep := fieldRepAndOrderRep[0]
	orderRep := fieldRepAndOrderRep[1]
	runes := []rune(fieldRep)
	orderList := commaSeparatedToLit(orderRep)
	return field.Init(runes), order.Init(orderList)
}

func commaSeparatedToLit(s string) []int {
	intRep := strings.Split(s, ",")
	ints := []int{}
	for i := range intRep {
		a, err := strconv.Atoi(intRep[i])
		if err != nil {
			panic("error occured while parsing inputlist")
		}
		ints = append(ints, a)
	}
	return ints
}
