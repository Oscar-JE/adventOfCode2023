package hashone

import "day15/ascii"

func hash(runes []rune) int {
	hash := 0
	for _, r := range runes {
		acii := ascii.Value(r)
		hash += acii
		hash *= 17
		hash = hash % 256
	}
	return hash
}

func Hash(rep string) int {
	return hash([]rune(rep))
}
