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

type Relay struct {
	signalTmp []Signal
	outgoing  []*Relay
	com       communicator
}

func Init(outgoing []*Relay, com communicator) Relay {
	return Relay{outgoing: outgoing, com: com}
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
