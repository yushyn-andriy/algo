package graph

import (
	"errors"
	"math"
)

// BellmanFord ...
func BellmanFord(g []Edge, rootIdx int) (*map[int]float64, error) {
	if len(g) < rootIdx {
		return nil, errors.New("rootIdx out of range")
	}

	d := make(map[int]float64, len(g))

	for _, e := range g {
		d[e.B], d[e.A] = math.Inf(1), math.Inf(1)
	}
	d[g[rootIdx].A] = 0

	for i := 0; i <= len(g)-1; i++ {
		for _, e := range g {
			d[e.B] = math.Min(d[e.B], d[e.A]+e.W)
		}
	}

	return &d, nil
}
