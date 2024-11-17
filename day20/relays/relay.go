package relays

type Signal struct {
	heigh  bool
	source *Relay
}

func InitSignal(heigh bool, source *Relay) Signal {
	return Signal{heigh: heigh, source: source}
}

func (s Signal) GetHeigh() bool {
	return s.heigh
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
	isOutputHeigh := r.com.IsOutputHeigh(r.signalTmp)
	r.signalTmp = []Signal{}
	for _, downStream := range r.outgoing {
		downStream.signalTmp = append(downStream.signalTmp, Signal{isOutputHeigh, r})
	}
}

type communicator interface {
	IsOutputHeigh(inputs []Signal) bool
	ShouldSendSignal(inputs []Signal) bool
}

func (r *Relay) JumpStart() {
	isOutputHeigh := r.com.IsOutputHeigh(r.signalTmp)
	for _, downStream := range r.outgoing {
		downStream.signalTmp = append(downStream.signalTmp, Signal{isOutputHeigh, r})
	}
}

//function för att sätta kommunikator beroende på sträng kan vi ta här nere
