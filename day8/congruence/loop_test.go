package congruence

import (
	"testing"
)

func assertEq(expected TailedLoopMultipleEndPoints, result TailedLoopMultipleEndPoints, t *testing.T) {
	if !(expected.Eq(result)) {
		t.Errorf("\nExpected: " + expected.String() + "\nResult:   " + result.String())
	}
}

func TestCutWithSameLoop(t *testing.T) {
	loop := InitTailedLoop(3, 4, []int{0, 1, 2, 3})
	loop2 := loop.Synk(loop)
	assertEq(loop, loop2, t)
}

func TestCutWithdifferentNumberOfwiningPoints(t *testing.T) {
	loop1 := InitTailedLoop(3, 4, []int{0, 1, 2, 3})
	loop2 := InitTailedLoop(3, 4, []int{0, 2})
	result := loop1.Synk(loop2)
	assertEq(loop2, result, t)
}

func TestTailEnd(t *testing.T) {
	loop1 := InitTailedLoop(5, 0, []int{})
	loop2 := InitTailedLoop(7, 0, []int{})
	expected := InitTailedLoop(7, 0, []int{})
	result := loop1.Synk(loop2)
	assertEq(expected, result, t)
}

func TestDifferentLoopLength5and7(t *testing.T) {
	loop1 := InitTailedLoop(0, 5, []int{})
	loop2 := InitTailedLoop(0, 7, []int{})
	expected := InitTailedLoop(0, 5*7, []int{}) // men detta bör förstöra
	res := loop1.Synk(loop2)
	assertEq(expected, res, t)
}

func TestDifferentLoopLength2and4(t *testing.T) {
	loop1 := InitTailedLoop(0, 2, []int{})
	loop2 := InitTailedLoop(0, 4, []int{})
	expected := InitTailedLoop(0, 4, []int{}) // men detta bör förstöra
	res := loop1.Synk(loop2)
	assertEq(expected, res, t)
}

func TestTailedLoopAndWin(t *testing.T) { //lite för tidigt för denna
	t.Skip("not written yett")
	// bör ha ett exempel med match och en utan match
	// vi har ingen identitet loop blir det ett problem ?
	// vad är lägsta talen att prova 2 , 3 , 5 så 2*5 och 3*5
	// Vi f¨r ta och rita lite
}
