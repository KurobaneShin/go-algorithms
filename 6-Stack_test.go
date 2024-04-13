package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	queue := StackInterface[int]{}

	queue.Push(5)
	queue.Push(7)
	queue.Push(9)

	assert.Equal(t, queue.Pop(), 9)
	assert.Equal(t, queue.length, 2)

	queue.Push(11)

	assert.Equal(t, queue.Pop(), 11)
	assert.Equal(t, queue.Pop(), 7)
	assert.Equal(t, queue.Peek(), 5)

	assert.Equal(t, queue.Pop(), 5)

	// when queue is empty it should return the default value from especified type
	assert.Equal(t, queue.Pop(), 0)
	assert.Equal(t, queue.length, 0)
	//
	// //make sure stack not died
	queue.Push(69)
	assert.Equal(t, queue.Peek(), 69)
	assert.Equal(t, queue.length, 1)
	//
}
