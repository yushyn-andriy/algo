package tries

import (
	"reflect"
	"testing"
)

func TestMultiStringSearch1(t *testing.T) {
	cases := []struct {
		bigString    string
		smallStrings []string
		expected     []bool
	}{
		{
			"this is a big string",
			[]string{"this", "yo", "is", "a", "bigger", "string", "kappa"},
			[]bool{true, false, true, true, false, true, false},
		},
	}

	for i, testCase := range cases {
		out := MultiStringSearch1(testCase.bigString, testCase.smallStrings)
		if !reflect.DeepEqual(out, testCase.expected) {
			t.Errorf("%d error expected %v but got %v", i, testCase.expected, out)
		}
	}
}

func TestMultiStringSearch2(t *testing.T) {
	cases := []struct {
		bigString    string
		smallStrings []string
		expected     []bool
	}{
		{
			"this is a big string",
			[]string{"this", "yo", "is", "a", "bigger", "string", "kappa"},
			[]bool{true, false, true, true, false, true, false},
		},
	}

	for i, testCase := range cases {
		out := MultiStringSearch2(testCase.bigString, testCase.smallStrings)
		if !reflect.DeepEqual(out, testCase.expected) {
			t.Errorf("%d error expected %v but got %v", i, testCase.expected, out)
		}
	}
}

func TestMultiStringSearch3(t *testing.T) {
	cases := []struct {
		bigString    string
		smallStrings []string
		expected     []bool
	}{
		{
			"this is a big string",
			[]string{"this", "yo", "is", "a", "bigger", "string", "kappa"},
			[]bool{true, false, true, true, false, true, false},
		},
	}

	for i, testCase := range cases {
		out := MultiStringSearch3(testCase.bigString, testCase.smallStrings)
		if !reflect.DeepEqual(out, testCase.expected) {
			t.Errorf("%d error expected %v but got %v", i, testCase.expected, out)
		}
	}
}
