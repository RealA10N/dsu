package dsu_test

import (
	"fmt"
	"math/rand"
	"testing"

	"alon.kr/x/dsu"
)

// MARK: Tests

func assertAllSame(t *testing.T, values ...uint) {
	for i := 1; i < len(values); i++ {
		if values[i] != values[0] {
			t.Error("All values should be the same")
		}
	}
}

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
	assertAllSame(t, d.Find(0), d.Find(1), d.Find(2))
}

func TestUnionOfSameSet(t *testing.T) {
	d := dsu.NewDsu(2)
	d.Union(0, 1)
	d.Union(0, 1)
	if d.Find(0) != d.Find(1) {
		t.Error("0 and 1 should be in the same set")
	}
}

func TestReverseOrderUnion(t *testing.T) {
	d := dsu.NewDsu(3)
	d.Union(1, 2)
	d.Union(0, 1)
	assertAllSame(t, d.Find(0), d.Find(1), d.Find(2))
}

// MARK: Benchmarks

type unionQuery struct {
	v1, v2 uint
}

func createUnionQueries(dsuSize int) []unionQuery {
	queries := make([]unionQuery, 0)
	for setSize := 2; setSize < dsuSize; setSize *= 2 {
		oldSetSize := setSize / 2
		for i := 0; i+setSize <= dsuSize; i += setSize {
			a := i + rand.Intn(oldSetSize)
			b := i + oldSetSize + rand.Intn(oldSetSize)
			queries = append(queries, unionQuery{uint(a), uint(b)})
		}
	}

	rand.Shuffle(len(queries), func(i, j int) {
		queries[i], queries[j] = queries[j], queries[i]
	})

	return queries
}

const dsuSmallestTestSize = 0x10000

func calculateRequiredDsuSize(minimumQueryAmount int) int {
	dsuSize := dsuSmallestTestSize
	for dsuSize < minimumQueryAmount {
		dsuSize *= 2
	}
	return dsuSize
}

func BenchmarkRandomUnions(b *testing.B) {
	dsuSize := calculateRequiredDsuSize(b.N)
	dsu := dsu.NewDsu(uint(dsuSize))
	queries := createUnionQueries(dsuSize)
	b.ResetTimer()
	for bi := 0; bi < b.N; bi++ {
		dsu.Union(queries[bi].v1, queries[bi].v2)
	}
}

// MARK: Example

func ExampleDsu() {
	d := dsu.NewDsu(3)
	d.Union(0, 1)

	if d.Find(0) == d.Find(1) {
		fmt.Println("0 and 1 are in the same set")
	}

	if d.Find(1) != d.Find(2) {
		fmt.Println("1 and 2 are in different sets")
	}

	if d.Size(0) == 2 {
		fmt.Println("Set with the element 0 has 2 elements")
	}

	// Output:
	// 0 and 1 are in the same set
	// 1 and 2 are in different sets
	// Set with the element 0 has 2 elements
}
