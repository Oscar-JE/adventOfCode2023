package main

import (
	"bufio"
	"day8/graph"
	"day8/traverser"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("file not found")
	}
	scan := bufio.NewScanner(file)
	scan.Scan()
	policyLine := scan.Text()
	scan.Scan() // we know whis is a blank row
	graphInfo := []graph.NameLeftRight{}
	for scan.Scan() {
		nodeLine := scan.Text()
		splited := strings.Split(nodeLine, "=")
		name := strings.Trim(splited[0], " ")
		cleaned := strings.Trim(splited[1], "() ")
		leftAndRight := strings.Split(cleaned, ", ")
		graphInfo = append(graphInfo, graph.NameLeftRight{Name: name, Left: leftAndRight[0], Right: leftAndRight[1]})
	}
	fmt.Println(traverser.NrOfMutualSteps(policyLine, graphInfo))
}
