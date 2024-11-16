package conjunction

import (
	"day20/relays"
	"testing"
)

func TestIngenSignalIn(t *testing.T) {
	con := InitConjunction([]*relays.Relay{})
	if !con.ShouldSendSignal([]relays.Signal{relays.InitSignal(false, nil)}) { // skulle kunna vara en lista av bool istället
		t.Errorf("should always send signal when receiving input signal")
	}
}
