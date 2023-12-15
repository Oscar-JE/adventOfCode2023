package hand

import (
	"day7/card"
	"testing"
)

func TestNumberOfMatch(t *testing.T) {
	hand := Init("22333")
	if hand.nrOfMatch(card.Init("2")) != 2 && hand.nrOfMatch(card.Init("3")) != 3 {
		t.Errorf("rethink how you sort your hand")
	}
}

func TestCardGrouping(t *testing.T) {
	hand := Init("22333")
	groups := hand.CardGrouping()
	if !groups[0].whatCard.Eq(card.Init("2")) && groups[0].number != 2 &&
		!groups[1].whatCard.Eq(card.Init("3")) && groups[1].number != 3 {
		t.Errorf("grouping of cards done wrong")
	}
}

func TestCatOrder1(t *testing.T) {
	h1 := Init("KTJJT")
	if h1.CategoryScore() != 2 {
		t.Errorf("wrong category for T55J5")
	}
}

func TestCatOrder2(t *testing.T) {
	h1 := Init("KK677")
	if h1.CategoryScore() != 2 {
		t.Errorf("wrong category for T55J5")
	}
}

func TestValueScore(t *testing.T) {
	h1 := Init("KTJJT")
	h2 := Init("KK677")
	v1 := h1.ValueScore()
	v2 := h2.ValueScore()
	if !(v1 < v2) {
		t.Errorf("wrong values")
	}
}

func TestOrdering(t *testing.T) {
	h1 := Init("KTJJT")
	h2 := Init("KK677")
	low := Lower(h1, h2)
	if !low {
		t.Errorf("faulse ordering")
	}
}

func TestHandCatogorising(t *testing.T)
