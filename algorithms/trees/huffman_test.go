package trees_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/strings"
	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestBuildHuffmanTree(t *testing.T) {
	s := "AAAABBCCDE"

	e := strings.CalculateRunesProbability(strings.CountRunes(s), len(s))
	ht := trees.BuildHuffmanTree(e)

	// A:ABCED:B:BCED:C:CED:E:ED:D:
	// A:ABCED:B:BCED:C:CED:E:ED:D:
	// A:ACBDE:C:CBDE:B:BDE:D:DE:E:
	// A:ABCED:B:BCED:C:CED:E:ED:D:
	// A:ACBDE:C:CBDE:B:BDE:D:DE:E:
	exps := ""
	trees.HuffmanInOrderWalk(ht, func(tree *trees.HuffmanNode) {
		exps += fmt.Sprintf("%s:", tree.Value)
	})

	expected := "A:ABCDE:B:BCDE:C:CDE:D:DE:E:"
	if !(exps != expected || exps != "A:ABCED:B:BCED:C:CED:E:ED:D:") {
		t.Errorf("Expected %s, got %s", expected, exps)
	}
}

func TestHuffmanCodesFromTree(t *testing.T) {
	// s := "AAAABBCCDE"
	// e := strings.CalculateRunesProbability(strings.CountRunes(s), len(s))
	ht := trees.BuildHuffmanTree(map[rune]float64{
		'a': 45.0,
		'b': 13.0,
		'c': 12.0,
		'd': 16.0,
		'e': 9.0,
		'f': 5.0,
	})
	hc := trees.HuffmanCodesFromTree(ht)
	expected := map[string]string{
		"a": "0",
		"b": "101",
		"c": "100",
		"d": "111",
		"e": "1101",
		"f": "1100",
	}
	if !reflect.DeepEqual(hc, expected) {
		t.Errorf("Expected %+v but got %v", expected, hc)

	}
}
