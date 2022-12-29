package main

import (
	"testing"
)

func TestNumberOfDivisors(t *testing.T) {
	tests := []struct {
		k int // number
		n int // number of divisors
	}{
		{
			1,
			1,
		},
		{
			2,
			2,
		},
		{
			3,
			2,
		},
		{
			4,
			3,
		},
		{
			6,
			4,
		},
		{
			12,
			6,
		},
		{
			24,
			8,
		},
		{
			36,
			9,
		},
		{
			48,
			10,
		},
		{
			60,
			12,
		},
		{
			120,
			16,
		},
		{
			180,
			18,
		},
		{
			240,
			20,
		},
		{
			360,
			24,
		},
		{
			720,
			30,
		},
		{ // #####
			840,
			32,
		},
		{
			1260,
			36,
		},
		{
			1680,
			40,
		},
		{
			2520,
			48,
		},
		{
			5040,
			60,
		},
		{
			7560,
			64,
		},
		{
			10080,
			72,
		},
		{
			15120,
			80,
		},
		{
			20160,
			84,
		},
		{
			25200,
			90,
		},
		{
			262144,
			19,
		},
		{
			27720,
			96,
		},
		{
			45360,
			100,
		},
		{
			50400,
			108,
		},
		{
			55440,
			120,
		},
		{
			83160,
			128,
		},
		{
			110880,
			144,
		},
		{
			166320,
			160,
		},
		{
			221760,
			168,
		},
		{
			277200,
			180,
		},
		{
			332640,
			192,
		},
		{
			498960,
			200,
		},
		{
			554400,
			216,
		},
		{
			665280,
			224,
		},
		{
			720720,
			240,
		},
	}

	for i, test := range tests {
		n := NumberOfDivisors(test.k)
		if n != test.n {
			t.Errorf("test(%d) Expected %d got %d", i, test.n, n)
		}
	}
}

/*
func TestFactoringNumbers(t *testing.T) {
	tests := []struct {
		k int   // number
		f []int // factoring numbers
	}{
		{
			1,
			[]int{},
		},
		{
			2,
			[]int{2},
		},
		{
			3,
			[]int{3},
		},
		{
			4,
			[]int{2, 2},
		},
		{
			6,
			[]int{2, 3},
		},
		{
			240,
			[]int{2, 2, 2, 2, 3, 5},
		},
	}

	for i, test := range tests {
		f := FactoringNumbers(test.k)
		if !reflect.DeepEqual(test.f, f) {
			t.Errorf("test(%d) Expected %v got %v", i, test.f, f)
		}
	}
}

func TestCalculateT(t *testing.T) {
	tests := []struct {
		f []int // factoring numbers
		k map[int]int
	}{
		{
			[]int{},
			map[int]int{},
		},
		{
			[]int{2},
			map[int]int{2: 1},
		},
		{
			[]int{3},
			map[int]int{3: 1},
		},
		{
			[]int{2, 2},
			map[int]int{2: 2},
		},
		{
			[]int{2, 3},
			map[int]int{2: 1, 3: 1},
		},
		{
			[]int{2, 2, 2, 2, 3, 5},
			map[int]int{2: 4, 3: 1, 5: 1},
		},
	}

	for i, test := range tests {
		k := CalculateT(test.f)
		if !reflect.DeepEqual(test.k, k) {
			t.Errorf("test(%d) Expected %v got %v", i, test.k, k)
		}
	}
}

func TestSearchSubsets(t *testing.T) {
	tests := []struct {
		f []int // factoring numbers
		k [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				{1, 2, 3},
				{1, 2},
				{1, 3},
				{1},
				{2, 3},
				{2},
				{3},
				{},
			},
		},
	}

	for i, test := range tests {
		k := SearchSubsets(test.f)
		if !reflect.DeepEqual(test.k, k) {
			t.Errorf("test(%d) Expected %v got %v", i, test.k, k)
		}
	}
}

func TestSubstract(t *testing.T) {
	tests := []struct {
		x []int // factoring numbers
		y []int
		z []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1},
			[]int{2, 3},
		},
		{
			[]int{2, 2, 2, 3, 5},
			[]int{2, 2, 3},
			[]int{2, 5},
		},
		{
			[]int{1, 2, 3},
			[]int{2},
			[]int{1, 3},
		},
		{
			[]int{1, 2, 3},
			[]int{},
			[]int{1, 2, 3},
		},
	}

	for i, test := range tests {
		z := Substract(test.x, test.y)
		if !reflect.DeepEqual(test.z, z) {
			t.Errorf("test(%d) Expected %v got %v", i, test.z, z)
		}
	}
}
*/
