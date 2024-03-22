package math

import (
	"testing"
)

func TestGCDPrime(t *testing.T) {
	res := GCD(3, 5)
	if res != 1 {
		t.Errorf("coprimes should hava a GCD of one")
	}
}

func TestGCD(t *testing.T) {
	res := GCD(5*3, 11*3)
	if res != 3 {
		t.Errorf("expected the common faktor")
	}
}
