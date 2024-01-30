package congruence

type TailedLoopMultipleEndPoints struct {
	tail         int
	loopLength   int
	winingPoints []int
}

func InitTailedLoop(tail int, loopLength int, winningIndexes []int) TailedLoopMultipleEndPoints {
	return TailedLoopMultipleEndPoints{tail: tail, loopLength: loopLength, winingPoints: winningIndexes}
}
