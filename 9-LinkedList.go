package main

type LinkedList[T any] struct {
	Length int
	head   *LinkedListNode[T]
}

type LinkedListNode[T any] struct {
	value T
	prev  *LinkedListNode[T]
	next  *LinkedListNode[T]
}

func (lk *LinkedList[T]) Get(idx int) T {

	curr := lk.head
	for i := 0; i < idx && curr != nil; i++ {
		curr = curr.next
	}
	return curr.value
}
