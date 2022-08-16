package collections

import (
	"github.com/nelsonlpco/mazeresolver/domain/maze"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("should returns true when i search 3 in [1,2,3,4,5]", func(t *testing.T) {
		elements := []int{1, 2, 3, 4, 5}
		element := 3

		assert.True(t, Contains(element, elements))
	})

	t.Run("should returns false when i search 6 in [1,2,3,4,5]", func(t *testing.T) {
		elements := []int{1, 2, 3, 4, 5}
		element := 6

		assert.False(t, Contains(element, elements))
	})

	t.Run("should returns true when i find a [1,2] in location collection [[2,3],[1,2],[0,0]]", func(t *testing.T) {
		locations := []maze.Location{
			{Row: 2, Column: 3},
			{Row: 1, Column: 2},
			{Row: 0, Column: 0},
		}

		location := maze.Location{Row: 1, Column: 2}

		assert.True(t, Contains(location, locations))
	})

}
