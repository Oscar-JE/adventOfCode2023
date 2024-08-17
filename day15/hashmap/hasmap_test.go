package hashmap

import "testing"

func TestShortExample(t *testing.T) {
	instructions := []string{"rn=1", "cm-", "qp=3", "cm=2", "qp-", "pc=4", "ot=9", "ab=5", "pc-", "pc=6", "ot=7"}
	hashMap := Init()
	res := hashMap.partOne(instructions)
	expected := 145
	if expected != res {
		t.Errorf("short example in spec not fullfilled")
	}
}
