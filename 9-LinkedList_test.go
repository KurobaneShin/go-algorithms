package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	linkedList := LinkedList[int]{}

	linkedList.Append(5)
	linkedList.Append(7)
	linkedList.Append(9)

	assert.Equal(t, linkedList.Get(2), 9)
	assert.Equal(t, 7, linkedList.RemoveAt(1))
	assert.Equal(t, 2, linkedList.Length)

	linkedList.Append(11)
	assert.Equal(t, 9, linkedList.RemoveAt(1))
	assert.Equal(t, 0, linkedList.Remove(9))
	assert.Equal(t, 5, linkedList.RemoveAt(0))
	assert.Equal(t, 11, linkedList.RemoveAt(0))

	linkedList.Prepend(5)
	linkedList.Prepend(7)
	linkedList.Prepend(9)

	assert.Equal(t, 5, linkedList.Get(2))
	assert.Equal(t, 9, linkedList.Get(0))
	assert.Equal(t, 9, linkedList.Remove(9))
	assert.Equal(t, 2, linkedList.Length)
	assert.Equal(t, 7, linkedList.Get(0))
}
