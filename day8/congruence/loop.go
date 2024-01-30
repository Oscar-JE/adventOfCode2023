package congruence

/*
	I den här filen undersöker congruence problem och hur man löser dem.
	Vi börjar i svansen och rör oss sedan in i loopen
*/

type TailedLoopMultipleEndPoints struct {
	tail         int
	loopLength   int
	winingPoints []int
}

func InitTailedLoop(tail int, loopLength int, winningIndexes []int) TailedLoopMultipleEndPoints {
	if someNegative(append([]int{tail, loopLength}, winningIndexes...)) {
		panic("negative integers not allowed in congruence problem")
	}
	return TailedLoopMultipleEndPoints{tail: tail, loopLength: loopLength, winingPoints: winningIndexes}
}

func someNegative(numbers []int) bool {
	negative := false
	for _, nr := range numbers {
		negative = negative || nr < 0
	}
	return negative
}

type TailedLoop struct {
	tail         int
	loopLength   int
	winningPoint int
}

func cut(loop1 TailedLoop, loop2 TailedLoop) (TailedLoop, error) {
	return loop2, nil
}
