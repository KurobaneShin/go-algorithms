package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearchList(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	assert.Equal(t, LinearSearchList(foo, 69), true)
	assert.Equal(t, LinearSearchList(foo, 1336), false)
}
