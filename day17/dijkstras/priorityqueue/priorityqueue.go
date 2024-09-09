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
	return findIndexInSortedList2(priorities, prio)
}

func findIndexInSortedList(list []int, v int) int {
	if len(list) == 0 || len(list) == 1 || v < list[0] {
		return 0
	}
	for i := 0; i < len(list)-2; i++ {
		if list[i] <= v && v <= list[i+1] {
			return i
		}
	}
	if v > list[len(list)-1] {
		return len(list)
	} else {
		return len(list) - 1
	}
}

func findIndexInSortedList2(list []int, v int) int {
	return findIndexInSortedListInternal(list, v, 0, len(list)-1)
}

func findIndexInSortedListInternal(list []int, v int, startIndex int, slutIndex int) int {
	if startIndex == slutIndex {
		if len(list) > 0 {
			if v < list[startIndex] {
				return startIndex
			} else {
				return startIndex + 1
			}
		}
		return startIndex
	}
	midIndex := (startIndex + slutIndex) / 2
	if list[midIndex] <= v && v <= list[midIndex+1] {
		return midIndex
	}
	if v < list[midIndex] {
		return findIndexInSortedListInternal(list, v, startIndex, midIndex)
	}
	return findIndexInSortedListInternal(list, v, midIndex+1, slutIndex)
}
