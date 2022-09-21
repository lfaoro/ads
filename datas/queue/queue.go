package queue

import "sync"

// Queue is a restricted access collection, similar to a stack but
// organized after FIFO (first in, first out).
//
// The first inserted element in the queue is the first element to
// be removed.
//
// Queue is threadsafe, acquires a lock when needed.
type Queue[T comparable] struct {
	mu sync.Mutex
	e  []T
}

// Enque adds the element at the tail of the queue.
func (q *Queue[T]) Enque(e T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.e = append(q.e, e)
}

// Deque returns & removes the element from the head of the queue.
func (q *Queue[T]) Deque() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.Len() == 0 {
		var zero T
		return zero
	}

	element := q.e[0] // save first element
	q.e = q.e[1:]     // remove first element
	return element
}

// Search traverses the queue and returns the element found or nil
// w/o removing it from the queue if found.
func (q *Queue[T]) Search(e T) T {
	for _, element := range q.e {
		if element == e {
			return element
		}
	}
	var zero T
	return zero
}

// Head returns the first element in the queue w/o removing it.
func (q *Queue[T]) Head() T {
	if q.Len() == 0 {
		var zero T
		return zero
	}
	return q.e[0]
}

// Tail returns the last element in the queue w/o removing it
func (q *Queue[T]) Tail() T {
	if q.Len() == 0 {
		var zero T
		return zero
	}
	return q.e[len(q.e)-1]
}

// Len returns the length of the queue.
func (q *Queue[T]) Len() int {
	return len(q.e)
}

// EnqueFirst adds an element at the head of the queue and changing
// the queue priority.
func (q *Queue[T]) EnqueFirst(e T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.e = append([]T{e}, q.e...)
}
