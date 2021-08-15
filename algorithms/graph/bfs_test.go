package graph

import (
	"reflect"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	tests := []struct {
		in  *Node1
		out []string
	}{
		{
			&Node1{
				Name: "A",
				Children: []*Node1{
					{
						Name: "B",
					},
					{
						Name: "C",
					},
				},
			},
			[]string{"A", "B", "C"},
		},
	}

	for i, test := range tests {
		out := test.in.BreadthFirstSearch([]string{})
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
		}
	}
}
