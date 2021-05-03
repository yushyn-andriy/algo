package graph_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/graph"
)

func TestBellmanFord(t *testing.T) {
	g := []graph.Edge{
		{1, 2, 2},
		{4, 6, 1},
		{5, 6, 2},
		{1, 3, 3},
		{1, 4, 7},
		{2, 5, 5},
		{2, 4, 3},
		{3, 4, -2},
		{4, 5, 2},
	}
	expDistances := map[int]float64{
		1: 0,
		2: 2,
		3: 3,
		4: 1,
		5: 3,
		6: 2,
	}
	d, err := graph.BellmanFord(g, 0)
	if err != nil {
		t.Errorf("Expected error is nil, got %v", err)
	}

	for k, v := range expDistances {
		if ev, ok := (*d)[k]; !ok || ev != v {
			t.Errorf("Expected distance is %f, got %f, found in d %v", ev, v, ok)
		}
	}
}

func BenchmarkBellmanFord(b *testing.B) {
	g := []graph.Edge{
		{1, 2, 2},
		{1, 3, 3},
		{1, 4, 7},
		{2, 5, 5},
		{2, 4, 3},
		{3, 4, -2},
		{4, 5, 2},
	}

	for i := 0; i < b.N; i++ {
		graph.BellmanFord(g, 1)
	}
}
