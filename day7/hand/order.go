package hand

import "day7/card"

func Lower(hand Hand, other Hand) bool {
	catScoreLow := hand.CategoryScore()
	catScoreHigh := other.CategoryScore()
	if catScoreLow < catScoreHigh {
		return true
	}
	valueScoreLow := hand.ValueScore()
	valueScoreHigh := other.ValueScore()
	if hand.CategoryScore() == other.CategoryScore() && valueScoreLow < valueScoreHigh {
		return true
	}
	return false
}

func (h Hand) ValueScore() int {
	value := 0
	factor := 1
	for i := len(h.cards) - 1; i >= 0; i-- {
		value += h.cards[i].Value() * factor
		factor *= factor * len(card.Ordering)
	}
	return value
}

func (h Hand) CategoryScore() int {
	retValue := -1
	numberActivations := 0
	activations := ""
	if h.highCard() {
		retValue += 1
		numberActivations++
		activations += "highCard "
	}
	if h.onePair() {
		retValue += 2
		numberActivations++
		activations += "onePair "
	}
	if h.twoPair() {
		retValue += 3
		numberActivations++
		activations += "twoPair "
	}
	if h.threeOfAKind() {
		retValue += 4
		numberActivations++
		activations += "threeOfAKind "
	}
	if h.fullHouse() {
		retValue += 5
		numberActivations++
		activations += "fullHouse "
	}
	if h.fourOfAKind() {
		retValue += 6
		numberActivations++
		activations += "fourOfAKind  "
	}
	if h.fiveOfAKind() {
		retValue += 7
		numberActivations++
		activations += "fiveOfAKind "
	}
	if retValue == -1 {
		panic("uncategoirised hand was found")
	}
	if numberActivations != 1 {
		panic("wrong number of activations for hand :" + h.cards[0].Rep + h.cards[1].Rep + h.cards[2].Rep + h.cards[3].Rep + h.cards[4].Rep + "  activations: " + activations)
	}
	return retValue
}
