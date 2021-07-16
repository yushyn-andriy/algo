package graph_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/graph"
)

func TestReconstructBst(t *testing.T) {
	tests := []struct {
		bst graph.Node
		out []string
	}{
		{
			graph.Node{
				Name: "A",
				Children: []*graph.Node{
					{
						Name: "B",
						Children: []*graph.Node{
							{
								Name:     "E",
								Children: nil,
							},
							{
								Name: "F",
								Children: []*graph.Node{
									{
										Name:     "I",
										Children: nil,
									},
									{
										Name:     "J",
										Children: nil,
									},
								},
							},
						},
					},
					{
						Name:     "C",
						Children: nil,
					},
					{
						Name: "D",
						Children: []*graph.Node{
							{
								Name: "G",
								Children: []*graph.Node{
									{
										Name:     "K",
										Children: nil,
									},
								},
							},
							{
								Name:     "H",
								Children: nil,
							},
						},
					},
				},
			},
			[]string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"},
		},
	}

	for i, test := range tests {
		out := test.bst.DepthFirstSearch([]string{})
		if !reflect.DeepEqual(test.out, out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}

	}
}
