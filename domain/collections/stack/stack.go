package stack

import (
	"errors"
	"log"
	"sync"
)

var ErrorEmptyStack = errors.New("empty stack")

type Stack[T any] struct {
	l       sync.Mutex
	data    []T
	pos     int
	isEmpty bool
}

func New[T any]() Stack[T] {
	return Stack[T]{isEmpty: true}
}

func (s *Stack[T]) Push(value T) {
	s.l.Lock()
	defer s.l.Unlock()

	s.isEmpty = false
	s.data = append(s.data, value)
	s.pos++
}

func (s *Stack[T]) Pop() T {
	s.l.Lock()
	defer s.l.Unlock()

	if s.isEmpty {
		log.Panic(ErrorEmptyStack)
	}

	s.pos--
	v := s.data[s.pos]
	s.data = s.data[:s.pos]
	s.isEmpty = s.pos == 0

	return v
}

func (s Stack[T]) IsEmpty() bool {
	return s.isEmpty
}
