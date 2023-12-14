package maps

import (
	"fmt"
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
	first := SubTransform{fromStart: 56, toStart: 60, length: 37}
	second := SubTransform{fromStart: 93, toStart: 56, length: 4}
	trans := Transform{[]SubTransform{first, second}}
	fmt.Println(trans)
	in := []int{78, 43, 82, 35}
	out := []int{82, 43, 86, 35}
	for i := range in {
		if trans.Do(in[i]) != out[i] {
			t.Errorf("wrong in test humidity to location")
		}
	}
}
