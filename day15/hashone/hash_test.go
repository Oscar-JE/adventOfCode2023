package hashone

import "testing"

func TestHash(t *testing.T) {
	hashTester("rn=1", 30, t)
	hashTester("cm-", 253, t)
	hashTester("qp=3", 97, t)
	hashTester("cm=2", 47, t)
	hashTester("qp-", 14, t)
	hashTester("pc=4", 180, t)
	hashTester("ot=9", 9, t)
	hashTester("ab=5", 197, t)
	hashTester("pc-", 48, t)
	hashTester("pc=6", 214, t)
	hashTester("ot=7", 231, t)
	hashTester("rn", 0, t)
}

func hashTester(in string, exp int, t *testing.T) {
	res := Hash(in)
	if res != exp {
		t.Errorf(in+" should result in %d but was %d", exp, res)
	}
}
