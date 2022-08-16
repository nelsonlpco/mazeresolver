package dfs

import (
	"github.com/nelsonlpco/mazeresolver/domain/maze"
	"github.com/nelsonlpco/mazeresolver/domain/stack"
	"log"
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

// DepthFirstSearch percore o labirinto em busca do local de chegada, em caso de sucesso retorn o Node de sucesso
// em falha não retorna nada
func DepthFirstSearch(
	initial Node,
	goalTest func(state maze.Location) bool,
	successors func(state maze.Location) []maze.Location,
) *Node {
	frontier := stack.New[Node]()
	frontier.Push(initial)
	explored := []maze.Location{initial.State}

	for !frontier.IsEmpty() {
		currentNode := frontier.Pop()
		currentStat := currentNode.State

		if goalTest(currentStat) {
			return &currentNode
		}

		for _, location := range successors(currentStat) {
			if containsLocation(location, explored) {
				continue
			}

			explored = append(explored, location)
			frontier.Push(*New(&currentNode, location, 0.0, 0.0))
		}
	}

	return nil
}

// NodeToPath mapeia a rota a partir do nó principal
func NodeToPath(node Node) []maze.Location {
	log.Println("Mapping maze")
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

func containsLocation(el maze.Location, registers []maze.Location) bool {
	for _, register := range registers {
		if el == register {
			return true
		}
	}

	return false
}
