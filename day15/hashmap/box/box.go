package box

import (
	"day15/hashmap/lens"
	"fmt"
)

type Box struct {
	lenses []lens.Lens
}

func Init(lenses []lens.Lens) Box {
	return Box{lenses: lenses}
}

func (b *Box) Add(lens lens.Lens) {
	for i := range b.lenses {
		if b.lenses[i].Label() == lens.Label() {
			b.lenses[i] = lens
			return
		}
	}
	b.lenses = append(b.lenses, lens)
}

func (b *Box) Remove(lens lens.Lens) {
	for i := range b.lenses {
		if b.lenses[i].Label() == lens.Label() {
			b.lenses = append(b.lenses[:i], b.lenses[i+1:]...)
			return
		}
	}
}

func (b Box) FocusingPower() int {
	sum := 0
	for index, lens := range b.lenses {
		sum += (index + 1) * lens.FocalLength()
	}
	return sum
}

func (b Box) String() string {
	rep := ""
	for _, lens := range b.lenses {
		rep += fmt.Sprint(lens) + " "
	}
	return rep
}

func (b Box) IsEmpty() bool {
	return len(b.lenses) == 0
}
