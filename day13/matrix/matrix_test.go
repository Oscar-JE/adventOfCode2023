package matrix

import (
	"testing"
)

func TestConstruction(t *testing.T) {
	m := Init()
	println(m.Get(0, 0))
}

func TestSetGet(t *testing.T) {
	m := InitZero(2, 3)
	m.Set(1, 2, 5)
	if m.Get(1, 2) != 5 {
		t.Errorf("Examin set and get")
	}
}

func TestTranspose(t *testing.T) {
	m := InitZero(2, 3)
	m.Set(1, 2, 5)
	m.Transpose()
	if m.Get(2, 1) != 5 {
		t.Errorf("Examin Transpose")
	}
}

func TestTransponseQube(t *testing.T) {
	m := InitZero(2, 2)
	m.Set(0, 0, 1)
	m.Set(0, 1, 2)
	m.Set(1, 0, 3)
	m.Set(1, 1, 4)
	m.Transpose()
	q1 := m.Get(0, 0) != 1
	q2 := m.Get(0, 1) != 3
	q3 := m.Get(1, 0) != 2
	q4 := m.Get(1, 1) != 4
	if q1 || q2 || q3 || q4 {
		t.Errorf("transpose failed in cube case")
	}
}
