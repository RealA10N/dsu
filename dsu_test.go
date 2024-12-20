package dsu_test

import (
	"testing"

	"alon.kr/x/dsu"
)

func TestSimpleUnion(t *testing.T) {
	d := dsu.NewDsu(3)
	d.Union(0, 1)
	if d.Find(0) != d.Find(1) {
		t.Error("0 and 1 should be in the same set")
	}
	if d.Find(2) == d.Find(1) {
		t.Error("2 should not be in the same set as 0 and 1")
	}
}

func TestDoubleUnion(t *testing.T) {
	d := dsu.NewDsu(3)
	d.Union(0, 1)
	d.Union(1, 2)
	if d.Find(0) != d.Find(1) || d.Find(1) != d.Find(2) || d.Find(0) != d.Find(2) {
		t.Error("0, 1, and 2 should be in the same set")
	}
}

func TestUnionOfSameSet(t *testing.T) {
	d := dsu.NewDsu(2)
	d.Union(0, 1)
	d.Union(0, 1)
	if d.Find(0) != d.Find(1) {
		t.Error("0 and 1 should be in the same set")
	}
}
