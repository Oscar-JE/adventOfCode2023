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
		i++
	}
	return TicketTally{multiples: multiples, tickets: tickets}
}

func (tally *TicketTally) Multiply() {
	for ticketIndex := range tally.tickets {
		ticket := tally.tickets[ticketIndex]
		nrOfmatches := NrOfMatches(ticket.winning, ticket.holding)
		nrTicket := tally.multiples[ticketIndex]
		for index := ticketIndex + 1; index < min(ticketIndex+1+nrOfmatches, len(tally.multiples)); index++ {
			tally.multiples[index] += nrTicket
		}
	}
}

func (tally TicketTally) SumOfTally() int {
	sum := 0
	for _, nr := range tally.multiples {
		sum += nr
	}
	return sum
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
