package dsu

const noParent = ^uint(0)

type Dsu struct {
	Parent []uint
	Size   []uint
}

// Create a new Disjoint Set Union data structure with n elements.
// The elements are referred to by their index in the range [0, n).
// The element values are not stored explicitly in the data structure,
// and the user of the data structure is responsible for maintaining
// a correspondence between the element values and their indices.
func NewDsu(n uint) Dsu {
	dsu := Dsu{
		Parent: make([]uint, n),
		Size:   make([]uint, n),
	}
	for i := range dsu.Parent {
		dsu.Parent[i] = noParent
		dsu.Size[i] = 1
	}
	return dsu
}

func (d *Dsu) Find(v uint) uint {
	if d.Parent[v] == noParent {
		return v
	}
	d.Parent[v] = d.Find(d.Parent[v])
	return d.Parent[v]
}

func (d *Dsu) Union(v1, v2 uint) {
	v1 = d.Find(v1)
	v2 = d.Find(v2)
	if v1 == v2 {
		return
	}
	if d.Size[v1] < d.Size[v2] {
		v1, v2 = v2, v1
	}
	d.Parent[v2] = v1
	d.Size[v1] += d.Size[v2]
}
