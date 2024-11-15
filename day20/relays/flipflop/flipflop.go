package flipflop

import "day20/relays"

func InitFlipFlop(downStream []*relays.Relay) relays.Relay {
	return relays.Init(downStream, &flipflop{false})
}

type flipflop struct {
	isOn bool
}

func (f flipflop) IsOutputHeigh(inputs []relays.Signal) bool {
	return f.isOn
}

func (f *flipflop) ShouldSendSignal(inputs []relays.Signal) bool {
	containsLow := containsLow(inputs)
	if containsLow {
		f.isOn = !f.isOn
	}
	return containsLow
}

func containsLow(inputs []relays.Signal) bool {
	res := false
	for _, s := range inputs {
		if !s.GetHeigh() {
			res = true
			break
		}
	}
	return res
}
