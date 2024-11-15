package broadcaster

import "day20/relays"

func InitBroadcaster() relays.Relay {
	return relays.Init([]*relays.Relay{}, alwaysLow{})
}

type alwaysLow struct {
}

func (a alwaysLow) IsOutputHeigh(signals []relays.Signal) bool {
	return false
}
