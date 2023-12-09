package maps

import (
	"testing"
)

func TestDoMapping(t *testing.T) {
	in := 79
	firstMap := Mapping{fromStart: 98, toStart: 50, length: 2}
	res := firstMap.Do(in)
	if res != 79 {
		t.Errorf("Wrong answer at first mapping")
	}
	secondMap := Mapping{fromStart: 50, toStart: 52, length: 48}
	res2 := secondMap.Do(79)
	if res2 != 81 {
		t.Errorf("wrong answer at second mapping")
	}

}
