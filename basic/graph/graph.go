package graph

import (
	"fmt"
	"math/rand"
	"time"
)

func NewGraph(v int, directed bool, kind GraphType) (Graph, error) {
	switch kind {
	case Dense:
		adj := make([][]bool, v)
		for i := range adj {
			adj[i] = make([]bool, v)
		}
		return &DenseGraph{
			v:        v,
			directed: directed,
			adj:      adj,
			vcnt:     v,
			ecnt:     0,
		}, nil
	case Sparse:
		adj := make([][]int, v)
		for i := range adj {
			adj[i] = make([]int, 0)
		}
		return &SparseGraph{
			v:        v,
			directed: directed,
			adj:      adj,
			vcnt:     v,
			ecnt:     0,
		}, nil
	default:
		return nil, fmt.Errorf("Graph type %d not recognized", kind)
	}

}

func RandGraph(g Graph, e int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < e; i++ {
		v := rand.Intn(g.Vcnt())
		w := rand.Intn(g.Vcnt())
		g.Insert(&Edge{V: v, W: w})
	}
}
