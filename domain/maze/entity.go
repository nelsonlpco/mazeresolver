package maze

import (
	"fmt"
	"github.com/nelsonlpco/mazeresolver/domain/cell"
	"math/rand"
)

type Location struct {
	Row    int
	Column int
}

type Entity struct {
	rows    int
	columns int
	start   Location
	goal    Location
	grid    [][]cell.Type
}

func (m Entity) String() string {
	output := ""
	for row := 0; row < m.rows; row++ {
		for column := 0; column < m.columns; column++ {
			output = fmt.Sprintf("%v %v ", output, m.grid[row][column])
		}
		output = fmt.Sprintf("%v%v", output, "\n")
	}
	return output
}

func (m Entity) GoalMatch(ml Location) bool {
	return ml == m.goal
}

func (m Entity) Successors(l Location) []Location {
	var locations []Location

	bottomIsEmpty := l.Row+1 < m.rows && m.grid[l.Row+1][l.Column] != cell.BLOCKED
	if bottomIsEmpty {
		locations = append(locations, Location{Row: l.Row + 1, Column: l.Column})
	}

	topIsEmpty := l.Row-1 >= 0 && m.grid[l.Row-1][l.Column] != cell.BLOCKED
	if topIsEmpty {
		locations = append(locations, Location{Row: l.Row - 1, Column: l.Column})
	}

	rightIsEmpty := l.Column+1 < m.rows && m.grid[l.Row][l.Column+1] != cell.BLOCKED
	if rightIsEmpty {
		locations = append(locations, Location{Row: l.Row, Column: l.Column + 1})
	}

	leftIsEmpty := l.Column-1 >= 0 && m.grid[l.Row][l.Column-1] != cell.BLOCKED
	if leftIsEmpty {
		locations = append(locations, Location{Row: l.Row, Column: l.Column - 1})
	}

	return locations
}

func (m Entity) Mark(path []Location) {
	for _, loc := range path {
		m.grid[loc.Row][loc.Column] = cell.PATH
	}
	m.grid[m.goal.Row][m.goal.Column] = cell.GOAL
	m.grid[m.start.Row][m.start.Column] = cell.START
}

func (m Entity) Clear(path []Location) {
	for _, loc := range path {
		m.grid[loc.Row][loc.Column] = cell.EMPTY
	}

	m.grid[m.goal.Row][m.goal.Column] = cell.GOAL
	m.grid[m.start.Row][m.start.Column] = cell.START
}

func (m Entity) Start() Location {
	return m.start
}

func New(rows int, columns int, sparseness float32, start Location, goal Location) *Entity {
	maze := &Entity{rows: rows, columns: columns, start: start, goal: goal}

	maze.grid = randomlyFill(rows, columns, sparseness)
	maze.grid[maze.goal.Row][maze.goal.Column] = cell.GOAL
	maze.grid[maze.start.Row][maze.start.Column] = cell.START

	return maze
}

func randomlyFill(rows int, columns int, sparseness float32) [][]cell.Type {
	grid := make([][]cell.Type, rows)
	for row := 0; row < rows; row++ {
		grid[row] = make([]cell.Type, columns)
		for column := 0; column < columns; column++ {
			s := rand.Float32()
			if s < sparseness {
				grid[row][column] = cell.BLOCKED
			} else {
				grid[row][column] = cell.EMPTY
			}
		}
	}

	return grid
}
