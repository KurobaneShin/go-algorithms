package main

type StackInterface[T any] struct {
	head   *SNode[T]
	tail   *SNode[T]
	length int
}

type SNode[T any] struct {
	value T
	prev  *SNode[T]
}

func (s *StackInterface[T]) Push(value T) {
	node := &SNode[T]{value: value}

	s.length++
	if s.head == nil {
		s.head = node
		return
	}

	node.prev = s.head
	s.head = node
}

func (s *StackInterface[T]) Pop() T {
	s.length = max(0, s.length-1)

	if s.length == 0 {
		head := s.head
		s.head = nil

		if head != nil {
			return head.value
		}

		var defaultValue T
		return defaultValue
	}

	head := s.head
	s.head = head.prev

	return head.value
}

func (s *StackInterface[T]) Peek() T {
	return s.head.value
}
