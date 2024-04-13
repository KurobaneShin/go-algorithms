package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {

	queue := QueueInterface[int]{}

	queue.Enqueue(5)
	queue.Enqueue(7)
	queue.Enqueue(9)

	assert.Equal(t, queue.Deque(), 5)
	assert.Equal(t, queue.length, 2)

	queue.Enqueue(11)

	assert.Equal(t, queue.Deque(), 7)
	assert.Equal(t, queue.Deque(), 9)
	assert.Equal(t, queue.Peek(), 11)

	assert.Equal(t, queue.Deque(), 11)

	//when queue is empty it should return the default value from especified type
	assert.Equal(t, queue.Deque(), 0)
	assert.Equal(t, queue.length, 0)

	//make sure list not died
	queue.Enqueue(69)
	assert.Equal(t, queue.Peek(), 69)
	assert.Equal(t, queue.length, 1)

}
