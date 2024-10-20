package parse

import (
	"day19/item"
	"day19/workflow"
	"strconv"
	"strings"
)

func ParseNodes(nodesRep string) []workflow.NodeInfo {
	lines := strings.Split(nodesRep, "\n\r")
	nodeInfos := []workflow.NodeInfo{}
	for _, line := range lines {
		nodeInfos = append(nodeInfos, parseNode(line))
	}
	return nodeInfos
}

func parseNode(nodeRep string) workflow.NodeInfo {
	nameAndRest := strings.Split(nodeRep, "{")
	name := nameAndRest[0]
	rulesRep := nameAndRest[1]
	rulesRep = strings.Replace(rulesRep, "}", "", -1)
	rules := parseRules(rulesRep)
	return workflow.NodeInfo{Id: name, Rules: rules}
}

func parseRules(rulesRep string) []workflow.RuleRep {
	ruleReps := strings.Split(rulesRep, ",")
	rules := []workflow.RuleRep{}
	for _, ruleRep := range ruleReps {
		rules = append(rules, parseRule(ruleRep))
	}
	return rules
}

func parseRule(ruleRep string) workflow.RuleRep {
	//a<2006:qkq
	demandAndDestination := strings.Split(ruleRep, ":")
	if len(demandAndDestination) == 1 {
		return workflow.RuleRep{Demand: item.NoDemand{}, DestinationId: demandAndDestination[0]}
	}
	demandRep := demandAndDestination[0]
	destinationId := demandAndDestination[1]
	return workflow.RuleRep{Demand: parseDemand(demandRep), DestinationId: destinationId}
}

func parseDemand(demandRep string) item.Demand {
	index := strings.Index(demandRep, ">") + strings.Index(demandRep, "<") + 1
	numRep := demandRep[index+1:]
	number, err := strconv.Atoi(numRep)
	if err != nil {
		panic("numberconversions failed")
	}
	if strings.Contains(demandRep, "x") {
		if strings.Contains(demandRep, ">") {
			return item.Cooler{CoolDemand: number}
		} else {
			return item.LessCool{CoolDemand: number}
		}
	} else if strings.Contains(demandRep, "m") {
		if strings.Contains(demandRep, ">") {
			return item.MoreMusical{Musical: number}
		} else {
			return item.LessMusical{Musical: number}
		}
	} else if strings.Contains(demandRep, "a") {
		if strings.Contains(demandRep, ">") {
			return item.MoreAerodynamic{Aerodynamic: number}
		} else {
			return item.LessAerodynamic{Aerodynamic: number}
		}
	} else if strings.Contains(demandRep, "s") {
		if strings.Contains(demandRep, ">") {
			return item.MoreShiny{Shiny: number}
		} else {
			return item.LessShiny{Shiny: number}
		}
	} else {
		panic("Unknown characteristics")
	}
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
