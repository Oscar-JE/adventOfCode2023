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
	dir := direction.LEFT
	if p.sequence[p.currentIndex] == 'R' {
		dir = direction.RIGHT
	}
	p.currentIndex = (p.currentIndex + 1) % len(p.sequence)
	return dir
}
