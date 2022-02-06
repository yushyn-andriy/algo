package heapsort

import (
	"reflect"
	"testing"
)

func TestLeft(t *testing.T) {
	cases := []struct {
		i, left int
	}{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 6},
	}

	for i, c := range cases {
		left := Left(c.i)
		if left != c.left {
			t.Errorf("%d error expected %v but got %v", i, c.left, left)
		}
	}
}

func TestRight(t *testing.T) {
	cases := []struct {
		i, right int
	}{
		{0, 1},
		{1, 3},
		{2, 5},
		{3, 7},
	}

	for i, c := range cases {
		right := Right(c.i)
		if right != c.right {
			t.Errorf("%d error expected %v but got %v", i, c.right, right)
		}
	}
}

func TestParent(t *testing.T) {
	cases := []struct {
		i, parent int
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
	}

	for i, c := range cases {
		parent := Parent(c.i)
		if parent != c.parent {
			t.Errorf("%d error expected %v but got %v", i, c.parent, parent)
		}
	}

}

func TestMaxHeapify(t *testing.T) {
	cases := []struct {
		arr, expected []int
	}{
		{
			[]int{0, 3, 2, 1},
			[]int{0, 2, 3, 1},
		},
		{
			[]int{0, 3, 5, 1},
			[]int{0, 1, 5, 3},
		},
		{
			[]int{0, 16, 4, 10, 14, 7, 9, 3, 2, 8, 1},
			[]int{0, 16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
		},
	}
	for i, c := range cases {
		h := *NewHeap(c.arr)
		MaxHeapify(h, 1)
		arr := []int(h)
		if reflect.DeepEqual(arr, c.expected) {
			t.Errorf("%d error expected %v but got %v", i, c.expected, arr)
		}
	}
}

func TestBuildMaxHeap(t *testing.T) {
	cases := []struct {
		arr, expected []int
	}{
		{
			[]int{0, 3, 2, 1},
			[]int{0, 2, 3, 1},
		},
		{
			[]int{0, 3, 5, 1},
			[]int{0, 1, 5, 3},
		},
		{
			[]int{0, 16, 4, 10, 14, 7, 9, 3, 2, 8, 1},
			[]int{0, 16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
		},
		{
			[]int{0, 4, 1, 3, 2, 16, 9, 10, 14, 8, 7},
			[]int{0, 16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
		},
	}
	for i, c := range cases {
		h := *NewHeap(c.arr)
		BuildMaxHeap(h)
		arr := []int(h)
		if reflect.DeepEqual(arr, c.expected) {
			t.Errorf("%d error expected %v but got %v", i, c.expected, arr)
		}
	}
}
