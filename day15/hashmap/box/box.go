package box

type Box struct {
	lenses []int
}

func Init(lenses []int) Box {
	return Box{lenses: lenses}
}

func (b *Box) Add(lens int) {
	for i := range b.lenses {
		if b.lenses[i] == lens {
			return
		}
	}
	b.lenses = append(b.lenses, lens)
}

func (b *Box) Remove(lens int) {
	for i := range b.lenses {
		if b.lenses[i] == lens { // sista elementet?
			b.lenses = append(b.lenses[:i], b.lenses[i+1:]...)
			return
		}
	}
}
