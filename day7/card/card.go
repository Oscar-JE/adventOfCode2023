package card

import "strings"

const Ordering = "23456789TJQKA"

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
