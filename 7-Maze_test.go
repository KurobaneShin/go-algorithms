package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaze(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	mazeResult := []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}

	result := MazeSolver(maze, "x", Point{x: 10, y: 0}, Point{x: 1, y: 5})
	assert.Equal(t, DrawPath(maze, result), DrawPath(maze, mazeResult))
}

func DrawPath(data []string, path []Point) []string {
	data2 := make([][]string, len(data))
	for i, row := range data {
		data2[i] = strings.Split(row, "")
	}

	for _, p := range path {
		if p.y < len(data2) && p.x < len(data2[p.y]) {
			data2[p.y][p.x] = "*"
		}
	}

	result := make([]string, len(data2))
	for i, row := range data2 {
		result[i] = strings.Join(row, "")
	}

	return result
}
