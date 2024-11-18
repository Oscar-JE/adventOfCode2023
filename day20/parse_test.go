package main

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	net := parseFile("example1.txt")
	fmt.Println(net)
}
