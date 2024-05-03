package field

import "day12/springs"

type Field struct {
	layout []springs.Condition
}

func Init(runes []rune) Field {
	layout := []springs.Condition{}
	for _, r := range runes {
		layout = append(layout, springs.Init(r))
	}
	return Field{layout: layout}
}
