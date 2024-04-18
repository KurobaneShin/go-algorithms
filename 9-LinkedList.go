package main

type LinkedList[T any] struct {
	Length int
	head   *LinkedListNode[T]
	tail   *LinkedListNode[T]
}

type LinkedListNode[T any] struct {
	value T
	prev  *LinkedListNode[T]
	next  *LinkedListNode[T]
}

func (lk *LinkedList[T]) InsertAt(item T, at int) {
	if at > lk.Length {
		panic("cant do")
	}

	if at == lk.Length {
		lk.Append(item)
		return
	} else if at == 0 {
		lk.Prepend(item)
		return
	}

	lk.Length++

	curr := lk.head

	for i := 0; curr != nil && i < at; i++ {
		curr = curr.next
	}

	node := &LinkedListNode[T]{value: item}

	node.next = curr
	node.prev = curr.prev
	curr.prev = node

	if curr.prev != nil {
		curr.prev.next = curr
	}
}

func (lk *LinkedList[T]) Append(item T) {
	lk.Length++

	node := &LinkedListNode[T]{
		value: item,
	}

	if lk.tail == nil {
		lk.head, lk.tail = node, node
		return
	}

	node.prev = lk.tail
	lk.tail.next = node
}

func (lk *LinkedList[T]) Prepend(item T) {
	node := &LinkedListNode[T]{value: item}

	lk.Length++

	if lk.head == nil {
		lk.head, lk.tail = node, node
		return
	}

	node.next = lk.head
	lk.head.prev = node
	lk.head = node
}

func (lk *LinkedList[T]) Get(idx int) T {
	return lk.GetAt(idx).value
}

func (lk *LinkedList[T]) GetAt(idx int) *LinkedListNode[T] {
	if idx > lk.Length {
		panic("cant do")
	}

	curr := lk.head

	for i := 0; i < idx && curr != nil; i++ {
		curr = curr.next
	}

	return curr
}

func (lk *LinkedList[T]) RemoveAt(idx int) T {
	node := lk.GetAt(idx)

	if node != nil {
		var defaultValue T
		return defaultValue
	}

	return lk.RemoveNode(node)
}

func (lk *LinkedList[T]) RemoveNode(node *LinkedListNode[T]) T {
	lk.Length--

	if lk.Length == 0 {
		out := lk.head.value
		lk.head, lk.tail = nil, nil

		return out
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == lk.head {
		lk.head = node.next
	}

	if node == lk.tail {
		lk.tail = node.prev
	}

	node.prev, node.next = nil, nil

	return node.value
}
