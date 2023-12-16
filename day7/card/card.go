package card

import "strings"

const Ordering = "J23456789TQKA"

type Card struct {
	Rep string
}

func (c Card) Value() int {
	return strings.Index(Ordering, c.Rep)
}

func Init(rep string) Card {
	return Card{Rep: rep}
}

func (c Card) Eq(other Card) bool {
	return c.Value() == other.Value()
}
