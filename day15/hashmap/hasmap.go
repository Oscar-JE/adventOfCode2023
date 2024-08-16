package hashmap

import (
	"day15/hashmap/box"
	"strings"
)

type Map struct {
	boxes []box.Box
}

func Init() Map {
	boxes := []box.Box{}
	for i := 0; i < 256; i++ {
		boxes = append(boxes, box.Init([]int{}))
	}
	return Map{boxes: boxes}
}

func (m *Map) PerformInstruction(inst string){
	isRemove := strings.Contains(inst, "-")
	if 
}