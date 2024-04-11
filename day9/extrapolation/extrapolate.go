package extrapolation

import "day9/extrapolation/differentiate"

func Extrapolate(series []int) int {
	if isAllZeros(series) {
		return 0
	}
	dif := differentiate.Diff(series)
	return series[len(series)-1] + Extrapolate(dif)
}

func isAllZeros(series []int) bool {
	res := true
	for _, el := range series {
		res = res && el == 0
	}
	return res
}
