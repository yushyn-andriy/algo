package tries

import "testing"

func TestSuffixTrie(t *testing.T) {
	tests := []struct {
		populateStr, containsStr string
		contains                 bool
	}{
		{
			"acdc",
			"dc",
			true,
		},
		{
			"acdc",
			"da",
			false,
		},
		{
			"algo and datastructure",
			"datastructure",
			true,
		},
	}

	for i, test := range tests {
		trie := NewSuffixTrie()
		trie.PopulateSuffixTrieFrom(test.populateStr)
		out := trie.Contains(test.containsStr)
		if out != test.contains {
			t.Errorf("test(%d) expected %v got %v", i, test.contains, out)
		}
	}

}
