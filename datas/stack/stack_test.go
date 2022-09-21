package stack

import "testing"

func TestStack(t *testing.T) {
	s := Stack[int]{}

	v := s.Pop()
	if v != 0 {
		t.Fail()
	}

	for i := 1; i <= 100; i++ {
		s.Push(i)
	}

	v = s.Top()
	if v != 100 {
		t.Fail()
	}

	switch {
	case s.Search(50) != true:
		t.Fail()
	case s.Search(150) != false:
		t.Fail()
	}

	for i := 100; i > 0; i-- {
		v := s.Pop()
		if v != i {
			t.Fatal(v, i)
		}
	}

	v = s.Top()
	if v != 0 {
		t.Fatal(v)
	}

}
