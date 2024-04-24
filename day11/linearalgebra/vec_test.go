package linearalgebra

import "testing"

func TestManhattan(t *testing.T) {
	v1 := InitVec(6, 1)
	v2 := InitVec(11, 5)
	if ManhattanDistans(v1, v2) != 9 {
		t.Errorf("Examin wheter the assignement calls for manhattandistance")
	}
}
