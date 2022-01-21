package arrays

import "testing"

func TestApartmentHunting(t *testing.T) {
	cases := []struct {
		blocks []Block
		reqs   []string
		idx    int
	}{
		{
			[]Block{
				{
					"gym":    false,
					"school": true,
					"store":  false,
				},
				{
					"gym":    true,
					"school": false,
					"store":  false,
				},
				{
					"gym":    true,
					"school": true,
					"store":  false,
				},
				{
					"gym":    false,
					"school": true,
					"store":  false,
				},
				{
					"gym":    false,
					"school": true,
					"store":  true,
				},
			},
			[]string{"gym", "school", "store"},
			3,
		},
	}

	for i, c := range cases {
		idx := ApartmentHunting(c.blocks, c.reqs)
		if idx != c.idx {
			t.Errorf("%d error expected %v got %v", i, c.idx, idx)
		}
	}
}
