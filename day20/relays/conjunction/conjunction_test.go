package conjunction

import (
	"day20/relays"
	"testing"
)

func TestShouldSendSignal(t *testing.T) {
	con := InitConjunction([]*relays.Relay{})
	if !con.ShouldSendSignal([]relays.Signal{relays.InitSignal(false, nil)}) {
		t.Errorf("should always send signal when receiving input signal")
	}
	if con.ShouldSendSignal([]relays.Signal{}) {
		t.Errorf("no input signal should not send a signal")
	}
}

func TestIsSignalHeigh(t *testing.T) {
	conDupe := InitConjunction([]*relays.Relay{})
	source1 := relays.Init([]*relays.Relay{}, &conDupe)
	source2 := relays.Init([]*relays.Relay{}, &conDupe)
	con := InitConjunction([]*relays.Relay{&source1, &source2})
	if !con.IsOutputHigh([]relays.Signal{}) {
		t.Errorf("all previous low should return in heigh")
	}
	if !con.IsOutputHigh([]relays.Signal{relays.InitSignal(true, &source1)}) {
		t.Errorf("Only one previous heigh should result in Heigh")
	}
	if con.IsOutputHigh([]relays.Signal{relays.InitSignal(true, &source2)}) {
		t.Errorf("All previous heigh should result in low output")
	}
	if !con.IsOutputHigh([]relays.Signal{relays.InitSignal(false, &source1)}) {
		t.Errorf("mix leads to heigh")
	}

}
