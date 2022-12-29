package graph

import "testing"

func TestSparseMatrixSearchR(t *testing.T) {
	tests := []struct {
		edges    [][]int
		vertexes int
		v, w     int
		search   bool
	}{
		{
			[][]int{
				{0, 1},
				{1, 2},
				{2, 3},
			},
			4,
			0, 3,
			true,
		},
		{
			[][]int{
				{0, 1},
				{1, 2},
			},
			4,
			0, 3,
			false,
		},
	}

	for i, test := range tests {
		g, _ := NewGraph(test.vertexes, false, Sparse)
		for t := range test.edges {
			vw := test.edges[t]
			g.Insert(&Edge{
				V: vw[0],
				W: vw[1],
			})
		}
		out := g.SearchR(test.v, test.w)
		if out != test.search {
			t.Errorf("test(%d) expected %v got %v", i, test.search, out)
		}
	}
}
