package queue

import (
	"log"
	"sync"
)

type Queue[T comparable] struct {
	data    []T
	isEmpty bool
	size    int
	l       sync.Mutex
}

func New[T comparable]() *Queue[T] {
	return &Queue[T]{
		isEmpty: true,
	}
}

func (q Queue[T]) IsEmpty() bool {
	return q.isEmpty
}

func (q Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Push(item T) {
	q.l.Lock()
	q.data = append(q.data, item)
	q.size++
	q.isEmpty = false
	q.l.Unlock()
}

func (q *Queue[T]) Pop() T {
	q.l.Lock()
	if q.isEmpty {
		log.Panicln("queue is empty")
	}
	item := q.data[0]
	q.data = q.data[1:]
	q.size--
	q.isEmpty = q.size == 0
	q.l.Unlock()

	return item
}
