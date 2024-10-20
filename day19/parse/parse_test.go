package parse

import (
	"day19/item"
	"testing"
)

func TestParseItem(t *testing.T) {
	strRep := "{x=1,m=2,a=3,s=4}"
	expected := item.Init(1, 2, 3, 4)
	res := parseItem(strRep)
	if res != expected {
		t.Errorf("parse item does not work")
	}
}
