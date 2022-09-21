package singly

import (
	"fmt"
	"sync"
)

type node[T comparable] struct {
	value T
	next  *node[T]
}

type Singly[T comparable] struct {
	mu   sync.Mutex
	node *node[T]
}

// Insert traverses all the nodes until it finds an empty node where to store the value.
func (s *Singly[T]) Insert(v T) {
	if s.node == nil {
		s.node = &node[T]{value: v}
		return
	}

	var ptr = s.node
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = &node[T]{value: v}
}

// Delete traverses the linked list and removes the matching node.
func (s *Singly[T]) Delete(v T) bool {
	if s.node.value == v {
		s.node = s.node.next
		return true
	}

	var prev = s.node
	var ptr = s.node.next
	for ptr != nil {
		if ptr.value == v {
			prev.next = ptr.next
			ptr = nil
			return true
		}
		prev = ptr
		ptr = ptr.next
	}
	return false
}

// Search traverses the linked-list and returns the position where
// the element (v T) has been found or 0.
func (s *Singly[T]) Search(v T) int {
	var position = 0
	var ptr = s.node
	for ptr.next != nil {
		if ptr.value == v {
			return position
		}
		ptr = ptr.next
		position++
	}
	return 0
}

// Print traverses the linked list and prints the value of every node.
func (s *Singly[T]) Print() {
	var ptr = s.node
	for ptr != nil {
		fmt.Println(ptr.value)
		ptr = ptr.next
	}
}
