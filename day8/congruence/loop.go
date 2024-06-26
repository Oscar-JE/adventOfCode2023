package congruence

import (
	"day8/congruence/math"
	"fmt"
)

type TailedLoopMultipleEndPoints struct {
	tail         int
	loopLength   int
	winingPoints []int // meh hur gör om de finns på svansen? Kanske inte krävs av uppgiften eller så är 0 från början av svansen. Ja så gör vi
}

func InitTailedLoop(tail int, loopLength int, winningIndexes []int) TailedLoopMultipleEndPoints {
	return TailedLoopMultipleEndPoints{tail: tail, loopLength: loopLength, winingPoints: winningIndexes}
}

func (loop TailedLoopMultipleEndPoints) Eq(other TailedLoopMultipleEndPoints) bool {
	taledEq := loop.tail == other.tail
	loopLenEq := loop.loopLength == other.loopLength
	if !taledEq || !loopLenEq {
		return false
	}
	winingPointsEq := true
	if len(loop.winingPoints) != len(other.winingPoints) {
		return false
	}
	for i := range loop.winingPoints {
		winingPointsEq = winingPointsEq && loop.winingPoints[i] == other.winingPoints[i]
	}
	return winingPointsEq
}

func (t TailedLoopMultipleEndPoints) String() string {
	return fmt.Sprintf("tail : %d , loopLength: %d , winingPoints %v ", t.tail, t.loopLength, t.winingPoints)
}

func (t TailedLoopMultipleEndPoints) Synk(other TailedLoopMultipleEndPoints) TailedLoopMultipleEndPoints {
	if t.Eq(other) {
		return other
	}
	winningIndexes := math.SetIntersect(t.winingPoints, other.winingPoints)
	tail := math.Max(t.tail, other.tail)
	loopLength := t.loopLength * other.loopLength / math.GCD(t.loopLength, other.loopLength) // antagligen något annat vid inte primtal
	return InitTailedLoop(tail, loopLength, winningIndexes)
}

func (l TailedLoopMultipleEndPoints) GetFirstWiningIndex() int { // bättre att skriva med error men lämnar såhär när vi testar
	if len(l.winingPoints) != 0 {
		return l.winingPoints[0]
	}
	return -1
}
