package dfs

import (
	"github.com/nelsonlpco/mazeresolver/domain/maze"
)

type Node struct {
	State     maze.Location
	Cost      float32
	Heuristic float32
	Parent    *Node
}

func (s Node) LT(other Node) bool {
	return (s.Cost + s.Heuristic) < (other.Cost + other.Heuristic)
}

func New(parent *Node, state maze.Location, cost float32, heuristic float32) *Node {
	return &Node{Parent: parent, State: state, Cost: cost, Heuristic: heuristic}
}

// NodeToPath mapeia a rota a partir do nó principal
func NodeToPath(node Node) []maze.Location {
	path := []maze.Location{node.State}

	parent := node.Parent
	for {
		if parent == nil {
			break
		}

		path = append(path, parent.State)
		parent = parent.Parent
	}

	return path
}
