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

func TestTailedLoopAndWin(t *testing.T) {
	t.Skip("not written yett")
	// bör ha ett exempel med match och en utan match
	// vi har ingen identitet loop blir det ett problem ?
}
