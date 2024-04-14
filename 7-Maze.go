package main

type Point struct {
	x int
	y int
}

var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func Walk(maze []string, wall string, curr Point, end Point, seen *[][]bool, path *[]Point) bool {
	// off the map
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	// on the wall
	if string(maze[curr.y][curr.x]) == wall {
		return false
	}

	if curr.y == end.y && curr.x == end.x {
		// add end to path
		*path = append(*path, end)
		return true
	}

	if (*seen)[curr.y][curr.x] {
		return false
	}

	// add path
	(*seen)[curr.y][curr.x] = true
	*path = append(*path, curr)

	// recurse
	for i := 0; i < len(dir); i++ {
		x, y := dir[i][0], dir[i][1]
		if Walk(maze, wall, Point{curr.x + x, curr.y + y}, end, seen, path) {
			return true
		}
	}

	// pop from path
	*path = (*path)[:len(*path)-1]
	return false
}

func MazeSolver(maze []string, wall string, start Point, end Point) []Point {
	seen := make([][]bool, len(maze))
	path := []Point{}

	for i := 0; i < len(maze); i++ {
		seen[i] = make([]bool, len(maze[0]))
	}

	Walk(maze, wall, start, end, &seen, &path)

	return path
}
