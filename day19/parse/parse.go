package parse

import (
	"day19/item"
	"day19/workflow"
	"strconv"
	"strings"
)

func ParseNodes(nodeRep string) []workflow.NodeInfo {
	// fortsätt här
	
	return []workflow.NodeInfo{}
}

func ParseItems(strItems string) []item.Item {
	lines := strings.Split(strItems, "\n\r")
	items := []item.Item{}
	for _, line := range lines {
		items = append(items, parseItem(line))
	}
	return items
}

func parseItem(strItem string) item.Item {
	line := strings.Replace(strItem, "{", "", -1)
	line = strings.Replace(line, "}", "", -1)
	line = strings.Replace(line, "=", "", -1)
	line = strings.Replace(line, "x", "", -1)
	line = strings.Replace(line, "m", "", -1)
	line = strings.Replace(line, "a", "", -1)
	line = strings.Replace(line, "s", "", -1)
	intReps := strings.Split(line, ",")
	ints := []int{}
	for _, intRep := range intReps {
		val, err := strconv.Atoi(intRep)
		if err != nil {
			panic("integer parsing failed")
		}
		ints = append(ints, val)
	}
	return item.Init(ints[0], ints[1], ints[2], ints[3])
}
