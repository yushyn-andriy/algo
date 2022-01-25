package stack

import "testing"

func TestStack(t *testing.T) {
	st := New()
	if st.Size() != 0 {
		t.Errorf("wrong stack initialization expected size %v got %v", 0, st.Size())
	}

	st.Push(1)
	st.Push(2)
	st.Push(3)

	peek := st.Peek()
	if peek == nil || *peek != 3 {
		t.Errorf("wrong peek value expected %v got %v", 3, peek)
	}

	v1, v2, v3 := st.Pop(), st.Pop(), st.Pop()
	if v1 == nil || v2 == nil || v3 == nil {
		t.Errorf("wrong pop expected non nil %v, %v, %v", v1, v2, v3)
		return
	}

	if *v1 != 3 {
		t.Errorf("wrong pop expected %v got %v", v1, 3)
	}

	if *v2 != 2 {
		t.Errorf("wrong pop expected %v got %v", v2, 2)
	}

	if *v3 != 1 {
		t.Errorf("wrong pop expected %v got %v", v3, 1)
	}

	n := st.Pop()
	if n != nil {
		t.Errorf("wrong pop expected %v got %v", n, nil)
	}

	if st.Size() != 0 {
		t.Errorf("wrong size expected %v got %v", 0, st.Size())
	}
}
