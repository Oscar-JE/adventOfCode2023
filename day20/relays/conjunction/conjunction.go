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
	c.updatePreviousSignals(inputs)
	heigh := !c.isAllPreviousHeigh()
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

func (c *conjunction) updatePreviousSignals(inputs []relays.Signal) {
	for _, inSignal := range inputs {
		c.updateSingular(inSignal)
	}
}

func (c *conjunction) updateSingular(inSig relays.Signal) {
	for i, prevSig := range c.previousSignals {
		if prevSig.GetSrc() == inSig.GetSrc() {
			c.previousSignals[i] = inSig
			return
		}
	}
}
