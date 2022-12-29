package linkedlist

import (
	_ "errors"
)

type LinkedList struct {
	Key   int
	Data  string
	Value int
	Next  *LinkedList
	Prev  *LinkedList
}

type L struct {
	Head *LinkedList
	Size int
}

func Insert(ll *L, l *LinkedList) {
	if ll.Head == nil {
		ll.Head = l
	} else {
		ll.Head.Prev = l
		l.Next = ll.Head
		ll.Head = l
	}
	ll.Size++
}

func Find(ll *L, key int) *LinkedList {
	for x := ll.Head; x != nil; x = x.Next {
		if x.Key == key {
			return x
		}
	}
	return nil
}

func Delete(ll *L, key int) {
	x := Find(ll, key)
	if x != nil {
		if x.Prev != nil {
			x.Prev.Next = x.Next
		} else {
			ll.Head = x.Next
		}
		ll.Size--
	}
}

func Traverse(ll *L, f func(key int, data string)) {
	for x := ll.Head; x != nil; x = x.Next {
		f(x.Key, x.Data)
	}
}

func NewHead() *L {
	return &L{nil, 0}
}
