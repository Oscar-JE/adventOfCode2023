package policy

import (
	"day8/direction"
)

type Policy struct {
	sequence     string
	currentIndex int
}

func Init(seq string) Policy {
	for _, r := range seq {
		if r != 'R' && r != 'L' {
			panic("faulty runes in sequence string for policy")
		}
	}
	return Policy{sequence: seq, currentIndex: 0}
}

func (p *Policy) GetNext() direction.Direction {
	dir := p.currentDirection()
	p.incrementIndex()
	return dir
}

func (p *Policy) incrementIndex() {
	p.currentIndex = (p.currentIndex + 1) % len(p.sequence)
}

func (p Policy) currentDirection() direction.Direction {
	dir := direction.LEFT
	if p.sequence[p.currentIndex] == 'R' {
		dir = direction.RIGHT
	}
	return dir
}

func (p *Policy) GetNextAndindex() (direction.Direction, int) {
	dir := p.currentDirection()
	index := p.currentIndex
	p.incrementIndex()
	return dir, index
}

func (p *Policy) Recett() {
	p.currentIndex = 0
}
