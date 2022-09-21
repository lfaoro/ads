package stack

import "sync"

// Stack is a generic, abstract and restricted data type
// which restriction is LIFO (last in, first out).
//
// Stack is generic(comparable) and thread safe.
type Stack[T comparable] struct {
	mu sync.Mutex
	e  []T
}

// Push adds the element e on top of the stack.
func (s *Stack[T]) Push(e T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.e = append(s.e, e)
}

// Pop returns & removes the first element from the stack.
func (s *Stack[T]) Pop() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.Len() == 0 {
		var zero T
		return zero
	}

	element := s.e[len(s.e)-1]
	s.e = s.e[:len(s.e)-1]
	return element
}

// Top returns the first element of the stack without removing it.
func (s *Stack[T]) Top() T {
	if s.Len() == 0 {
		var zero T
		return zero
	}
	return s.e[len(s.e)-1]
}

// Search returns true if is able to find (e T) in the stack, otherwise
// returns false.
func (s *Stack[T]) Search(e T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, element := range s.e {
		if element == e {
			return true
		}
	}
	return false
}

// Len returns the length of the stack.
func (s *Stack[T]) Len() int {
	return len(s.e)
}
