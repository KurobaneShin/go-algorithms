package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {

	arr := []int{9, 3, 7, 4, 69, 420, 42}

	ordered := []int{3, 4, 7, 9, 42, 69, 420}

	assert.Equal(t, ordered, BubbleSort(arr))
}
