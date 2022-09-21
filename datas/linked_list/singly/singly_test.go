package singly

import (
	"testing"
)

func TestSingly(t *testing.T) {
	s := Singly[int]{}
	s.Insert(1)
	s.Insert(2)
	s.Insert(3)
	s.Insert(4)
	s.Print()
	ok := s.Delete(3)
	ok = s.Delete(1)
	if !ok {
		t.Fail()
	}
	ok = s.Delete(10)
	if ok {
		t.Fail()
	}
	s.Print()
}
