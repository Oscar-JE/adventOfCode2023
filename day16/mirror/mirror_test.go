package mirror

import (
	"day16/vec"
	"testing"
)

func TestSpegling(t *testing.T) {
	normal := vec.Init(1, 0)
	inkommande := vec.Init(-2, 1)
	expected := vec.Init(2, 1)
	result := spegla(normal, inkommande)
	if expected != result {
		t.Errorf("spegling failade")
	}
}

func TestInitMirrorFromBelow(t *testing.T) {
	ref := Init('/')
	input := vec.Init(-1, 0)
	expected := vec.Init(0, 1)
	res := ref.Reflect(input)[0]
	if expected != res {
		t.Errorf("reflection from bellow failed")
	}
}

func TestInitMirrorFromLeft(t *testing.T) {
	ref := Init('/')
	input := vec.Init(0, 1)
	expected := vec.Init(-1, 0)
	res := ref.Reflect(input)[0]
	if expected != res {
		t.Errorf("reflection from left failed")
	}
}

func TestInitMirrorFromAbove(t *testing.T) {
	ref := Init('/')
	input := vec.Init(1, 0)
	expected := vec.Init(0, -1)
	res := ref.Reflect(input)[0]
	if expected != res {
		t.Errorf("reflection from above failed")
	}
}

func TestInitMirrorFromRight(t *testing.T) {
	ref := Init('/')
	input := vec.Init(0, -1)
	expected := vec.Init(1, 0)
	res := ref.Reflect(input)[0]
	if expected != res {
		t.Errorf("reflection from right failed")
	}
}

func TestAntiMirrorFromAbove(t *testing.T) {
	ref := Init('\\')
	input := vec.Init(1, 0)
	expected := vec.Init(0, 1)
	res := ref.Reflect(input)[0]
	if expected != res {
		t.Errorf("reflection from above failed")
	}
}

func TestSplitterFromAbove(t *testing.T) {
	ref := Init('|')
	input := vec.Init(1, 0)
	res := ref.Reflect(input)
	if len(res) != 1 || res[0] != input {
		t.Errorf("splitter | from above failed")
	}
}

func TestSplitterFromLeft(t *testing.T) {
	ref := Init('|')
	input := vec.Init(0, 1)
	res := ref.Reflect(input)
	if res[0] != vec.Init(1, 0) || res[1] != vec.Init(-1, 0) {
		t.Errorf("splitter | from above failed")
	}
}

func TestSplitterFromAboveFlat(t *testing.T) {
	ref := Init('-')
	input := vec.Init(1, 0)
	res := ref.Reflect(input)
	if len(res) != 2 || res[0] != vec.Init(0, 1) || res[1] != vec.Init(0, -1) {
		t.Errorf("splitter | from above failed")
	}
}
