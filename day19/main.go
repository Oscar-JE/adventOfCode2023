package main

import (
	"day19/parse"
	"day19/workflow"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
}

func part1() {
	fileName := "small.txt"
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic("file reading failed")
	}
	// divide on double  blank line
	workflowsAndItems := strings.Split(string(content), "\r\n\r\n")
	fmt.Println("first")
	fmt.Println(workflowsAndItems[0])
	fmt.Println("second")
	fmt.Println(workflowsAndItems[1])
	nodesInfos := parse.ParseNodes(workflowsAndItems[0])
	workFlow := workflow.Init(nodesInfos)
	items := parse.ParseItems(workflowsAndItems[1])
	fmt.Println(workFlow)
	fmt.Println(items)
	sum := workFlow.ProcessItems(items)
	fmt.Println(sum)
}
