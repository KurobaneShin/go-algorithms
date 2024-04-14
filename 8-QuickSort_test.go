package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	ordered := []int{3, 4, 7, 9, 42, 69, 420}
	QuickSort(&arr)
	assert.Equal(t, ordered, arr)
}
