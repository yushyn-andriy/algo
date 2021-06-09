package trees_test

import (
	"fmt"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/strings"
	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestBuildHuffmanTree(t *testing.T) {
	s := "AAAABBCCDE"
	e := strings.CalculateRunesProbability(strings.CountRunes(s), len(s))
	ht := trees.BuildHuffmanTree(e)

	exps := ""
	trees.HuffmanInOrderWalk(ht, func(tree *trees.HuffmanNode) {
		exps += fmt.Sprintf("%s:", tree.Value)
	})
	if exps != "A:AB:B:ABC:C:ABCD:D:ABCDE:E:" {
		t.Errorf("%s", exps)
	}
}
