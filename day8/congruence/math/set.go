package math

func SetIntersect(numbers1 []int, numbers2 []int) []int {
	shortest := numbers1
	longest := numbers2
	result := []int{}
	if len(numbers2) < len(numbers1) {
		shortest = numbers2
		longest = numbers1
	}
	for i := range shortest {
		if Contains(shortest[i], longest) {
			result = append(result, shortest[i])
		}
	}
	return result
}

func Contains(element int, numbers []int) bool {
	for i := range numbers {
		if numbers[i] == element {
			return true
		}
	}
	return false
}
