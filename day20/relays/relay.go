package relays

type Signal struct {
	high   bool
	source *Relay
}

func InitSignal(high bool, source *Relay) Signal {
	return Signal{high: high, source: source}
}

func (s Signal) GetHigh() bool {
	return s.high
}

func (s Signal) GetSrc() *Relay {
	return s.source
}

type Relay struct {
	signalTmp []Signal
	outgoing  []*Relay
	ingoing   []*Relay
	com       communicator
}

func Init(outgoing []*Relay, com communicator) Relay {
	return Relay{outgoing: outgoing, com: com}
}

func (r *Relay) SetOutgoing(relays []*Relay) {
	r.outgoing = relays
}

func (r *Relay) AppendOutgoing(rel *Relay) {
	r.outgoing = append(r.outgoing, rel)
}

func (r *Relay) SetIngoing(relays []*Relay) {
	r.ingoing = relays
}

func (r *Relay) AppendIngoing(rel *Relay) {
	r.ingoing = append(r.ingoing, rel)
}

func (r *Relay) Propagate() {
	isOutputHigh := r.com.IsOutputHigh(r.signalTmp)
	r.signalTmp = []Signal{}
	for _, downStream := range r.outgoing {
		downStream.signalTmp = append(downStream.signalTmp, Signal{isOutputHigh, r})
	}
}

type communicator interface {
	IsOutputHigh(inputs []Signal) bool
	ShouldSendSignal(inputs []Signal) bool
}

func (r *Relay) JumpStart() {
	isOutputHigh := r.com.IsOutputHigh(r.signalTmp)
	for _, downStream := range r.outgoing {
		downStream.signalTmp = append(downStream.signalTmp, Signal{isOutputHigh, r})
	}
}

//function för att sätta kommunikator beroende på sträng kan vi ta här nere
