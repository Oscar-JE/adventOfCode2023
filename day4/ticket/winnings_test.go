package ticket

import "testing"

func TestMatches(t *testing.T) {
	winning := []int{1, 2, 3, 4}
	holding := []int{4, 0, 90, 2}
	matched := NrOfMatches(winning, holding)
	expected := 2
	if matched != expected {
		t.Errorf("matches: %d , expected : %d ", matched, expected)
	}
}

func TestWinnings(t *testing.T) {
	testInt(Winnings(0), 0, t)
	testInt(Winnings(1), 1, t)
	testInt(Winnings(2), 2, t)
	testInt(Winnings(3), 4, t)
}

func testInt(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Errorf("actual : %d , expected : %d ", actual, expected)
	}
}
