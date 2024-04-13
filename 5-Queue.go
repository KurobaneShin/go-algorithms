package main

type QueueInterface[T any] struct {
	head   *QNode[T]
	tail   *QNode[T]
	length int
}

type QNode[T any] struct {
	value T
	next  *QNode[T]
}

func (q *QueueInterface[T]) Enqueue(value T) {
	pointer := &QNode[T]{value: value}
	q.length++

	if q.length == 1 {
		q.tail = pointer
		q.head = pointer
		return
	}

	q.tail.next = pointer
	q.tail = pointer
}

func (q *QueueInterface[T]) Peek() T {
	return q.head.value
}

func (q *QueueInterface[T]) Deque() T {

	if q.head == nil {
		var defaultValue T
		return defaultValue
	}

	q.length--

	head := q.head

	q.head = q.head.next

	head.next = nil

	if q.length == 0 {
		q.tail = nil
	}

	return head.value

}
