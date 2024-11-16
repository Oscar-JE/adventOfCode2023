package conjunction

import "day20/relays"

func InitConjuncttion(inputs []*relays.Relay, outgoing []*relays.Relay) relays.Relay {
	con := InitConjunction(inputs)
	return relays.Init(outgoing, &con)
}

type conjunction struct {
	previousSignals []relays.Signal
}

func InitConjunction(inputs []*relays.Relay) conjunction {
	previousSignals := []relays.Signal{}
	for _, relay := range inputs {
		previousSignals = append(previousSignals, relays.InitSignal(false, relay))
	}
	return conjunction{previousSignals: previousSignals}
}

func (c conjunction) ShouldSendSignal(inputs []relays.Signal) bool {
	return len(inputs) > 0
}

func (c *conjunction) IsOutputHeigh(inputs []relays.Signal) bool {
	heigh := c.isAllPreviousHeigh()
	c.setPreviousSignals(inputs)
	return heigh
}

func (c conjunction) isAllPreviousHeigh() bool {
	for _, sig := range c.previousSignals {
		if !sig.GetHeigh() {
			return false
		}
	}
	return true
}

func (c *conjunction) setPreviousSignals(inputs []relays.Signal) {
	c.previousSignals = inputs
}
