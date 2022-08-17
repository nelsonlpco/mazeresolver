package main

import (
	"fmt"
	"github.com/nelsonlpco/mazeresolver/domain/bfs"
	"github.com/nelsonlpco/mazeresolver/domain/dfs"
	"github.com/nelsonlpco/mazeresolver/domain/maze"
	"github.com/nelsonlpco/mazeresolver/domain/node"
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

	root := node.New(nil, maze.Start(), 0.0, 0.0)
	solution1 := dfs.DepthFirstSearch(*root, maze.GoalMatch, maze.Successors)

	if solution1 == nil {
		log.Println("No solution using deep first search")
		return
	}

	fmt.Println("solution DFS")
	path := node.NodeToPath(*solution1)
	fmt.Println("Size ", len(path))
	maze.Mark(path)
	fmt.Println(maze)
	maze.Clear(path)

	rootBfs := node.New(nil, maze.Start(), 0.0, 0.0)
	solution2 := bfs.BreadthFirstSearch(*rootBfs, maze.GoalMatch, maze.Successors)

	fmt.Println("solution BFS")
	path = node.NodeToPath(*solution2)
	fmt.Println("Size ", len(path))
	maze.Mark(path)
	fmt.Println(maze)
	maze.Clear(path)
}
