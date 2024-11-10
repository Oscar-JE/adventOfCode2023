package main

import (
	"day19/item"
	"day19/parse"
	"day19/workflow"
	"fmt"
	"os"
	"strings"
)

func main() {
	part2()
}

func part2() {
	fileName := "small.txt"
	workflow, _ := Parse(fileName)
	fmt.Println(workflow.FindNumberOfPossibleAccept())
}

func part1() {
	//fileName := "small.txt"
	fileName := "big.txt"
	workFlow, items := Parse(fileName)
	fmt.Println(workFlow)
	fmt.Println(items)
	sum := workFlow.ProcessItems(items)
	fmt.Println(sum)
}

func Parse(fileName string) (workflow.Workflow, []item.Item) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic("file reading failed")
	}
	workflowsAndItems := strings.Split(string(content), "\r\n\r\n")
	fmt.Println("first")
	fmt.Println(workflowsAndItems[0])
	fmt.Println("second")
	fmt.Println(workflowsAndItems[1])
	nodesInfos := parse.ParseNodes(workflowsAndItems[0])
	workFlow := workflow.Init(nodesInfos)
	items := parse.ParseItems(workflowsAndItems[1])
	return workFlow, items
}
