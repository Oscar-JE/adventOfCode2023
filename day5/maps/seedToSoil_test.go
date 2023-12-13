package maps

import (
	"testing"
)

func TestDoMapping(t *testing.T) {
	in := 79
	firstMap := SubTransform{fromStart: 98, toStart: 50, length: 2}
	res, _ := firstMap.Do(in)
	if res != 79 {
		t.Errorf("Wrong answer at first mapping")
	}
	secondMap := SubTransform{fromStart: 50, toStart: 52, length: 48}
	res2, _ := secondMap.Do(79)
	if res2 != 81 {
		t.Errorf("wrong answer at second mapping")
	}
}

func TestTransform(t *testing.T) {
	in := 78
}
