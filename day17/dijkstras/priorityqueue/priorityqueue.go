package priorityqueue

type element[T any] struct { // nu rör vi oss äntligen bort från any. Det måste kunna göra en eq
	state    T
	priority int
}

type PriorityQueue[T any] struct {
	elements []element[T]
}

func Init[T any]() PriorityQueue[T] {
	return PriorityQueue[T]{}
}

func InitWithValues[T any](elements []element[T]) PriorityQueue[T] {
	return PriorityQueue[T]{elements: elements}
}

func (p PriorityQueue[T]) HasElement() bool {
	return len(p.elements) > 0
}

func (p *PriorityQueue[T]) Pop() (T, int) {
	// poppa från listan dela upp i interna grejer och returnera
	popped := p.elements[0]
	p.elements = p.elements[1:]
	return popped.state, popped.priority
}

func (p *PriorityQueue[T]) Update(s T, prio int) {
	// skriver ett test så att vi vet att denna ska skrivas
}

func (p *PriorityQueue[T]) Add(s T, prio int) {
	insertIndex := p.findInsertionIndex(prio)
	tail := append([]element[T]{{state: s, priority: prio}}, p.elements[insertIndex:]...)
	p.elements = append(p.elements[:insertIndex], tail...)
}

func (p PriorityQueue[T]) findInsertionIndex(prio int) int {
	priorities := []int{}
	for _, elem := range p.elements {
		priorities = append(priorities, elem.priority)
	}
	return findIndexInSortedList(priorities, prio)
}

func findIndexInSortedList(list []int, v int) int {
	if len(list) == 0 {
		return 0
	}
	gapIndex := findNearestGap(list, v)
	indexCandidate := gapIndex + 1
	if indexCandidate == 1 && v < list[0] {
		return 0
	}
	if indexCandidate == len(list)-1 && v > list[len(list)-1] {
		return len(list)
	}
	return indexCandidate
}

func findNearestGap(list []int, v int) int {
	nrOfGaps := len(list) - 1
	if nrOfGaps <= 0 {
		return 0
	}
	minIndex := 0            // firstGapIndex
	maxIndex := nrOfGaps - 1 // lastGapIndex
	for minIndex < maxIndex {
		midIndex := (minIndex + maxIndex) / 2
		if list[midIndex] <= v && v <= list[midIndex] {
			return midIndex
		}
		if v < list[midIndex] {
			maxIndex = midIndex - 1
		}
		if list[midIndex+1] < v {
			minIndex = midIndex + 1
		}
	}
	return minIndex
}

func (p PriorityQueue[T]) Eq(other PriorityQueue[T]) bool {
	if len(p.elements) != len(other.elements) {
		return false
	}
	same := true
	for i := range p.elements {
		pE := p.elements[i]
		oE := other.elements[i]
		same = same && pE == oE
	}
}
