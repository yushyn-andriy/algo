package trees_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/strings"
	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestBuildHuffmanTree(t *testing.T) {
	s := "AAAAABBCDE"
	e := strings.CalculateRunesProbability(strings.CountRunes(s), len(s))
	ht := trees.BuildHuffmanTree(e)
	t.Errorf("%s", ht)
	t.Errorf("%s", ht.Left)
	t.Errorf("%s", ht.Left.Left)
	t.Errorf("%s", ht.Left.Right)

	t.Errorf("%s", ht.Right)
}
