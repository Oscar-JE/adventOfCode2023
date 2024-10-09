package hex

import (
	"testing"
)

func TestHexFull(t *testing.T) {
	res := Parse("0dc57")
	expected := 56407
	if res != expected {
		t.Errorf("parse of 0dc57 should be 56407")
	}
}

func TestHexSingular(t *testing.T) {
	sample := "f"
	expected := 15
	res := Parse(sample)
	if res != expected {
		t.Errorf("paring of f should be 15")
	}
}
