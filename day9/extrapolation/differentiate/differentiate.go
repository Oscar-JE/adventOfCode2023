package differentiate

func Diff(series []int) []int {
	diffed := []int{}
	for i := 1; i < len(series); i++ {
		diffed = append(diffed, series[i]-series[i-1])
	}
	return diffed
}
