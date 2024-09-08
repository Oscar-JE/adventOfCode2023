package priorityqueue

import "testing"

func TestFindIndexEmpty(t *testing.T) {
	sorted := []int{}
	value := 2
	index := findIndexInSortedList(sorted, value)
	if index != 0 {
		t.Errorf("correkt insert index in empty list is always zero")
	}
}

func TestFindIndexEevenLowest(t *testing.T) {
	sorted := []int{1, 3}
	i1 := findIndexInSortedList(sorted, 0)
	if i1 != 0 {
		t.Errorf("i1 shoud be 0 but was %d", i1)
	}
}

func TestFindIndexEvenEqualLowest(t *testing.T) {
	sorted := []int{1, 3}
	i2 := findIndexInSortedList(sorted, 1)
	if !(i2 == 0 || i2 == 1) {
		t.Errorf("i2 should be 0 or 1 but was %d", i2)
	}

}

func TestFindIndexEevenMiddle(t *testing.T) {
	sorted := []int{1, 3}
	i3 := findIndexInSortedList(sorted, 2)

	if i3 != 1 {
		t.Errorf("i3 should be 1 but was %d", i3)
	}
}

func TestFindIndexEvenEqualhigest(t *testing.T) {
	sorted := []int{1, 3}
	i4 := findIndexInSortedList(sorted, 3)
	if !(i4 == 1 || i4 == 2) {
		t.Errorf("i4 should either be 1 or 2 but was %d", i4)
	}
}

func TestFindIndexEvenNumberBigger(t *testing.T) {
	sorted := []int{1, 3}
	i5 := findIndexInSortedList(sorted, 4)
	if i5 != 2 {
		t.Errorf("i5 should be 2 but was %d", i5)
	}

}
