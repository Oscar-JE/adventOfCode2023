package card

import (
	"testing"
)

func TestValue(t *testing.T) {
	card := Card{"2"}
	if card.Value() != 0 {
		t.Errorf("Examin the value representation of cards")
	}
}
