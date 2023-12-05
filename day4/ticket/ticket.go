package ticket

type Ticket struct {
	winning []int
	holding []int
}

type TicketTally struct {
	tickets   []Ticket
	multiples []int
}

func InitTally(tickets []Ticket) TicketTally {
	multiples := []int{}
	i := 0
	for i < len(tickets) {
		multiples = append(multiples, 1)
	}
	return TicketTally{multiples: multiples, tickets: tickets}
}

func (tally *TicketTally) multiply() {
}

func (tally TicketTally) copyedIndexes(ticket Ticket) []int {
	copiedIndexes := []int{}
	matched := ticket.matched()
	for i := range matched {
		matched[i] = matched[i] - 1
	}
	for i := range matched {
		if 0 <= matched[i] && matched[i] < len(tally.tickets) {
			copiedIndexes = append(copiedIndexes, matched[i])
		}
	}
	return copiedIndexes
}

func (t Ticket) matched() []int {
	matched := []int{}
	for i := range t.winning {
		for j := range t.holding {
			if t.winning[i] == t.holding[j] {
				matched = append(matched, t.winning[i])
			}
		}
	}
	return matched
}
