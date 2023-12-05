package ticket

import "testing"

func TestMatched(t *testing.T) {
	winning := []int{1, 2, 3, 4}
	holding := []int{4, 0, 90, 2}
	ticket := Ticket{winning: winning, holding: holding}
	matched := ticket.matched()

	if matched[0] != 2 && matched[1] != 4 {
		t.Errorf("Error in matched")
	}
}
