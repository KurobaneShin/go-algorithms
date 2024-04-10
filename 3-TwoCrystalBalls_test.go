package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoCrystalBalls(t *testing.T) {

	// Generate a random index between 0 and 9999 (inclusive)
	idx := rand.Intn(10000)

	// Create a slice with 10000 elements, initialized to false
	data := make([]bool, 10000)

	// Set elements from idx to 9999 to true
	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	assert.Equal(t, idx, TwoCrystalBalls(data))
	assert.Equal(t, -1, TwoCrystalBalls(make([]bool, 821)))
}
