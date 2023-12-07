package solution

import (
	"testing"
)

func TestNrOfWins(t *testing.T) {
	race := Race{Time: 7, Distance: 9}
	res := NrOfWins(race)
	if res != 4 {
		t.Errorf("wrong number of wins")
	}
}
