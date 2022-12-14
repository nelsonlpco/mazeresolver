package dfs

import (
	"github.com/nelsonlpco/mazeresolver/domain/collections"
	"github.com/nelsonlpco/mazeresolver/domain/collections/stack"
	"github.com/nelsonlpco/mazeresolver/domain/maze"
	"github.com/nelsonlpco/mazeresolver/domain/node"
)

// DepthFirstSearch percore o labirinto em busca do local de chegada, em caso de sucesso retorn o Node de sucesso
// em falha não retorna nada
func DepthFirstSearch(
	initial node.Node,
	goalTest func(state maze.Location) bool,
	successors func(state maze.Location) []maze.Location,
) *node.Node {
	frontier := stack.New[node.Node]()
	frontier.Push(initial)
	explored := []maze.Location{initial.State}

	for !frontier.IsEmpty() {
		currentNode := frontier.Pop()
		currentStat := currentNode.State

		if goalTest(currentStat) {
			return &currentNode
		}

		for _, location := range successors(currentStat) {
			if collections.Contains(location, explored) {
				continue
			}

			explored = append(explored, location)
			frontier.Push(*node.New(&currentNode, location, 0.0, 0.0))
		}
	}

	return nil
}
