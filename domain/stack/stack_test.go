package stack

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestStack(t *testing.T) {
	var stack Stack[int]

	t.Run("should create an empty stack", func(t *testing.T) {
		stack := New[int]()
		assert.True(t, stack.IsEmpty())
	})

	t.Run("should push two items", func(t *testing.T) {
		stack.Push(1)
		stack.Push(2)

		assert.False(t, stack.IsEmpty())
	})

	t.Run("should pop one items", func(t *testing.T) {
		item := stack.Pop()

		assert.False(t, stack.IsEmpty())
		assert.Equal(t, 2, item)
	})

	t.Run("should pop one and stack is empty", func(t *testing.T) {
		item := stack.Pop()

		assert.True(t, stack.IsEmpty())
		assert.Equal(t, 1, item)
	})

	t.Run("should expect an panic when try to pop an empty stack", func(t *testing.T) {
		stack := New[int]()

		assert.Panics(t, func() { stack.Pop() }, "expected an panic")
	})
}

func TestStack_paralel(t *testing.T) {
	stack := New[int]()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 100; i++ {
			stack.Push(i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			stack.Push(1)
		}

		wg.Done()
	}()

	wg.Wait()

	assert.Equal(t, 200, stack.pos)
}
