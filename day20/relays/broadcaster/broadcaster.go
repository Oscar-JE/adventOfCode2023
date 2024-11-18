package broadcaster

import "day20/relays"

func InitBroadcaster() relays.Relay {
	return relays.Init([]*relays.Relay{}, sameAsFirst{})
}

type sameAsFirst struct {
}

func (a sameAsFirst) IsOutputHigh(signals []relays.Signal) bool {
	if len(signals) == 0 {
		return false
	}
	return signals[0].GetHigh()
}

func (a sameAsFirst) ShouldSendSignal(signals []relays.Signal) bool {
	return len(signals) > 0
}
