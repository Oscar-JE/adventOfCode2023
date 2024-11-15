package relays

type Signal struct {
	heigh  bool
	source *Relay
}

type Relay struct {
	outgoing []*Relay
	Com      communicator
}

func Init(outgoing []*Relay, com communicator) Relay {
	return Relay{outgoing: outgoing, Com: com}
}

type communicator interface {
	IsOutputHeigh(inputs []Signal) bool
}
