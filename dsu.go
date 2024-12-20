package dsu

const noParent = ^uint(0)

type Dsu struct {
	parent []uint
	size   []uint
}

// Create a new Disjoint Set Union data structure with n elements.
// The elements are referred to by their index in the range [0, n).
// The element values are not stored explicitly in the data structure,
// and the user of the data structure is responsible for maintaining
// a correspondence between the element values and their indices.
func NewDsu(n uint) *Dsu {
	dsu := &Dsu{
		parent: make([]uint, n),
		size:   make([]uint, n),
	}
	for i := range dsu.parent {
		dsu.parent[i] = noParent
		dsu.size[i] = 1
	}
	return dsu
}

func (d *Dsu) Find(v uint) uint {
	if d.parent[v] == noParent {
		return v
	}
	d.parent[v] = d.Find(d.parent[v])
	return d.parent[v]
}

func (d *Dsu) Union(v1, v2 uint) {
	v1 = d.Find(v1)
	v2 = d.Find(v2)
	if v1 == v2 {
		return
	}
	if d.size[v1] < d.size[v2] {
		v1, v2 = v2, v1
	}
	d.parent[v2] = v1
	d.size[v1] += d.size[v2]
}
