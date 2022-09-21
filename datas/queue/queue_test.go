package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue[int]{}

	// add to queue
	for i := 100; i > 0; i-- {
		q.Enque(i)
	}

	q.EnqueFirst(101)

	v := q.Search(50)
	if v != 50 {
		t.Fatal()
	}
	v = q.Search(150)
	if v != 0 {
		t.Fatal()
	}

	v = q.Head()
	if v != 101 {
		t.Fatal()
	}

	v = q.Tail()
	if v != 1 {
		t.Fatal(v)
	}

	// remove from queue
	for i := 101; i >= 0; i-- {
		v := q.Deque()
		if v != i {
			t.Fatal(v)
		}
	}

	// queue overflow
	v = q.Deque()
	if v != 0 {
		t.Fatal(v)
	}

	v = q.Head()
	if v != 0 {
		t.Fatal(v)
	}
	v = q.Tail()
	if v != 0 {
		t.Fatal(v)
	}

}

var q = Queue[int]{}

func BenchmarkEnque(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q.Enque(i)
	}
}

func BenchmarkDeque(b *testing.B) {
	BenchmarkEnque(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if q.Len() == 0 {
			b.Logf("seq: %v", i)
			break
		}
		q.Deque()
	}
}
