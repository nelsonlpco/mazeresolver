package bfs

import (
	"github.com/nelsonlpco/mazeresolver/domain/collections"
	"github.com/nelsonlpco/mazeresolver/domain/collections/queue"
	"github.com/nelsonlpco/mazeresolver/domain/maze"
	"github.com/nelsonlpco/mazeresolver/domain/node"
)

func BreadthFirstSearch(
	initial node.Node,
	goalTest func(state maze.Location) bool,
	successors func(state maze.Location) []maze.Location,
) *node.Node {
	frontier := queue.New[node.Node]()
	frontier.Push(initial)
	explored := []maze.Location{initial.State}

	for !frontier.IsEmpty() {
		currentNode := frontier.Pop()
		currentState := currentNode.State

		if goalTest(currentState) {
			return &currentNode
		}

		for _, location := range successors(currentState) {
			if collections.Contains(location, explored) {
				continue
			}

			explored = append(explored, location)
			frontier.Push(*node.New(&currentNode, location, 0.0, 0.0))
		}
	}

	return nil
}
