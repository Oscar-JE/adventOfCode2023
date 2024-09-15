package priorityqueue

import "testing"

func TestUpdateWithMatch(t *testing.T) {
	pQ := InitWithValues([]element[int]{element[int]{1, 2}, element[int]{3, 4}})
	pQ.Update(3, 5)
	expected := InitWithValues([]element[int]{element[int]{1, 2}, element[int]{3, 5}})
	if !pQ.Eq(expected) {
		t.Errorf("update with match is faulty")
	}
}

func TestUpdateWithNoMatch(t *testing.T) {
	pQ := InitWithValues([]element[int]{element[int]{1, 2}, element[int]{3, 4}})
	pQ.Update(10, 3)
	expected := InitWithValues([]element[int]{element[int]{1, 2}, element[int]{10, 3}, element[int]{3, 4}})
	if !pQ.Eq(expected) {
		t.Errorf("update with no match is faulty")
	}
}

func TestPop(t *testing.T) {
	pQ := InitWithValues([]element[int]{element[int]{1, 1}, element[int]{1, 1}})
	val, prio := pQ.Pop()
	if val != 1 {
		t.Errorf("wrong value in popped element")
	}
	if prio != 1 {
		t.Errorf("wrong prie of popped element")
	}
	if len(pQ.elements) != 1 {
		t.Errorf("wrong number of remaining elements after pop")
	}
}

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

func TestFindINdexOddSameValues(t *testing.T) {
	list := []int{8, 8, 8}
	value := 10
	expected := 3
	res := findIndexInSortedList(list, value)
	if res != expected {
		t.Errorf("res was %d but should bee %d", res, expected)
	}
}

func TestFindINdexOddSameValues6(t *testing.T) {
	list := []int{7, 10}
	value := 8
	expected := 1
	res := findIndexInSortedList(list, value)
	if res != expected {
		t.Errorf("res was %d but should bee %d", res, expected)
	}
}

func TestFindNearestGap(t *testing.T) {
	list := []int{7, 10}
	value := 8
	expected := 0
	res := findNearestGap(list, value)
	if res != expected {
		t.Errorf("res was %d but should bee %d", res, expected)
	}
}

func TestFindNearestGapLonger(t *testing.T) {
	list := []int{7, 7, 10}
	value := 8
	expected := 1
	res := findNearestGap(list, value)
	if res != expected {
		t.Errorf("res was %d but should bee %d", res, expected)
	}
}
