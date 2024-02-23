package congruence

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
		return false;
	}
}