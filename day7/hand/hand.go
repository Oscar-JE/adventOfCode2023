package hand

import (
	"day7/card"
)

type Hand struct {
	cards [5]card.Card
}

func Init(cards string) Hand {
	if len(cards) != 5 {
		panic("wrong number of cards dealt to a hand")
	}
	content := [5]card.Card{card.Init(string(cards[0])), card.Init(string(cards[1])), card.Init(string(cards[2])), card.Init(string(cards[3])), card.Init(string(cards[4]))}
	return Hand{content}
}

func (h Hand) nrOfMatch(c card.Card) int {
	nr := 0
	for _, card := range h.cards {
		if card.Eq(c) {
			nr++
		}
	}
	return nr
}

type CardGroup struct {
	whatCard card.Card
	number   int
}

func (h Hand) CardGrouping() []CardGroup {
	cardGroups := []CardGroup{}
	for _, rep := range card.Ordering {
		cardCandidate := card.Init(string(rep))
		matches := h.nrOfMatch(cardCandidate)
		if matches > 0 {
			cardGroups = append(cardGroups, CardGroup{whatCard: cardCandidate, number: matches})
		}
	}
	cardGroups = convertJokers(cardGroups)
	return cardGroups
}

func convertJokers(cardGroups []CardGroup) []CardGroup {
	if len(cardGroups) == 1 {
		// no conversion possible
		return cardGroups
	}
	indexOfJoker := -1
	for i, group := range cardGroups { // find first not joker
		if group.whatCard.Eq(card.Init("J")) {
			indexOfJoker = i
			break
		}
	}
	if indexOfJoker == -1 { // there is no Jokers
		return cardGroups
	}
	// find largest groupp nor joker
	largestNotJoker := 0
	if indexOfJoker == largestNotJoker {
		largestNotJoker++
	}
	for i, group := range cardGroups {
		if cardGroups[largestNotJoker].number < group.number && !group.whatCard.Eq(card.Init("J")) {
			largestNotJoker = i
		}
	}
	// add jokers to largestgroup
	cardGroups[largestNotJoker].number += cardGroups[indexOfJoker].number

	// remove Jokers
	beforeJoker := cardGroups[0:indexOfJoker]
	afterJoker := []CardGroup{}
	if indexOfJoker != len(cardGroups)-1 {
		afterJoker = append(afterJoker, cardGroups[indexOfJoker+1:]...)
	}
	return append(beforeJoker, afterJoker...)
}

func (h Hand) highCard() bool {
	groups := h.CardGrouping()
	return len(groups) == 5
}

func (h Hand) onePair() bool {
	groups := h.CardGrouping()
	nrOfPairs := nrOfGroupsOf(groups, 2)
	nrOf3 := nrOfGroupsOf(groups, 3)
	return nrOfPairs == 1 && nrOf3 == 0
}

func (h Hand) twoPair() bool {
	groups := h.CardGrouping()
	nrOfPairs := nrOfGroupsOf(groups, 2)
	return nrOfPairs == 2
}

func (h Hand) threeOfAKind() bool {
	groups := h.CardGrouping()
	nrOf3 := nrOfGroupsOf(groups, 3)
	nrOf2 := nrOfGroupsOf(groups, 2)
	return nrOf3 == 1 && nrOf2 == 0
}

func (h Hand) fullHouse() bool {
	groups := h.CardGrouping()
	nrOf3 := nrOfGroupsOf(groups, 3)
	nrOf2 := nrOfGroupsOf(groups, 2)
	return nrOf3 == 1 && nrOf2 == 1
}

func (h Hand) fiveOfAKind() bool {
	groups := h.CardGrouping()
	return len(groups) == 1
}

func (h Hand) fourOfAKind() bool {
	groups := h.CardGrouping()
	nrOf4 := nrOfGroupsOf(groups, 4)
	return nrOf4 == 1
}

func nrOfGroupsOf(groups []CardGroup, nr int) int {
	nrOfPairs := 0
	for _, group := range groups {
		if group.number == nr {
			nrOfPairs++
		}
	}
	return nrOfPairs
}
