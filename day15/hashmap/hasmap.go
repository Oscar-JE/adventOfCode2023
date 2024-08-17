package hashmap

import (
	"day15/hashmap/box"
	"day15/hashmap/lens"
	"day15/hashone"
	"fmt"
	"strconv"
	"strings"
)

type Map struct {
	boxes []box.Box
}

func Init() Map {
	boxes := []box.Box{}
	for i := 0; i < 256; i++ {
		boxes = append(boxes, box.Init([]lens.Lens{}))
	}
	return Map{boxes: boxes}
}

func (m *Map) PerformInstruction(inst string) {
	isRemove := strings.Contains(inst, "-")
	isAdd := strings.Contains(inst, "=")
	if (!isRemove) && (!isAdd) {
		panic("unknown operator")
	}
	if isRemove {
		label := strings.Replace(inst, "-", "", 1)
		hashValue := hashone.Hash(label)
		m.remove(label, hashValue)
	}
	if isAdd {
		labelAndFokalLength := strings.Split(inst, "=")
		label := labelAndFokalLength[0]
		fokalLength, err := strconv.Atoi(labelAndFokalLength[1])
		if err != nil {
			panic("fokal length extraction failed")
		}
		hashValue := hashone.Hash(label)
		lens := lens.Init(label, fokalLength)
		m.add(lens, hashValue)
	}
}

func (m *Map) remove(label string, hashValue int) {
	m.boxes[hashValue].Remove(lens.Init(label, -1))
}

func (m *Map) add(lens lens.Lens, hashValue int) {
	m.boxes[hashValue].Add(lens)
}

func (m Map) FocusingPower() int {
	sum := 0
	for index, box := range m.boxes {
		sum += (index + 1) * box.FocusingPower()
	}
	return sum
}

func (m Map) partOne(instructions []string) int {
	for _, inst := range instructions {
		m.PerformInstruction(inst)
	}
	return m.FocusingPower()
}

func (m Map) String() string {
	rep := ""
	for index, box := range m.boxes {
		if box.IsEmpty() {
			continue
		}
		rep += "box nr " + fmt.Sprint(index) + ": " + box.String() + "\n"
	}
	return rep
}
