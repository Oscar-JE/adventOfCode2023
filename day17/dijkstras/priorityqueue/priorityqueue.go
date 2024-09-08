package priorityqueue

type element[T any] struct {
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
	lowerIndex := 0
	upperIndex := len(list) - 1
	for lowerIndex < upperIndex {
		mid := (lowerIndex + upperIndex) / 2
		midValue := list[mid]
		if v < midValue {
			upperIndex = mid
		}
		if midValue <= v {
			lowerIndex = mid
		}
	}
	return lowerIndex
}
