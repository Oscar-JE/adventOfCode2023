package flipflop

import (
	"day20/relays"
	"testing"
)

func TestIngenSignalOmIngenLågSignal(t *testing.T) {
	flipFlopAv := flipflop{false}
	flipFlopPå := flipflop{true}
	if flipFlopAv.ShouldSendSignal([]relays.Signal{relays.InitSignal(true, nil)}) || flipFlopPå.ShouldSendSignal([]relays.Signal{relays.InitSignal(true, nil)}) {
		t.Errorf("heigh signal should not give any output ")
	}
}

func TestFlipOmLågSignal(t *testing.T) {
	flipFlopAv := flipflop{false}
	flipFlopPå := flipflop{true}
	singalPaket := []relays.Signal{relays.InitSignal(false, nil)}
	if !flipFlopAv.ShouldSendSignal(singalPaket) || !flipFlopPå.ShouldSendSignal(singalPaket) {
		t.Errorf("låg signal ska alltid ge en signal")
	}
	if !flipFlopAv.isOn || flipFlopPå.isOn {
		t.Errorf("Tillstånden shiftat på oväntat sätt")
	}

}
