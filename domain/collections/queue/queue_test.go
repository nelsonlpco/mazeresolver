package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	var queue *Queue[int]

	t.Run("should be create an empty queue", func(t *testing.T) {
		queue = New[int]()

		assert.True(t, queue.IsEmpty())
	})

	t.Run("should add two itens to queue", func(t *testing.T) {
		queue.Push(1)
		queue.Push(2)

		assert.Equal(t, 2, queue.Size())
		assert.False(t, queue.IsEmpty())
	})

	t.Run("should pop one item to queue", func(t *testing.T) {
		item := queue.Pop()

		assert.Equal(t, 1, item, "expected queue item to 1")
		assert.Equal(t, 1, queue.size, "expected queue size equals 1")
		assert.False(t, queue.IsEmpty())
	})

	t.Run("should pop last item and  mark queue as empty", func(t *testing.T) {
		item := queue.Pop()

		assert.Equal(t, 2, item, "expected queue item to 1")
		assert.Equal(t, 0, queue.size, "expected queue size equals 1")
		assert.True(t, queue.IsEmpty(), "expect queue is empty")
	})

	t.Run("should pop an empty queue and expect a panic", func(t *testing.T) {
		assert.Panics(t, func() { queue.Pop() }, "expected an panic")
	})
}
