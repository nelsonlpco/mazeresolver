package main

import (
	"fmt"
	"github.com/nelsonlpco/mazeresolver/domain/dfs"
	"github.com/nelsonlpco/mazeresolver/domain/maze"
	"log"
)

func main() {
	maze := maze.New(
		10,
		10,
		0.2,
		maze.Location{},
		maze.Location{Row: 9, Column: 9},
	)

	fmt.Println(maze)

	node := dfs.New(nil, maze.Start(), 0.0, 0.0)
	solution1 := dfs.DepthFirstSearch(*node, maze.GoalMatch, maze.Successors)

	if solution1 == nil {
		log.Println("No solution using deep first search")
		return
	}

	path := dfs.NodeToPath(*solution1)
	maze.Mark(path)
	fmt.Println(maze)
	maze.Clear(path)
}
